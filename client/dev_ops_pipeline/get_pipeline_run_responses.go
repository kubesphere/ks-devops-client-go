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

	"kubesphere.io/devops-client/models"
)

// GetPipelineRunReader is a Reader for the GetPipelineRun structure.
type GetPipelineRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/pipelineruns/{pipelinerun}] getPipelineRun", response, response.Code())
	}
}

// NewGetPipelineRunOK creates a GetPipelineRunOK with default headers values
func NewGetPipelineRunOK() *GetPipelineRunOK {
	return &GetPipelineRunOK{}
}

/*
GetPipelineRunOK describes a response with status code 200, with default header values.

ok
*/
type GetPipelineRunOK struct {
	Payload *models.V1alpha3PipelineRun
}

// IsSuccess returns true when this get pipeline run o k response has a 2xx status code
func (o *GetPipelineRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pipeline run o k response has a 3xx status code
func (o *GetPipelineRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline run o k response has a 4xx status code
func (o *GetPipelineRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipeline run o k response has a 5xx status code
func (o *GetPipelineRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline run o k response a status code equal to that given
func (o *GetPipelineRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get pipeline run o k response
func (o *GetPipelineRunOK) Code() int {
	return 200
}

func (o *GetPipelineRunOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/pipelineruns/{pipelinerun}][%d] getPipelineRunOK %s", 200, payload)
}

func (o *GetPipelineRunOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/pipelineruns/{pipelinerun}][%d] getPipelineRunOK %s", 200, payload)
}

func (o *GetPipelineRunOK) GetPayload() *models.V1alpha3PipelineRun {
	return o.Payload
}

func (o *GetPipelineRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha3PipelineRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
