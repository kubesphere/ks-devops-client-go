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

// CreatePipelineReader is a Reader for the CreatePipeline structure.
type CreatePipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/pipelines] CreatePipeline", response, response.Code())
	}
}

// NewCreatePipelineOK creates a CreatePipelineOK with default headers values
func NewCreatePipelineOK() *CreatePipelineOK {
	return &CreatePipelineOK{}
}

/*
CreatePipelineOK describes a response with status code 200, with default header values.

ok
*/
type CreatePipelineOK struct {
	Payload *models.V1alpha3Pipeline
}

// IsSuccess returns true when this create pipeline o k response has a 2xx status code
func (o *CreatePipelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create pipeline o k response has a 3xx status code
func (o *CreatePipelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create pipeline o k response has a 4xx status code
func (o *CreatePipelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create pipeline o k response has a 5xx status code
func (o *CreatePipelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create pipeline o k response a status code equal to that given
func (o *CreatePipelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create pipeline o k response
func (o *CreatePipelineOK) Code() int {
	return 200
}

func (o *CreatePipelineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/pipelines][%d] createPipelineOK %s", 200, payload)
}

func (o *CreatePipelineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/pipelines][%d] createPipelineOK %s", 200, payload)
}

func (o *CreatePipelineOK) GetPayload() *models.V1alpha3Pipeline {
	return o.Payload
}

func (o *CreatePipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha3Pipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
