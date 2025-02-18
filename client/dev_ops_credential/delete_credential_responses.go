// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubesphere/ks-devops-client-go/models"
)

// DeleteCredentialReader is a Reader for the DeleteCredential structure.
type DeleteCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/credentials/{credential}] DeleteCredential", response, response.Code())
	}
}

// NewDeleteCredentialOK creates a DeleteCredentialOK with default headers values
func NewDeleteCredentialOK() *DeleteCredentialOK {
	return &DeleteCredentialOK{}
}

/*
DeleteCredentialOK describes a response with status code 200, with default header values.

ok
*/
type DeleteCredentialOK struct {
	Payload *models.V1Secret
}

// IsSuccess returns true when this delete credential o k response has a 2xx status code
func (o *DeleteCredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete credential o k response has a 3xx status code
func (o *DeleteCredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete credential o k response has a 4xx status code
func (o *DeleteCredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete credential o k response has a 5xx status code
func (o *DeleteCredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete credential o k response a status code equal to that given
func (o *DeleteCredentialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete credential o k response
func (o *DeleteCredentialOK) Code() int {
	return 200
}

func (o *DeleteCredentialOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/credentials/{credential}][%d] deleteCredentialOK %s", 200, payload)
}

func (o *DeleteCredentialOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/credentials/{credential}][%d] deleteCredentialOK %s", 200, payload)
}

func (o *DeleteCredentialOK) GetPayload() *models.V1Secret {
	return o.Payload
}

func (o *DeleteCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Secret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
