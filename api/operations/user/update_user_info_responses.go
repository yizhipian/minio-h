// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// UpdateUserInfoOKCode is the HTTP code returned for type UpdateUserInfoOK
const UpdateUserInfoOKCode int = 200

/*
UpdateUserInfoOK A successful response.

swagger:response updateUserInfoOK
*/
type UpdateUserInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewUpdateUserInfoOK creates UpdateUserInfoOK with default headers values
func NewUpdateUserInfoOK() *UpdateUserInfoOK {

	return &UpdateUserInfoOK{}
}

// WithPayload adds the payload to the update user info o k response
func (o *UpdateUserInfoOK) WithPayload(payload *models.User) *UpdateUserInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user info o k response
func (o *UpdateUserInfoOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
UpdateUserInfoDefault Generic error response.

swagger:response updateUserInfoDefault
*/
type UpdateUserInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewUpdateUserInfoDefault creates UpdateUserInfoDefault with default headers values
func NewUpdateUserInfoDefault(code int) *UpdateUserInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateUserInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update user info default response
func (o *UpdateUserInfoDefault) WithStatusCode(code int) *UpdateUserInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update user info default response
func (o *UpdateUserInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update user info default response
func (o *UpdateUserInfoDefault) WithPayload(payload *models.APIError) *UpdateUserInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user info default response
func (o *UpdateUserInfoDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateUserInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
