package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
)

// Mydate this is mydate
// swagger:model mydate
type Mydate strfmt.Date

// Validate validates this mydate
func (m Mydate) Validate(formats strfmt.Registry) error {
	return nil
}
