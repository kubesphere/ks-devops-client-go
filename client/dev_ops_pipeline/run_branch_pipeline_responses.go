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

// RunBranchPipelineReader is a Reader for the RunBranchPipeline structure.
type RunBranchPipelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunBranchPipelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunBranchPipelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/branches/{branch}/runs] RunBranchPipeline", response, response.Code())
	}
}

// NewRunBranchPipelineOK creates a RunBranchPipelineOK with default headers values
func NewRunBranchPipelineOK() *RunBranchPipelineOK {
	return &RunBranchPipelineOK{}
}

/*
RunBranchPipelineOK describes a response with status code 200, with default header values.

ok
*/
type RunBranchPipelineOK struct {
	Payload *models.DevopsRunPipeline
}

// IsSuccess returns true when this run branch pipeline o k response has a 2xx status code
func (o *RunBranchPipelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this run branch pipeline o k response has a 3xx status code
func (o *RunBranchPipelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this run branch pipeline o k response has a 4xx status code
func (o *RunBranchPipelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this run branch pipeline o k response has a 5xx status code
func (o *RunBranchPipelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this run branch pipeline o k response a status code equal to that given
func (o *RunBranchPipelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the run branch pipeline o k response
func (o *RunBranchPipelineOK) Code() int {
	return 200
}

func (o *RunBranchPipelineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/branches/{branch}/runs][%d] runBranchPipelineOK %s", 200, payload)
}

func (o *RunBranchPipelineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/branches/{branch}/runs][%d] runBranchPipelineOK %s", 200, payload)
}

func (o *RunBranchPipelineOK) GetPayload() *models.DevopsRunPipeline {
	return o.Payload
}

func (o *RunBranchPipelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevopsRunPipeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
