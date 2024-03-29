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

package object

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/minio/console/models"
)

// DeleteObjectOKCode is the HTTP code returned for type DeleteObjectOK
const DeleteObjectOKCode int = 200

/*
DeleteObjectOK A successful response.

swagger:response deleteObjectOK
*/
type DeleteObjectOK struct {
}

// NewDeleteObjectOK creates DeleteObjectOK with default headers values
func NewDeleteObjectOK() *DeleteObjectOK {

	return &DeleteObjectOK{}
}

// WriteResponse to the client
func (o *DeleteObjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*
DeleteObjectDefault Generic error response.

swagger:response deleteObjectDefault
*/
type DeleteObjectDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIError `json:"body,omitempty"`
}

// NewDeleteObjectDefault creates DeleteObjectDefault with default headers values
func NewDeleteObjectDefault(code int) *DeleteObjectDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteObjectDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete object default response
func (o *DeleteObjectDefault) WithStatusCode(code int) *DeleteObjectDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete object default response
func (o *DeleteObjectDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete object default response
func (o *DeleteObjectDefault) WithPayload(payload *models.APIError) *DeleteObjectDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete object default response
func (o *DeleteObjectDefault) SetPayload(payload *models.APIError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteObjectDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
