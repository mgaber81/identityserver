package itsyouonline

import (
	"gopkg.in/validator.v2"
)

type DigitalAssetAddress struct {
	Address        string   `json:"address" validate:"nonzero"`
	Currencysymbol string   `json:"currencysymbol" validate:"nonzero"`
	Expire         DateTime `json:"expire" validate:"nonzero"`
	Label          Label    `json:"label" validate:"nonzero"`
	Noexpiration   bool     `json:"noexpiration,omitempty"`
}

func (s DigitalAssetAddress) Validate() error {

	return validator.Validate(s)
}
