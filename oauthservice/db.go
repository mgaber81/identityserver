package oauthservice

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"fmt"
	"github.com/itsyouonline/identityserver/db"
	"strings"
)

const (
	requestsCollectionName     = "oauth_authorizationrequests"
	tokensCollectionName       = "oauth_accesstokens"
	clientsCollectionName      = "oauth_clients"
	refreshTokenCollectionName = "oauth_refreshtokens"
)

//InitModels initialize models in mongo, if required.
func InitModels() {
	index := mgo.Index{
		Key:    []string{"authorizationcode"},
		Unique: true,
	} //Do not drop duplicates since it would hijack another authorizationrequest, better to error out

	db.EnsureIndex(requestsCollectionName, index)

	//TODO: unique username/clientid combination

	automaticExpiration := mgo.Index{
		Key:         []string{"createdat"},
		ExpireAfter: time.Second * 10,
		Background:  true,
	}
	db.EnsureIndex(requestsCollectionName, automaticExpiration)

	index = mgo.Index{
		Key:    []string{"accesstoken"},
		Unique: true,
	} //Do not drop duplicates since it would hijack another authorizationrequest, better to error out

	db.EnsureIndex(tokensCollectionName, index)

	//TODO: unique username/clientid combination

	automaticExpiration = mgo.Index{
		Key:         []string{"createdat"},
		ExpireAfter: AccessTokenExpiration,
		Background:  true,
	}
	db.EnsureIndex(tokensCollectionName, automaticExpiration)

	index = mgo.Index{
		Key:    []string{"clientid", "label"},
		Unique: true,
	}
	db.EnsureIndex(clientsCollectionName, index)

	index = mgo.Index{
		Key:    []string{"refreshtoken"},
		Unique: true,
	} //Do not drop duplicates since it would hijack another refreshtoken, better to error out

	db.EnsureIndex(refreshTokenCollectionName, index)
	automaticExpiration = mgo.Index{
		Key:         []string{"lastused"},
		ExpireAfter: time.Second * 86400 * 30,
		Background:  true,
	}
	db.EnsureIndex(refreshTokenCollectionName, automaticExpiration)

}

//Manager is used to store
type Manager struct {
	session *mgo.Session
}

//NewManager creates and initializes a new Manager
func NewManager(r *http.Request) *Manager {
	session := db.GetDBSession(r)
	return &Manager{
		session: session,
	}
}

//ClientManager defines a client persistence interface
type ClientManager interface {
	//AllByClientID retrieves all clients with a given ID
	AllByClientID(clientID string) ([]*Oauth2Client, error)
}

//getAuthorizationRequestCollection returns the mongo collection for the authorizationRequests
func (m *Manager) getAuthorizationRequestCollection() *mgo.Collection {
	return db.GetCollection(m.session, requestsCollectionName)
}

//getAccessTokenCollection returns the mongo collection for the accessTokens
func (m *Manager) getAccessTokenCollection() *mgo.Collection {
	return db.GetCollection(m.session, tokensCollectionName)
}

//getRefreshTokenCollection returns the mongo collection for the accessTokens
func (m *Manager) getRefreshTokenCollection() *mgo.Collection {
	return db.GetCollection(m.session, refreshTokenCollectionName)
}

// Get an authorizationRequest by it's authorizationcode.
func (m *Manager) getAuthorizationRequest(authorizationcode string) (ar *authorizationRequest, err error) {
	ar = &authorizationRequest{}

	err = m.getAuthorizationRequestCollection().Find(bson.M{"authorizationcode": authorizationcode}).One(ar)
	if err != nil && err == mgo.ErrNotFound {
		ar = nil
		err = nil
		return
	}
	return
}

// saveAuthorizationRequest stores an authorizationRequest, only adding new authorizationRequests is allowed, updating is not
func (m *Manager) saveAuthorizationRequest(ar *authorizationRequest) (err error) {
	// TODO: Validation!

	err = m.getAuthorizationRequestCollection().Insert(ar)

	return
}

// saveAccessToken stores an accessToken, only adding new accessTokens is allowed, updating is not
func (m *Manager) saveAccessToken(at *AccessToken) (err error) {
	// TODO: Validation!

	err = m.getAccessTokenCollection().Insert(at)

	return
}

//GetAccessToken gets an access token by it's actual token string
// If the token is not found or is expired, nil is returned
func (m *Manager) GetAccessToken(token string) (at *AccessToken, err error) {
	at = &AccessToken{}

	err = m.getAccessTokenCollection().Find(bson.M{"accesstoken": token}).One(at)
	if err != nil && err == mgo.ErrNotFound {
		at = nil
		err = nil
		return
	}
	if err != nil {
		at = nil
		return
	}
	if at.IsExpired() {
		at = nil
		err = nil
	}

	return
}

// RemoveOrganizationScopes removes all user:memberof:globalid scopes from all access tokens
func (m *Manager) RemoveOrganizationScopes(globalID string, username string) error {
	var accessTokens []AccessToken
	qry := bson.M{"username": username}
	err := m.getAccessTokenCollection().Find(qry).All(&accessTokens)
	if err == mgo.ErrNotFound {
		return nil
	}
	if err != nil {
		return err
	}
	for _, accessToken := range accessTokens {
		if !accessToken.IsExpired() {
			memberOfScope := fmt.Sprintf("user:memberof:%s", globalID)
			if strings.Contains(accessToken.Scope, memberOfScope) {
				accessToken.Scope = removeScope(accessToken.Scope, memberOfScope)
				err = m.getAccessTokenCollection().UpdateId(accessToken.ID, accessToken)
				if err != nil && err != mgo.ErrNotFound {
					return err
				}
			}
		}
	}
	return nil
}

func removeScope(scope string, scopeToRemove string) string {
	scopes := []string{}
	split := strings.Split(scope, ",")
	for _, part := range split {
		if part != scopeToRemove {
			scopes = append(scopes, part)
		}
	}
	return strings.Join(scopes, ",")
}

//getRefreshToken gets an refresh token by it's refresh token string
// If the token is not found or is expired, nil is returned
func (m *Manager) getRefreshToken(token string) (rt *refreshToken, err error) {
	rt = &refreshToken{}

	err = m.getRefreshTokenCollection().Find(bson.M{"refreshtoken": token}).One(rt)
	if err == mgo.ErrNotFound {
		rt = nil
		err = nil
	}
	if err != nil {
		rt = nil
	}

	return
}

// saveRefreshToken stores a refreshToken
func (m *Manager) saveRefreshToken(t *refreshToken) (err error) {
	if t == nil {
		return
	}

	_, err = m.getRefreshTokenCollection().Upsert(bson.M{"refreshtoken": t.RefreshToken}, t)

	return
}

//getClientsCollection returns the mongo collection for the clients
func (m *Manager) getClientsCollection() *mgo.Collection {
	return db.GetCollection(m.session, clientsCollectionName)
}

//GetClientLabels returns a list of labels for which there are apikeys registered for a specific client
func (m *Manager) GetClientLabels(clientID string) (labels []string, err error) {
	results := []struct{ Label string }{}
	err = m.getClientsCollection().Find(bson.M{"clientid": clientID}).Select(bson.M{"label": 1}).All(&results)
	labels = make([]string, len(results), len(results))
	for i, value := range results {
		labels[i] = value.Label
	}
	return
}

//CreateClient saves an Oauth2 client
func (m *Manager) CreateClient(client *Oauth2Client) (err error) {

	err = m.getClientsCollection().Insert(client)

	if err != nil && mgo.IsDup(err) {
		err = db.ErrDuplicate
	}
	return
}

//UpdateClient updates the label, callbackurl and clientCredentialsGrantType properties of a client
func (m *Manager) UpdateClient(clientID, oldLabel, newLabel string, callbackURL string, clientcredentialsGrantType bool) (err error) {

	_, err = m.getClientsCollection().UpdateAll(bson.M{"clientid": clientID, "label": oldLabel}, bson.M{"$set": bson.M{"label": newLabel, "callbackurl": callbackURL, "clientcredentialsgranttype": clientcredentialsGrantType}})

	if err != nil && mgo.IsDup(err) {
		err = db.ErrDuplicate
	}
	return
}

//DeleteClient removes a client secret by it's clientID and label
func (m *Manager) DeleteClient(clientID, label string) (err error) {
	_, err = m.getClientsCollection().RemoveAll(bson.M{"clientid": clientID, "label": label})
	return
}

//DeleteAllForOrganization removes al client secrets for the organization
func (m *Manager) DeleteAllForOrganization(clientID string) (err error) {
	_, err = m.getClientsCollection().RemoveAll(bson.M{"clientid": clientID})
	return
}

//GetClient retrieves a client given a clientid and a label
func (m *Manager) GetClient(clientID, label string) (client *Oauth2Client, err error) {
	client = &Oauth2Client{}
	err = m.getClientsCollection().Find(bson.M{"clientid": clientID, "label": label}).One(client)
	if err == mgo.ErrNotFound {
		err = nil
		client = nil
		return
	}
	return
}

//AllByClientID retrieves all clients with a given ID
func (m *Manager) AllByClientID(clientID string) (clients []*Oauth2Client, err error) {
	clients = make([]*Oauth2Client, 0)

	err = m.getClientsCollection().Find(bson.M{"clientid": clientID}).All(&clients)
	return
}

//GetClientByCredentials retrieves a client given a clientid and a secret
func (m *Manager) getClientByCredentials(clientID, secret string) (client *Oauth2Client, err error) {
	client = &Oauth2Client{}
	err = m.getClientsCollection().Find(bson.M{"clientid": clientID, "secret": secret}).One(client)
	if err == mgo.ErrNotFound {
		err = nil
		client = nil
		return
	}
	return
}

//RemoveTokensByGlobalID removes oauth tokens by global id
func (m *Manager) RemoveTokensByGlobalID(globalid string) error {
	_, err := m.getAccessTokenCollection().RemoveAll(bson.M{"globalid": globalid})
	return err
}

//RemoveClientsByID removes oauth clients by client id
func (m *Manager) RemoveClientsByID(clientid string) error {
	_, err := m.getAccessTokenCollection().RemoveAll(bson.M{"clientid": clientid})
	return err
}
