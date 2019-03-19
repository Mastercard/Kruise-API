// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "deploy-wizard/gen/models"
)

// CreateDeploymentCreatedCode is the HTTP code returned for type CreateDeploymentCreated
const CreateDeploymentCreatedCode int = 201

/*CreateDeploymentCreated created

swagger:response createDeploymentCreated
*/
type CreateDeploymentCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Application `json:"body,omitempty"`
}

// NewCreateDeploymentCreated creates CreateDeploymentCreated with default headers values
func NewCreateDeploymentCreated() *CreateDeploymentCreated {

	return &CreateDeploymentCreated{}
}

// WithPayload adds the payload to the create deployment created response
func (o *CreateDeploymentCreated) WithPayload(payload *models.Application) *CreateDeploymentCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create deployment created response
func (o *CreateDeploymentCreated) SetPayload(payload *models.Application) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDeploymentCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateDeploymentBadRequestCode is the HTTP code returned for type CreateDeploymentBadRequest
const CreateDeploymentBadRequestCode int = 400

/*CreateDeploymentBadRequest Bad request

swagger:response createDeploymentBadRequest
*/
type CreateDeploymentBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateDeploymentBadRequest creates CreateDeploymentBadRequest with default headers values
func NewCreateDeploymentBadRequest() *CreateDeploymentBadRequest {

	return &CreateDeploymentBadRequest{}
}

// WithPayload adds the payload to the create deployment bad request response
func (o *CreateDeploymentBadRequest) WithPayload(payload string) *CreateDeploymentBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create deployment bad request response
func (o *CreateDeploymentBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDeploymentBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*CreateDeploymentDefault Internal server error

swagger:response createDeploymentDefault
*/
type CreateDeploymentDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCreateDeploymentDefault creates CreateDeploymentDefault with default headers values
func NewCreateDeploymentDefault(code int) *CreateDeploymentDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateDeploymentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create deployment default response
func (o *CreateDeploymentDefault) WithStatusCode(code int) *CreateDeploymentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create deployment default response
func (o *CreateDeploymentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create deployment default response
func (o *CreateDeploymentDefault) WithPayload(payload string) *CreateDeploymentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create deployment default response
func (o *CreateDeploymentDefault) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDeploymentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}