package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Person person
// swagger:model person
type Person struct {

	// age
	Age int64 `json:"age,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this person
func (m *Person) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
