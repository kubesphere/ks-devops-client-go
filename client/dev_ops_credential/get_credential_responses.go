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

// GetCredentialReader is a Reader for the GetCredential structure.
type GetCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/credentials/{credential}] GetCredential", response, response.Code())
	}
}

// NewGetCredentialOK creates a GetCredentialOK with default headers values
func NewGetCredentialOK() *GetCredentialOK {
	return &GetCredentialOK{}
}

/*
GetCredentialOK describes a response with status code 200, with default header values.

ok
*/
type GetCredentialOK struct {
	Payload *models.V1Secret
}

// IsSuccess returns true when this get credential o k response has a 2xx status code
func (o *GetCredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get credential o k response has a 3xx status code
func (o *GetCredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get credential o k response has a 4xx status code
func (o *GetCredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get credential o k response has a 5xx status code
func (o *GetCredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get credential o k response a status code equal to that given
func (o *GetCredentialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get credential o k response
func (o *GetCredentialOK) Code() int {
	return 200
}

func (o *GetCredentialOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/credentials/{credential}][%d] getCredentialOK %s", 200, payload)
}

func (o *GetCredentialOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/credentials/{credential}][%d] getCredentialOK %s", 200, payload)
}

func (o *GetCredentialOK) GetPayload() *models.V1Secret {
	return o.Payload
}

func (o *GetCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Secret)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
