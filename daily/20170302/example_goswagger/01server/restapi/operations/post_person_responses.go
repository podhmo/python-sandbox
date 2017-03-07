package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/foo/booah/models"
)

// HTTP code for type PostPersonCreated
const PostPersonCreatedCode int = 201

/*PostPersonCreated created

swagger:response postPersonCreated
*/
type PostPersonCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Person `json:"body,omitempty"`
}

// NewPostPersonCreated creates PostPersonCreated with default headers values
func NewPostPersonCreated() *PostPersonCreated {
	return &PostPersonCreated{}
}

// WithPayload adds the payload to the post person created response
func (o *PostPersonCreated) WithPayload(payload *models.Person) *PostPersonCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post person created response
func (o *PostPersonCreated) SetPayload(payload *models.Person) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPersonCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
