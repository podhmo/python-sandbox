package main

import (
	"github.com/podhmo/errmap"
	"encoding/json"
)

func (p *Person) UnmarshalJSON(b []byte) error {
	var err *errmap.Error

	// loading internal data
	var inner struct {
		Name *string `json:"name"`// required
	}
	if rawErr := json.Unmarshal(b, &inner); rawErr != nil  {
		return err.addSummary(rawErr.Error())
	}

	// binding field value and required check
	if inner.Name != nil  {
		p.Name = *inner.Name
	} else  {
		err = err.Add("name", "required")
	}

	return err.Untyped()
}
