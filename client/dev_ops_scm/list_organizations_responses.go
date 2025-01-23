// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_scm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"kubesphere.io/devops-client/models"
)

// ListOrganizationsReader is a Reader for the ListOrganizations structure.
type ListOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha3/scms/{scm}/organizations] listOrganizations", response, response.Code())
	}
}

// NewListOrganizationsOK creates a ListOrganizationsOK with default headers values
func NewListOrganizationsOK() *ListOrganizationsOK {
	return &ListOrganizationsOK{}
}

/*
ListOrganizationsOK describes a response with status code 200, with default header values.

ok
*/
type ListOrganizationsOK struct {
	Payload []*models.ScmOrganization
}

// IsSuccess returns true when this list organizations o k response has a 2xx status code
func (o *ListOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list organizations o k response has a 3xx status code
func (o *ListOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organizations o k response has a 4xx status code
func (o *ListOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list organizations o k response has a 5xx status code
func (o *ListOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list organizations o k response a status code equal to that given
func (o *ListOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list organizations o k response
func (o *ListOrganizationsOK) Code() int {
	return 200
}

func (o *ListOrganizationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/scms/{scm}/organizations][%d] listOrganizationsOK %s", 200, payload)
}

func (o *ListOrganizationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/scms/{scm}/organizations][%d] listOrganizationsOK %s", 200, payload)
}

func (o *ListOrganizationsOK) GetPayload() []*models.ScmOrganization {
	return o.Payload
}

func (o *ListOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
