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

// GetBranchesReader is a Reader for the GetBranches structure.
type GetBranchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBranchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBranchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/pipelines/{pipeline}/branches] getBranches", response, response.Code())
	}
}

// NewGetBranchesOK creates a GetBranchesOK with default headers values
func NewGetBranchesOK() *GetBranchesOK {
	return &GetBranchesOK{}
}

/*
GetBranchesOK describes a response with status code 200, with default header values.

ok
*/
type GetBranchesOK struct {
	Payload *models.APIListResult
}

// IsSuccess returns true when this get branches o k response has a 2xx status code
func (o *GetBranchesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get branches o k response has a 3xx status code
func (o *GetBranchesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get branches o k response has a 4xx status code
func (o *GetBranchesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get branches o k response has a 5xx status code
func (o *GetBranchesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get branches o k response a status code equal to that given
func (o *GetBranchesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get branches o k response
func (o *GetBranchesOK) Code() int {
	return 200
}

func (o *GetBranchesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/pipelines/{pipeline}/branches][%d] getBranchesOK %s", 200, payload)
}

func (o *GetBranchesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/pipelines/{pipeline}/branches][%d] getBranchesOK %s", 200, payload)
}

func (o *GetBranchesOK) GetPayload() *models.APIListResult {
	return o.Payload
}

func (o *GetBranchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIListResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
