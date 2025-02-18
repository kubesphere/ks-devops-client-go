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

// StopBranchPipelineReader is a Reader for the StopBranchPipeline structure.
type StopBranchPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopBranchPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopBranchPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/stop] StopBranchPipeline", response, response.Code())
	}
}

// NewStopBranchPipelineOK creates a StopBranchPipelineOK with default headers values
func NewStopBranchPipelineOK() *StopBranchPipelineOK {
	return &StopBranchPipelineOK{}
}

/*
StopBranchPipelineOK describes a response with status code 200, with default header values.

ok
*/
type StopBranchPipelineOK struct {
	Payload *models.DevopsStopPipeline
}

// IsSuccess returns true when this stop branch pipeline o k response has a 2xx status code
func (o *StopBranchPipelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop branch pipeline o k response has a 3xx status code
func (o *StopBranchPipelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop branch pipeline o k response has a 4xx status code
func (o *StopBranchPipelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop branch pipeline o k response has a 5xx status code
func (o *StopBranchPipelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop branch pipeline o k response a status code equal to that given
func (o *StopBranchPipelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop branch pipeline o k response
func (o *StopBranchPipelineOK) Code() int {
	return 200
}

func (o *StopBranchPipelineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/stop][%d] stopBranchPipelineOK %s", 200, payload)
}

func (o *StopBranchPipelineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/stop][%d] stopBranchPipelineOK %s", 200, payload)
}

func (o *StopBranchPipelineOK) GetPayload() *models.DevopsStopPipeline {
	return o.Payload
}

func (o *StopBranchPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevopsStopPipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
