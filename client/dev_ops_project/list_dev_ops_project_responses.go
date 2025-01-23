// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_project

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

// ListDevOpsProjectReader is a Reader for the ListDevOpsProject structure.
type ListDevOpsProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDevOpsProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDevOpsProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha3/workspaces/{workspace}/namespaces] ListDevOpsProject", response, response.Code())
	}
}

// NewListDevOpsProjectOK creates a ListDevOpsProjectOK with default headers values
func NewListDevOpsProjectOK() *ListDevOpsProjectOK {
	return &ListDevOpsProjectOK{}
}

/*
ListDevOpsProjectOK describes a response with status code 200, with default header values.

ok
*/
type ListDevOpsProjectOK struct {
	Payload *models.APIListResult
}

// IsSuccess returns true when this list dev ops project o k response has a 2xx status code
func (o *ListDevOpsProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list dev ops project o k response has a 3xx status code
func (o *ListDevOpsProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list dev ops project o k response has a 4xx status code
func (o *ListDevOpsProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list dev ops project o k response has a 5xx status code
func (o *ListDevOpsProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list dev ops project o k response a status code equal to that given
func (o *ListDevOpsProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list dev ops project o k response
func (o *ListDevOpsProjectOK) Code() int {
	return 200
}

func (o *ListDevOpsProjectOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/workspaces/{workspace}/namespaces][%d] listDevOpsProjectOK %s", 200, payload)
}

func (o *ListDevOpsProjectOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/workspaces/{workspace}/namespaces][%d] listDevOpsProjectOK %s", 200, payload)
}

func (o *ListDevOpsProjectOK) GetPayload() *models.APIListResult {
	return o.Payload
}

func (o *ListDevOpsProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
