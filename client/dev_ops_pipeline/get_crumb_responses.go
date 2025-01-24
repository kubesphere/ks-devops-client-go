// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_pipeline

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

// GetCrumbReader is a Reader for the GetCrumb structure.
type GetCrumbReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCrumbReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCrumbOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha2/crumbissuer] GetCrumb", response, response.Code())
	}
}

// NewGetCrumbOK creates a GetCrumbOK with default headers values
func NewGetCrumbOK() *GetCrumbOK {
	return &GetCrumbOK{}
}

/*
GetCrumbOK describes a response with status code 200, with default header values.

ok
*/
type GetCrumbOK struct {
	Payload *models.DevopsCrumb
}

// IsSuccess returns true when this get crumb o k response has a 2xx status code
func (o *GetCrumbOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get crumb o k response has a 3xx status code
func (o *GetCrumbOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get crumb o k response has a 4xx status code
func (o *GetCrumbOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get crumb o k response has a 5xx status code
func (o *GetCrumbOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get crumb o k response a status code equal to that given
func (o *GetCrumbOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get crumb o k response
func (o *GetCrumbOK) Code() int {
	return 200
}

func (o *GetCrumbOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/crumbissuer][%d] getCrumbOK %s", 200, payload)
}

func (o *GetCrumbOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/crumbissuer][%d] getCrumbOK %s", 200, payload)
}

func (o *GetCrumbOK) GetPayload() *models.DevopsCrumb {
	return o.Payload
}

func (o *GetCrumbOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevopsCrumb)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
