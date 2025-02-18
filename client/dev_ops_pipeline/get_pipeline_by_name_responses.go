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

// GetPipelineByNameReader is a Reader for the GetPipelineByName structure.
type GetPipelineByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/pipelines/{pipeline}] getPipelineByName", response, response.Code())
	}
}

// NewGetPipelineByNameOK creates a GetPipelineByNameOK with default headers values
func NewGetPipelineByNameOK() *GetPipelineByNameOK {
	return &GetPipelineByNameOK{}
}

/*
GetPipelineByNameOK describes a response with status code 200, with default header values.

ok
*/
type GetPipelineByNameOK struct {
	Payload *models.V1alpha3Pipeline
}

// IsSuccess returns true when this get pipeline by name o k response has a 2xx status code
func (o *GetPipelineByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pipeline by name o k response has a 3xx status code
func (o *GetPipelineByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline by name o k response has a 4xx status code
func (o *GetPipelineByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipeline by name o k response has a 5xx status code
func (o *GetPipelineByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline by name o k response a status code equal to that given
func (o *GetPipelineByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get pipeline by name o k response
func (o *GetPipelineByNameOK) Code() int {
	return 200
}

func (o *GetPipelineByNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/pipelines/{pipeline}][%d] getPipelineByNameOK %s", 200, payload)
}

func (o *GetPipelineByNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/pipelines/{pipeline}][%d] getPipelineByNameOK %s", 200, payload)
}

func (o *GetPipelineByNameOK) GetPayload() *models.V1alpha3Pipeline {
	return o.Payload
}

func (o *GetPipelineByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha3Pipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
