package user

import (
	"encoding/json"
	"net/http"

	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/itsyouonline/identityserver/communication"
	"github.com/itsyouonline/identityserver/credentials/oauth2"
	"github.com/itsyouonline/identityserver/credentials/password"
	"github.com/itsyouonline/identityserver/credentials/totp"
	contractdb "github.com/itsyouonline/identityserver/db/contract"
	organizationDb "github.com/itsyouonline/identityserver/db/organization"
	"github.com/itsyouonline/identityserver/db/registry"
	"github.com/itsyouonline/identityserver/db/user"
	"github.com/itsyouonline/identityserver/db/user/apikey"
	validationdb "github.com/itsyouonline/identityserver/db/validation"
	"github.com/itsyouonline/identityserver/identityservice/contract"
	"github.com/itsyouonline/identityserver/identityservice/invitations"
	"github.com/itsyouonline/identityserver/identityservice/organization"
	"github.com/itsyouonline/identityserver/oauthservice"
	"github.com/itsyouonline/identityserver/validation"
	"gopkg.in/mgo.v2"
)

//UsersAPI is the actual implementation of the /users api
type UsersAPI struct {
	SmsService                    communication.SMSService
	PhonenumberValidationService  *validation.IYOPhonenumberValidationService
	EmailService                  communication.EmailService
	EmailAddressValidationService *validation.IYOEmailAddressValidationService
}

func isUniquePhonenumber(user *user.User, number string, label string) (unique bool) {
	unique = true
	for _, phonenumber := range user.Phonenumbers {
		if phonenumber.Label != label && phonenumber.Phonenumber == number {
			unique = false
			return
		}
	}
	return
}

func isLastVerifiedPhoneNumber(user *user.User, number string, label string, r *http.Request) (last bool, err error) {
	last = false
	valMgr := validationdb.NewManager(r)
	validated, err := valMgr.IsPhonenumberValidated(user.Username, string(number))
	if err != nil {
		return
	}
	if validated {
		// check if this phone number is the last verified one
		uniquelabel := isUniquePhonenumber(user, number, label)
		hasotherverifiednumbers := false
		verifiednumbers, err := valMgr.GetByUsernameValidatedPhonenumbers(user.Username)
		if err != nil {
			return false, err

		}
		for _, verifiednumber := range verifiednumbers {
			if verifiednumber.Phonenumber != string(number) {
				hasotherverifiednumbers = true
				break

			}
		}
		if uniquelabel && !hasotherverifiednumbers {
			return true, nil
		}

	}
	return
}

// It is handler for POST /users
func (api UsersAPI) Post(w http.ResponseWriter, r *http.Request) {

	var u user.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)
	if err := userMgr.Save(&u); err != nil {
		log.Error("ERROR while saving user:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&u)
}

// GetUser is handler for GET /users/{username}
func (api UsersAPI) GetUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	userMgr := user.NewManager(r)

	usr, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr)
}

// RegisterNewEmailAddress is the handler for POST /users/{username}/emailaddresses
// Register a new email address
func (api UsersAPI) RegisterNewEmailAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	body := user.EmailAddress{}
	lang := r.FormValue("lang")
	if lang == "" {
		lang = organization.DefaultLanguage
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || !body.Validate() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)
	u, err := userMgr.GetByName(username)
	if handleServerError(w, "getting user by name", err) {
		return
	}

	if _, err := u.GetEmailAddressByLabel(body.Label); err == nil {
		writeErrorResponse(w, http.StatusConflict, "duplicate_label")
		return
	}

	err = userMgr.SaveEmail(username, body)
	if handleServerError(w, "saving email", err) {
		return
	}

	valMgr := validationdb.NewManager(r)
	validated, err := valMgr.IsEmailAddressValidated(username, body.EmailAddress)
	if handleServerError(w, "checking if email address is validated", err) {
		return
	}
	if !validated {
		_, err = api.EmailAddressValidationService.RequestValidation(r, username, body.EmailAddress, fmt.Sprintf("https://%s/emailvalidation", r.Host), lang)
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}

// UpdateEmailAddress is the handler for PUT /users/{username}/emailaddresses/{label}
// Updates the label and/or value of an email address
func (api UsersAPI) UpdateEmailAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	oldlabel := mux.Vars(r)["label"]
	lang := r.FormValue("lang")
	if lang == "" {
		lang = organization.DefaultLanguage
	}

	var body user.EmailAddress
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if !body.Validate() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)

	if oldlabel != body.Label {
		u, err := userMgr.GetByName(username)
		if err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if _, err := u.GetEmailAddressByLabel(body.Label); err == nil {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
	}

	if err := userMgr.SaveEmail(username, body); err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if oldlabel != body.Label {
		if err := userMgr.RemoveEmail(username, oldlabel); err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	valMgr := validationdb.NewManager(r)
	validated, err := valMgr.IsEmailAddressValidated(username, body.EmailAddress)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if !validated {
		_, err = api.EmailAddressValidationService.RequestValidation(r, username, body.EmailAddress, fmt.Sprintf("https://%s/emailvalidation", r.Host), lang)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}

// Validate email address is the handler for GET /users/{username}/emailaddress/{label}/validate
func (api UsersAPI) ValidateEmailAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)
	lang := r.FormValue("lang")
	if lang == "" {
		lang = organization.DefaultLanguage
	}
	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	email, err := userobj.GetEmailAddressByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = api.EmailAddressValidationService.RequestValidation(r, username, email.EmailAddress, fmt.Sprintf("https://%s/emailvalidation", r.Host), lang)
	w.WriteHeader(http.StatusNoContent)
}

// ListEmailAddresses is the handler for GET /users/{username}/emailaddresses
func (api UsersAPI) ListEmailAddresses(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	validated := strings.Contains(r.URL.RawQuery, "validated")
	userMgr := user.NewManager(r)
	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var emails []user.EmailAddress
	if validated {
		emails, err = api.getValidatedEmails(r, *userobj)
		if handleServerError(w, "getting validated emails", err) {
			return
		}
	} else {
		emails = userobj.EmailAddresses
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}

func (api UsersAPI) getValidatedEmails(r *http.Request, userobj user.User) ([]user.EmailAddress, error) {
	emails := make([]user.EmailAddress, 0)
	valMngr := validationdb.NewManager(r)
	validatedEmails, err := valMngr.GetByUsernameValidatedEmailAddress(userobj.Username)
	if err == nil {
		for _, email := range userobj.EmailAddresses {
			for _, validatedEmail := range validatedEmails {
				if email.EmailAddress == validatedEmail.EmailAddress {
					emails = append(emails, email)
					break
				}
			}
		}
	}
	return emails, err
}

// DeleteEmailAddress is the handler for DELETE /users/{username}/emailaddresses/{label}
// Removes an email address
func (api UsersAPI) DeleteEmailAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]

	userMgr := user.NewManager(r)
	valMgr := validationdb.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	email, err := u.GetEmailAddressByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if len(u.EmailAddresses) == 1 {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err = userMgr.RemoveEmail(username, label); err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err = valMgr.RemoveValidatedEmailAddress(username, email.EmailAddress); err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeleteGithubAccount is the handler for DELETE /users/{username}/github
// Delete the associated Github account.
func (api UsersAPI) DeleteGithubAccount(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)
	err := userMgr.DeleteGithubAccount(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteFacebookAccount is the handler for DELETE /users/{username}/facebook
// Delete the associated facebook account
func (api UsersAPI) DeleteFacebookAccount(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	userMgr := user.NewManager(r)
	err := userMgr.DeleteFacebookAccount(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UpdatePassword is the handler for PUT /users/{username}/password
func (api UsersAPI) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	body := struct {
		Currentpassword string `json:"currentpassword"`
		Newpassword     string `json:"newpassword"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userMgr := user.NewManager(r)
	exists, err := userMgr.Exists(username)
	if !exists || err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	passwordMgr := password.NewManager(r)
	passwordok, err := passwordMgr.Validate(username, body.Currentpassword)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if !passwordok {
		writeErrorResponse(w, 422, "incorrect_password")
		return
	}
	err = passwordMgr.Save(username, body.Newpassword)
	if err != nil {
		writeErrorResponse(w, 422, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetUserInformation is the handler for GET /users/{username}/info
func (api UsersAPI) GetUserInformation(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	requestingClient, validClient := context.Get(r, "client_id").(string)
	if !validClient {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	availableScopes, _ := context.Get(r, "availablescopes").(string)
	isAdmin := oauth2.CheckScopes([]string{"user:admin"}, oauth2.SplitScopeString(availableScopes))

	authorization, err := userMgr.GetAuthorization(username, requestingClient)
	if handleServerError(w, "getting authorization", err) {
		return
	}

	//Create an administrator authorization
	if authorization == nil && isAdmin {
		authorization = &user.Authorization{
			Name:           true,
			Github:         true,
			Facebook:       true,
			Addresses:      []user.AuthorizationMap{},
			BankAccounts:   []user.AuthorizationMap{},
			DigitalWallet:  []user.DigitalWalletAuthorization{},
			EmailAddresses: []user.AuthorizationMap{},
			Phonenumbers:   []user.AuthorizationMap{},
			PublicKeys:     []user.AuthorizationMap{},
		}
		for _, address := range userobj.Addresses {
			authorization.Addresses = append(authorization.Addresses, user.AuthorizationMap{RequestedLabel: address.Label, RealLabel: address.Label})
		}
		for _, a := range userobj.BankAccounts {
			authorization.BankAccounts = append(authorization.BankAccounts, user.AuthorizationMap{RequestedLabel: a.Label, RealLabel: a.Label})
		}
		for _, a := range userobj.DigitalWallet {
			authorization.DigitalWallet = append(authorization.DigitalWallet, user.DigitalWalletAuthorization{Currency: a.CurrencySymbol, AuthorizationMap: user.AuthorizationMap{RequestedLabel: a.Label, RealLabel: a.Label}})
		}
		for _, a := range userobj.EmailAddresses {
			authorization.EmailAddresses = append(authorization.EmailAddresses, user.AuthorizationMap{RequestedLabel: a.Label, RealLabel: a.Label})
		}
		for _, a := range userobj.Phonenumbers {
			authorization.Phonenumbers = append(authorization.Phonenumbers, user.AuthorizationMap{RequestedLabel: a.Label, RealLabel: a.Label})
		}
		for _, a := range userobj.PublicKeys {
			authorization.PublicKeys = append(authorization.PublicKeys, user.AuthorizationMap{RequestedLabel: a.Label, RealLabel: a.Label})
		}

	}
	if authorization == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	respBody := &Userview{
		Username:       userobj.Username,
		Github:         user.GithubAccount{},
		Facebook:       user.FacebookAccount{},
		Addresses:      []user.Address{},
		EmailAddresses: []user.EmailAddress{},
		Phonenumbers:   []user.Phonenumber{},
		BankAccounts:   []user.BankAccount{},
		DigitalWallet:  []user.DigitalAssetAddress{},
		OwnerOf: user.OwnerOf{
			EmailAddresses: []string{},
		},
		PublicKeys: []user.PublicKey{},
	}

	if authorization.Name {
		respBody.Firstname = userobj.Firstname
		respBody.Lastname = userobj.Lastname
	}

	if authorization.Github {
		respBody.Github = userobj.Github
	}

	if authorization.Facebook {
		respBody.Facebook = userobj.Facebook
	}

	if authorization.Addresses != nil {
		for _, addressmap := range authorization.Addresses {
			address, err := userobj.GetAddressByLabel(addressmap.RealLabel)
			if err == nil {
				newaddress := user.Address{
					Label:      addressmap.RequestedLabel,
					City:       address.City,
					Country:    address.Country,
					Nr:         address.Nr,
					Other:      address.Other,
					Postalcode: address.Postalcode,
					Street:     address.Street,
				}
				respBody.Addresses = append(respBody.Addresses, newaddress)
			} else {
				log.Debug(err)
			}
		}
	}

	if authorization.EmailAddresses != nil {
		for _, emailmap := range authorization.EmailAddresses {
			email, err := userobj.GetEmailAddressByLabel(emailmap.RealLabel)
			if err == nil {
				newemail := user.EmailAddress{
					Label:        emailmap.RequestedLabel,
					EmailAddress: email.EmailAddress,
				}
				respBody.EmailAddresses = append(respBody.EmailAddresses, newemail)
			} else {
				log.Debug(err)
			}
		}
	}

	if authorization.Phonenumbers != nil {
		for _, phonemap := range authorization.Phonenumbers {
			phonenumber, err := userobj.GetPhonenumberByLabel(phonemap.RealLabel)
			if err == nil {
				newnumber := user.Phonenumber{
					Label:       phonemap.RequestedLabel,
					Phonenumber: phonenumber.Phonenumber,
				}
				respBody.Phonenumbers = append(respBody.Phonenumbers, newnumber)
			} else {
				log.Debug(err)
			}
		}
	}

	if authorization.BankAccounts != nil {
		for _, bankmap := range authorization.BankAccounts {
			bank, err := userobj.GetBankAccountByLabel(bankmap.RealLabel)
			if err == nil {
				newbank := user.BankAccount{
					Label:   bankmap.RequestedLabel,
					Bic:     bank.Bic,
					Country: bank.Country,
					Iban:    bank.Iban,
				}
				respBody.BankAccounts = append(respBody.BankAccounts, newbank)
			} else {
				log.Debug(err)
			}
		}
	}

	if authorization.DigitalWallet != nil {
		for _, addressMap := range authorization.DigitalWallet {
			walletAddress, err := userobj.GetDigitalAssetAddressByLabel(addressMap.RealLabel)
			if err == nil {
				walletAddress.Label = addressMap.RequestedLabel
				respBody.DigitalWallet = append(respBody.DigitalWallet, walletAddress)
			} else {
				log.Debug(err)
			}
		}
	}

	if authorization.PublicKeys != nil {
		for _, publicKeyMap := range authorization.PublicKeys {
			publicKey, err := userobj.GetPublicKeyByLabel(publicKeyMap.RealLabel)
			if err == nil {
				publicKey.Label = publicKeyMap.RequestedLabel
				respBody.PublicKeys = append(respBody.PublicKeys, publicKey)
			} else {
				log.Debug(err)
			}
		}
	}

	if authorization.OwnerOf.EmailAddresses != nil {
		respBody.OwnerOf.EmailAddresses = authorization.OwnerOf.EmailAddresses
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)
}

// RegisterNewPhonenumber is the handler for POST /users/{username}/phonenumbers
// Register a new phonenumber
func (api UsersAPI) RegisterNewPhonenumber(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	body := user.Phonenumber{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !body.Validate() {
		log.Debug("Invalid phonenumber: ", body.Phonenumber)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//Check if this label is already used
	_, err = u.GetPhonenumberByLabel(body.Label)
	if err == nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err := userMgr.SavePhone(username, body); err != nil {
		log.Error("ERROR while saving a phonenumber - ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// respond with created phone number.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(body)
}

// GetUserPhoneNumbers is the handler for GET /users/{username}/phonenumbers
func (api UsersAPI) GetUserPhoneNumbers(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	validated := strings.Contains(r.URL.RawQuery, "validated")
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var phonenumbers []user.Phonenumber
	if validated {
		phonenumbers = make([]user.Phonenumber, 0)
		valMngr := validationdb.NewManager(r)
		validatednumbers, err := valMngr.GetByUsernameValidatedPhonenumbers(username)
		if err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		for _, number := range userobj.Phonenumbers {
			for _, validatednumber := range validatednumbers {
				if number.Phonenumber == validatednumber.Phonenumber {
					phonenumbers = append(phonenumbers, number)
					break
				}
			}
		}
	} else {
		phonenumbers = userobj.Phonenumbers
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phonenumbers)
}

// GetUserPhonenumberByLabel is the handler for GET /users/{username}/phonenumbers/{label}
func (api UsersAPI) GetUserPhonenumberByLabel(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	phonenumber, err := userobj.GetPhonenumberByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phonenumber)
}

// ValidatePhoneNumber is the handler for POST /users/{username}/phonenumbers/{label}/validate
func (api UsersAPI) ValidatePhoneNumber(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	lang := r.FormValue("lang")
	if lang == "" {
		lang = "en"
	}

	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	phonenumber, err := userobj.GetPhonenumberByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	validationKey := ""
	validationKey, err = api.PhonenumberValidationService.RequestValidation(r, username, phonenumber, fmt.Sprintf("https://%s/phonevalidation", r.Host), lang)
	response := struct {
		ValidationKey string `json:"validationkey"`
	}{
		ValidationKey: validationKey,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

// VerifyPhoneNumber is the handler for PUT /users/{username}/phonenumbers/{label}/validate
func (api UsersAPI) VerifyPhoneNumber(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	values := struct {
		Smscode       string `json:"smscode"`
		ValidationKey string `json:"validationkey"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
		log.Debug("Error decoding the ProcessPhonenumberConfirmation request:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = userobj.GetPhonenumberByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	err = api.PhonenumberValidationService.ConfirmValidation(r, values.ValidationKey, values.Smscode)
	if err != nil {
		log.Debug(err)
		if err == validation.ErrInvalidCode || err == validation.ErrInvalidOrExpiredKey {
			writeErrorResponse(w, 422, err.Error())
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// UpdatePhonenumber is the handler for PUT /users/{username}/phonenumbers/{label}
// Update the label and/or value of an existing phonenumber.
func (api UsersAPI) UpdatePhonenumber(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	oldlabel := mux.Vars(r)["label"]

	body := user.Phonenumber{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !body.Validate() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	oldnumber, err := u.GetPhonenumberByLabel(oldlabel)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if oldlabel != body.Label {
		// Check if there already is another phone number with the new label
		_, err := u.GetPhonenumberByLabel(body.Label)
		if err == nil {
			writeErrorResponse(w, http.StatusConflict, "duplicate_label")
			return
		}
	}

	if oldnumber.Phonenumber != body.Phonenumber {
		last, err := isLastVerifiedPhoneNumber(u, oldnumber.Phonenumber, oldlabel, r)
		if err != nil {
			log.Error("ERROR while verifying last verified number - ", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if last {
			writeErrorResponse(w, http.StatusConflict, "cannot_delete_last_verified_phone_number")
			return
		}
	}

	if err = userMgr.SavePhone(username, body); err != nil {
		log.Error("ERROR while saving phonenumber - ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if oldlabel != body.Label {
		if err := userMgr.RemovePhone(username, oldlabel); err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	valMgr := validationdb.NewManager(r)
	if oldnumber.Phonenumber != body.Phonenumber && isUniquePhonenumber(u, oldnumber.Phonenumber, oldlabel) {
		valMgr.RemoveValidatedPhonenumber(username, oldnumber.Phonenumber)
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)

}

// DeletePhonenumber is the handler for DELETE /users/{username}/phonenumbers/{label}
// Removes a phonenumber
func (api UsersAPI) DeletePhonenumber(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)
	valMgr := validationdb.NewManager(r)
	force := r.URL.Query().Get("force") == "true"

	usr, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	number, err := usr.GetPhonenumberByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	last, err := isLastVerifiedPhoneNumber(usr, number.Phonenumber, label, r)
	if err != nil {
		log.Error("ERROR while checking if number can be deleted:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return

	}
	if last {
		hasTOTP := false
		if !force {
			writeErrorResponse(w, http.StatusConflict, "warning_delete_last_verified_phone_number")
			return
		} else {
			totpMgr := totp.NewManager(r)
			hasTOTP, err = totpMgr.HasTOTP(username)
		}
		if !hasTOTP {
			writeErrorResponse(w, http.StatusConflict, "cannot_delete_last_verified_phone_number")
			return
		}
	}

	// check if the phonenumber is unique or if there are duplicates
	uniqueNumber := isUniquePhonenumber(usr, number.Phonenumber, label)

	if err := userMgr.RemovePhone(username, label); err != nil {
		log.Error("ERROR while saving user:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// only remove the phonenumber from the validatedphonenumber collection if there are no duplicates
	if uniqueNumber {
		if err := valMgr.RemoveValidatedPhonenumber(username, number.Phonenumber); err != nil {
			log.Error("ERROR while saving user:\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// CreateUserBankAccount is handler for POST /users/{username}/banks
// Create new bank account
func (api UsersAPI) CreateUserBankAccount(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	bank := user.BankAccount{}

	if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
		log.Error("Error while decoding the body: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !bank.Validate() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	usr, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//Check if this label is already used
	_, err = usr.GetBankAccountByLabel(bank.Label)
	if err == nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err := userMgr.SaveBank(usr, bank); err != nil {
		log.Error("ERROR while saving address:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// respond with created bank account
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(bank)
}

// GetUserBankAccounts It is handler for GET /users/{username}/banks
func (api UsersAPI) GetUserBankAccounts(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.BankAccounts)
}

// GetUserBankAccountByLabel is handler for GET /users/{username}/banks/{label}
func (api UsersAPI) GetUserBankAccountByLabel(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	bank, err := userobj.GetBankAccountByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bank)
}

// UpdateUserBankAccount is handler for PUT /users/{username}/banks/{label}
// Update an existing bankaccount and label.
func (api UsersAPI) UpdateUserBankAccount(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	oldlabel := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	newbank := user.BankAccount{}

	if err := json.NewDecoder(r.Body).Decode(&newbank); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !newbank.Validate() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	oldbank, err := user.GetBankAccountByLabel(oldlabel)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if oldbank.Label != newbank.Label {
		_, err := user.GetBankAccountByLabel(newbank.Label)
		if err == nil {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
	}

	if err = userMgr.SaveBank(user, newbank); err != nil {
		log.Error("ERROR while saving bank - ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if oldlabel != newbank.Label {
		if err := userMgr.RemoveBank(user, oldlabel); err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newbank)
}

// DeleteUserBankAccount is handler for DELETE /users/{username}/banks/{label}
// Delete a BankAccount
func (api UsersAPI) DeleteUserBankAccount(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	usr, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = usr.GetBankAccountByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err := userMgr.RemoveBank(usr, label); err != nil {
		log.Error("ERROR while saving user:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RegisterNewAddress is the handler for POST /users/{username}/addresses
// Register a new address
func (api UsersAPI) RegisterNewAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	address := user.Address{}

	if err := json.NewDecoder(r.Body).Decode(&address); err != nil {
		log.Debug("Error while decoding the body: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !address.Validate() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	//Check if this label is already used
	_, err = u.GetAddressByLabel(address.Label)
	if err == nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err := userMgr.SaveAddress(username, address); err != nil {
		log.Error("ERROR while saving address:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// respond with created phone number.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(address)
}

// GetUserAddresses is handler for GET /users/{username}/addresses
func (api UsersAPI) GetUserAddresses(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)

	usr, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr.Addresses)
}

// GetUserAddressByLabel is handler for GET /users/{username}/addresses/{label}
func (api UsersAPI) GetUserAddressByLabel(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	address, err := userobj.GetAddressByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}

// UpdateAddress is the handler for PUT /users/{username}/addresses/{label}
// Update the label and/or value of an existing address.
func (api UsersAPI) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	oldlabel := mux.Vars(r)["label"]

	newaddress := user.Address{}
	if err := json.NewDecoder(r.Body).Decode(&newaddress); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !newaddress.Validate() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = u.GetAddressByLabel(oldlabel)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if oldlabel != newaddress.Label {
		_, err = u.GetAddressByLabel(newaddress.Label)
		if err == nil {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
	}

	if err = userMgr.SaveAddress(username, newaddress); err != nil {
		log.Error("ERROR while saving address - ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if oldlabel != newaddress.Label {
		if err := userMgr.RemoveAddress(username, oldlabel); err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newaddress)
}

// DeleteAddress is the handler for DELETE /users/{username}/addresses/{label}
// Removes an address
func (api UsersAPI) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = u.GetAddressByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err = userMgr.RemoveAddress(username, label); err != nil {
		log.Error("ERROR while saving address:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetUserContracts is handler for GET /users/{username}/contracts
// Get the contracts where the user is 1 of the parties. Order descending by date.
func (api UsersAPI) GetUserContracts(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	includedparty := contractdb.Party{Type: "user", Name: username}
	contract.FindContracts(w, r, includedparty)
}

// RegisterNewContract is handler for GET /users/{username}/contracts
func (api UsersAPI) RegisterNewContract(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	includedparty := contractdb.Party{Type: "user", Name: username}
	contract.CreateContract(w, r, includedparty)

}

// GetNotifications is handler for GET /users/{username}/notifications
// Get the list of notifications, these are pending invitations or approvals
func (api UsersAPI) GetNotifications(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	type NotificationList struct {
		Approvals               []invitations.JoinOrganizationInvitation `json:"approvals"`
		ContractRequests        []contractdb.ContractSigningRequest      `json:"contractRequests"`
		Invitations             []invitations.JoinOrganizationInvitation `json:"invitations"`
		MissingScopes           []organizationDb.MissingScope            `json:"missingscopes"`
		OrganizationInvitations []invitations.JoinOrganizationInvitation `json:"organizationinvitations"`
	}
	var notifications NotificationList

	invitationMgr := invitations.NewInvitationManager(r)
	valMgr := validationdb.NewManager(r)

	userOrgRequests, err := invitationMgr.FilterByUser(username, "pending")
	if handleServerError(w, "getting invitations by user", err) {
		return
	}

	// Add the invites for the users verified phone numbers. This is required if the invite
	// was added before the phone number was verified, and no invite sms was send
	validatedPhonenumbers, err := valMgr.GetByUsernameValidatedPhonenumbers(username)
	if handleServerError(w, "getting verified phone numbers", err) {
		return
	}
	for _, number := range validatedPhonenumbers {
		phonenumberRequests, err := invitationMgr.FilterByPhonenumber(number.Phonenumber, "pending")
		if handleServerError(w, "getting invitations by user for phonenumber", err) {
			return
		}
		userOrgRequests = append(userOrgRequests, phonenumberRequests...)
	}

	// Add the invites for the users verified email addresses. This is required if the invite
	// was added before the email address was verified, and no invite email was send
	validatedEmailaddresses, err := valMgr.GetByUsernameValidatedEmailAddress(username)
	if handleServerError(w, "getting verified email addresses", err) {
		return
	}
	for _, email := range validatedEmailaddresses {
		emailRequests, err := invitationMgr.FilterByEmail(email.EmailAddress, "pending")
		if handleServerError(w, "getting invitations by user for email", err) {
			return
		}
		userOrgRequests = append(userOrgRequests, emailRequests...)
	}

	// Add the invites for the organizations where this user is an owner
	orgMgr := organizationDb.NewManager(r)

	orgs, err := orgMgr.AllByUser(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var ownedOrgs []string

	for _, org := range orgs {
		if exists(username, org.Owners) {
			ownedOrgs = append(ownedOrgs, org.Globalid)
		}
	}

	//var orgInvites []invitations.JoinOrganizationInvitation
	orgInvites := make([]invitations.JoinOrganizationInvitation, 0)

	for _, org := range ownedOrgs {
		invites, err := invitationMgr.GetOpenOrganizationInvites(org)
		if err != nil {
			log.Error("Error while loading all invites where the organization ", org, " is invited")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		orgInvites = append(orgInvites, invites...)
	}

	notifications.Invitations = userOrgRequests
	notifications.OrganizationInvitations = orgInvites
	// TODO: Get Approvals and Contract requests
	notifications.Approvals = []invitations.JoinOrganizationInvitation{}
	notifications.ContractRequests = []contractdb.ContractSigningRequest{}
	extraOrganizations := []string{}
	for _, invitation := range notifications.Invitations {
		extraOrganizations = append(extraOrganizations, invitation.Organization)
	}
	err, notifications.MissingScopes = getMissingScopesForOrganizations(r, username, extraOrganizations)
	if handleServerError(w, "getting missing scopes", err) {
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(&notifications)
}

func getMissingScopesForOrganizations(r *http.Request, username string, extraOrganizations []string) (error, []organizationDb.MissingScope) {
	orgMgr := organizationDb.NewManager(r)
	userMgr := user.NewManager(r)
	err, organizations := orgMgr.ListByUserOrGlobalID(username, extraOrganizations)

	if err != nil {
		return err, nil
	}
	authorizations, err := userMgr.GetAuthorizationsByUser(username)
	if err != nil {
		return err, nil
	}
	missingScopes := []organizationDb.MissingScope{}
	authorizationsMap := make(map[string]user.Authorization)
	for _, authorization := range authorizations {
		authorizationsMap[authorization.Username] = authorization
	}
	for _, organization := range organizations {
		scopes := []string{}
		for _, requiredScope := range organization.RequiredScopes {
			hasScope := false
			if authorization, hasKey := authorizationsMap[username]; hasKey {
				hasScope = requiredScope.IsAuthorized(authorization)
			} else {
				hasScope = false
			}
			if !hasScope {
				scopes = append(scopes, requiredScope.Scope)
			}
		}
		if len(scopes) > 0 {
			missingScope := organizationDb.MissingScope{
				Scopes:       scopes,
				Organization: organization.Globalid,
			}
			missingScopes = append(missingScopes, missingScope)
		}
	}
	return nil, missingScopes
}

// usernameorganizationsGet is the handler for GET /users/{username}/organizations
// Get the list organizations a user is owner of member of
func (api UsersAPI) usernameorganizationsGet(w http.ResponseWriter, r *http.Request) {

}

// GetAllAuthorizations is the handler for GET /users/{username}/authorizations
// Get the list of authorizations.
func (api UsersAPI) GetAllAuthorizations(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	userMgr := user.NewManager(r)

	authorizations, err := userMgr.GetAuthorizationsByUser(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(authorizations)

}

// GetAuthorization is the handler for GET /users/{username}/authorizations/{grantedTo}
// Get the authorization for a specific organization.
func (api UsersAPI) GetAuthorization(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	grantedTo := mux.Vars(r)["grantedTo"]

	userMgr := user.NewManager(r)

	authorization, err := userMgr.GetAuthorization(username, grantedTo)
	if handleServerError(w, "Getting authorization by user", err) {
		return
	}
	if authorization == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(authorization)
}

func FilterAuthorizationMaps(s []user.AuthorizationMap) []user.AuthorizationMap {
	var p []user.AuthorizationMap
	for _, v := range s {
		if v.RealLabel != "" {
			p = append(p, v)
		}
	}
	return p
}

func FilterDigitalWallet(s []user.DigitalWalletAuthorization) []user.DigitalWalletAuthorization {
	var p []user.DigitalWalletAuthorization
	for _, v := range s {
		if v.RealLabel != "" {
			p = append(p, v)
		}
	}
	return p
}

func FilterOwnerOf(s user.OwnerOf, verifiedEmails []user.EmailAddress) user.OwnerOf {
	var o user.OwnerOf
	for _, verifiedEmail := range verifiedEmails {
		for _, mail := range s.EmailAddresses {
			if mail == verifiedEmail.EmailAddress {
				o.EmailAddresses = append(o.EmailAddresses, verifiedEmail.EmailAddress)
			}
		}
	}
	return o
}

// UpdateAuthorization is the handler for PUT /users/{username}/authorizations/{grantedTo}
// Modify which information an organization is able to see.
func (api UsersAPI) UpdateAuthorization(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	grantedTo := mux.Vars(r)["grantedTo"]

	authorization := &user.Authorization{}

	if err := json.NewDecoder(r.Body).Decode(authorization); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userMgr := user.NewManager(r)
	verifiedEmails := []user.EmailAddress{}
	if len(authorization.OwnerOf.EmailAddresses) != 0 {
		userobj, err := userMgr.GetByName(username)
		if handleServerError(w, "getting user by name", err) {
			return
		}
		verifiedEmails, err = api.getValidatedEmails(r, *userobj)
		if handleServerError(w, "getting verified emails", err) {
			return
		}
	}

	authorization.Username = username
	authorization.GrantedTo = grantedTo
	authorization.Addresses = FilterAuthorizationMaps(authorization.Addresses)
	authorization.EmailAddresses = FilterAuthorizationMaps(authorization.EmailAddresses)
	authorization.Phonenumbers = FilterAuthorizationMaps(authorization.Phonenumbers)
	authorization.BankAccounts = FilterAuthorizationMaps(authorization.BankAccounts)
	authorization.PublicKeys = FilterAuthorizationMaps(authorization.PublicKeys)
	authorization.DigitalWallet = FilterDigitalWallet(authorization.DigitalWallet)
	authorization.OwnerOf = FilterOwnerOf(authorization.OwnerOf, verifiedEmails)

	err := userMgr.UpdateAuthorization(authorization)
	if handleServerError(w, "updating authorization", err) {
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(authorization)
}

// DeleteAuthorization is the handler for DELETE /users/{username}/authorizations/{grantedTo}
// Remove the authorization for an organization, the granted organization will no longer
// have access the user's information.
func (api UsersAPI) DeleteAuthorization(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	grantedTo := mux.Vars(r)["grantedTo"]

	userMgr := user.NewManager(r)

	err := userMgr.DeleteAuthorization(username, grantedTo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (api UsersAPI) AddAPIKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	body := struct {
		Label string
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if !user.IsValidLabel(body.Label) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	apikeyMgr := apikey.NewManager(r)
	apiKey := apikey.NewAPIKey(username, body.Label)
	apikeyMgr.Save(apiKey)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(apiKey)
}

func (api UsersAPI) GetAPIKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	apikeyMgr := apikey.NewManager(r)
	apiKey, err := apikeyMgr.GetByUsernameAndLabel(username, label)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiKey)
}

func (api UsersAPI) UpdateAPIKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	apikeyMgr := apikey.NewManager(r)
	body := struct {
		Label string
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	apiKey, err := apikeyMgr.GetByUsernameAndLabel(username, label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	apiKey.Label = body.Label
	apikeyMgr.Save(apiKey)
	w.WriteHeader(http.StatusNoContent)

}
func (api UsersAPI) DeleteAPIKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	apikeyMgr := apikey.NewManager(r)
	apikeyMgr.Delete(username, label)

}
func (api UsersAPI) ListAPIKeys(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	apikeyMgr := apikey.NewManager(r)
	apikeys, err := apikeyMgr.GetByUser(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if apikeys == nil {
		apikeys = []apikey.APIKey{}
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(apikeys)
}

// AddPublicKey Add a public key
func (api UsersAPI) AddPublicKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	body := user.PublicKey{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(body.PublicKey, "ssh-rsa AAAAB3NzaC1yc2E") {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	mgr := user.NewManager(r)

	usr, err := mgr.GetByName(username)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		log.Error("Error while getting user: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = usr.GetPublicKeyByLabel(body.Label)
	if err == nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	err = mgr.SavePublicKey(username, body)
	if err != nil {
		log.Error("error while saving public key: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)
}

// GetPublicKey Get the public key associated with a label
func (api UsersAPI) GetPublicKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	publickey, err := userobj.GetPublicKeyByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publickey)
}

// UpdatePublicKey Update the label and/or value of an existing public key.
func (api UsersAPI) UpdatePublicKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	oldlabel := mux.Vars(r)["label"]

	body := user.PublicKey{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !body.Validate() {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(body.PublicKey, "ssh-rsa AAAAB3NzaC1yc2E") {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	userMgr := user.NewManager(r)

	u, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = u.GetPublicKeyByLabel(oldlabel)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if oldlabel != body.Label {
		// Check if there already is another public key with the new label
		_, err := u.GetPublicKeyByLabel(body.Label)
		if err == nil {
			writeErrorResponse(w, http.StatusConflict, "duplicate_label")
			return
		}
	}

	if err = userMgr.SavePublicKey(username, body); err != nil {
		log.Error("ERROR while saving public key - ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if oldlabel != body.Label {
		if err := userMgr.RemovePublicKey(username, oldlabel); err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(body)

}

// DeletePublicKey Deletes a public key
func (api UsersAPI) DeletePublicKey(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	label := mux.Vars(r)["label"]
	userMgr := user.NewManager(r)

	usr, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err = usr.GetPublicKeyByLabel(label)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	if err := userMgr.RemovePublicKey(username, label); err != nil {
		log.Error("ERROR while removing public key:\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//ListPublicKeys lists all public keys
func (api UsersAPI) ListPublicKeys(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)
	userobj, err := userMgr.GetByName(username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var publicKeys []user.PublicKey

	publicKeys = userobj.PublicKeys

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publicKeys)
}

// UpdateName is the handler for PUT /users/{username}/name
func (api UsersAPI) UpdateName(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	values := struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	userMgr := user.NewManager(r)
	exists, err := userMgr.Exists(username)
	if !exists || err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	err = userMgr.UpdateName(username, values.Firstname, values.Lastname)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetTwoFAMethods is the handler for GET /users/{username}/twofamethods
// Get the possible two factor authentication methods
func (api UsersAPI) GetTwoFAMethods(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	userMgr := user.NewManager(r)
	userFromDB, err := userMgr.GetByName(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := struct {
		Totp bool               `json:"totp"`
		Sms  []user.Phonenumber `json:"sms"`
	}{}
	totpMgr := totp.NewManager(r)
	response.Totp, err = totpMgr.HasTOTP(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	valMgr := validationdb.NewManager(r)
	verifiedPhones, err := valMgr.GetByUsernameValidatedPhonenumbers(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	for _, validatedPhoneNumber := range verifiedPhones {
		for _, number := range userFromDB.Phonenumbers {
			if number.Phonenumber == string(validatedPhoneNumber.Phonenumber) {
				response.Sms = append(response.Sms, number)
			}
		}
	}
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
	return
}

// GetTOTPSecret is the handler for GET /users/{username}/totp/
// Gets a new TOTP secret
func (api UsersAPI) GetTOTPSecret(w http.ResponseWriter, r *http.Request) {
	token, err := totp.NewToken()
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	response := struct {
		Totpsecret string `json:"totpsecret"`
		TotpIssuer string `json:"totpissuer"`
	}{
		Totpsecret: token.Secret,
		TotpIssuer: totp.GetIssuer(r),
	}
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

// SetupTOTP is the handler for POST /users/{username}/totp/
// Configures TOTP authentication for this user
func (api UsersAPI) SetupTOTP(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	values := struct {
		TotpSecret string `json:"totpsecret"`
		TotpCode   string `json:"totpcode"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&values); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	totpMgr := totp.NewManager(r)
	err := totpMgr.Save(username, values.TotpSecret)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	valid, err := totpMgr.Validate(username, values.TotpCode)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if !valid {
		err := totpMgr.Remove(username)
		if err != nil {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(422)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// RemoveTOTP is the handler for DELETE /users/{username}/totp/
// Removes TOTP authentication for this user, if possible.
func (api UsersAPI) RemoveTOTP(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	valMngr := validationdb.NewManager(r)
	hasValidatedPhones, err := valMngr.HasValidatedPhones(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if !hasValidatedPhones {
		w.WriteHeader(http.StatusConflict)
		return
	}
	totpMgr := totp.NewManager(r)
	err = totpMgr.Remove(username)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// LeaveOrganization is the handler for DELETE /users/{username}/organizations/{globalid}/leave
// Removes the user from an organization
func (api UsersAPI) LeaveOrganization(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	organizationGlobalId := mux.Vars(r)["globalid"]
	orgMgr := organizationDb.NewManager(r)
	userMgr := user.NewManager(r)
	oauthMgr := oauthservice.NewManager(r)
	err := orgMgr.RemoveUser(organizationGlobalId, username)
	if err == mgo.ErrNotFound {
		writeErrorResponse(w, http.StatusNotFound, "user_not_found")
		return
	} else if handleServerError(w, "removing user from organization", err) {
		return
	}
	err = userMgr.DeleteAuthorization(username, organizationGlobalId)
	if handleServerError(w, "removing authorization", err) {
		return
	}
	err = oauthMgr.RemoveOrganizationScopes(organizationGlobalId, username)
	if handleServerError(w, "removing organization scopes", err) {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// ListUserRegistry is the handler for GET /users/{username}/registry
// Lists the Registry entries
func (api UsersAPI) ListUserRegistry(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	mgr := registry.NewManager(r)
	registryEntries, err := mgr.ListRegistryEntries(username, "")
	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registryEntries)
}

// AddUserRegistryEntry is the handler for POST /users/{username}/registry
// Adds a RegistryEntry to the user's registry, if the key is already used, it is overwritten.
func (api UsersAPI) AddUserRegistryEntry(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	registryEntry := registry.RegistryEntry{}

	if err := json.NewDecoder(r.Body).Decode(&registryEntry); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if err := registryEntry.Validate(); err != nil {
		log.Debug("Invalid registry entry: ", registryEntry)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	mgr := registry.NewManager(r)
	err := mgr.UpsertRegistryEntry(username, "", registryEntry)

	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(registryEntry)
}

// GetUserRegistryEntry is the handler for GET /users/{username}/registry/{key}
// Get a RegistryEntry from the user's registry.
func (api UsersAPI) GetUserRegistryEntry(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	key := mux.Vars(r)["key"]

	mgr := registry.NewManager(r)
	registryEntry, err := mgr.GetRegistryEntry(username, "", key)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if registryEntry == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registryEntry)
}

// DeleteUserRegistryEntry is the handler for DELETE /users/{username}/registry/{key}
// Removes a RegistryEntry from the user's registry
func (api UsersAPI) DeleteUserRegistryEntry(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	key := mux.Vars(r)["key"]

	mgr := registry.NewManager(r)
	err := mgr.DeleteRegistryEntry(username, "", key)

	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func writeErrorResponse(responseWrite http.ResponseWriter, httpStatusCode int, message string) {
	log.Debug(httpStatusCode, message)
	errorResponse := struct {
		Error string `json:"error"`
	}{
		Error: message,
	}
	responseWrite.WriteHeader(httpStatusCode)
	json.NewEncoder(responseWrite).Encode(&errorResponse)
}

func handleServerError(responseWriter http.ResponseWriter, actionText string, err error) bool {
	if err != nil {
		log.Error("Users api: Error while "+actionText, " - ", err)
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return true
	}
	return false
}

func exists(value string, list []string) bool {
	for _, val := range list {
		if val == value {
			return true
		}
	}

	return false
}
