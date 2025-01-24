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

// ListPipelineRunsV1alpha2Reader is a Reader for the ListPipelineRunsV1alpha2 structure.
type ListPipelineRunsV1alpha2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPipelineRunsV1alpha2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPipelineRunsV1alpha2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/runs] ListPipelineRunsV1alpha2", response, response.Code())
	}
}

// NewListPipelineRunsV1alpha2OK creates a ListPipelineRunsV1alpha2OK with default headers values
func NewListPipelineRunsV1alpha2OK() *ListPipelineRunsV1alpha2OK {
	return &ListPipelineRunsV1alpha2OK{}
}

/*
ListPipelineRunsV1alpha2OK describes a response with status code 200, with default header values.

ok
*/
type ListPipelineRunsV1alpha2OK struct {
	Payload *models.DevopsPipelineRunList
}

// IsSuccess returns true when this list pipeline runs v1alpha2 o k response has a 2xx status code
func (o *ListPipelineRunsV1alpha2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list pipeline runs v1alpha2 o k response has a 3xx status code
func (o *ListPipelineRunsV1alpha2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list pipeline runs v1alpha2 o k response has a 4xx status code
func (o *ListPipelineRunsV1alpha2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list pipeline runs v1alpha2 o k response has a 5xx status code
func (o *ListPipelineRunsV1alpha2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this list pipeline runs v1alpha2 o k response a status code equal to that given
func (o *ListPipelineRunsV1alpha2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list pipeline runs v1alpha2 o k response
func (o *ListPipelineRunsV1alpha2OK) Code() int {
	return 200
}

func (o *ListPipelineRunsV1alpha2OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/runs][%d] listPipelineRunsV1alpha2OK %s", 200, payload)
}

func (o *ListPipelineRunsV1alpha2OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/runs][%d] listPipelineRunsV1alpha2OK %s", 200, payload)
}

func (o *ListPipelineRunsV1alpha2OK) GetPayload() *models.DevopsPipelineRunList {
	return o.Payload
}

func (o *ListPipelineRunsV1alpha2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevopsPipelineRunList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
