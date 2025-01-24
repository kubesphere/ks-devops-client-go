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

// GetBranchPipelineRunReader is a Reader for the GetBranchPipelineRun structure.
type GetBranchPipelineRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBranchPipelineRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBranchPipelineRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}] GetBranchPipelineRun", response, response.Code())
	}
}

// NewGetBranchPipelineRunOK creates a GetBranchPipelineRunOK with default headers values
func NewGetBranchPipelineRunOK() *GetBranchPipelineRunOK {
	return &GetBranchPipelineRunOK{}
}

/*
GetBranchPipelineRunOK describes a response with status code 200, with default header values.

ok
*/
type GetBranchPipelineRunOK struct {
	Payload *models.DevopsPipelineRun
}

// IsSuccess returns true when this get branch pipeline run o k response has a 2xx status code
func (o *GetBranchPipelineRunOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get branch pipeline run o k response has a 3xx status code
func (o *GetBranchPipelineRunOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get branch pipeline run o k response has a 4xx status code
func (o *GetBranchPipelineRunOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get branch pipeline run o k response has a 5xx status code
func (o *GetBranchPipelineRunOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get branch pipeline run o k response a status code equal to that given
func (o *GetBranchPipelineRunOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get branch pipeline run o k response
func (o *GetBranchPipelineRunOK) Code() int {
	return 200
}

func (o *GetBranchPipelineRunOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}][%d] getBranchPipelineRunOK %s", 200, payload)
}

func (o *GetBranchPipelineRunOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}][%d] getBranchPipelineRunOK %s", 200, payload)
}

func (o *GetBranchPipelineRunOK) GetPayload() *models.DevopsPipelineRun {
	return o.Payload
}

func (o *GetBranchPipelineRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevopsPipelineRun)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
