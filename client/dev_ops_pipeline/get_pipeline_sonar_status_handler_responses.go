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

// GetPipelineSonarStatusHandlerReader is a Reader for the GetPipelineSonarStatusHandler structure.
type GetPipelineSonarStatusHandlerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPipelineSonarStatusHandlerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPipelineSonarStatusHandlerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/sonarstatus] GetPipelineSonarStatusHandler", response, response.Code())
	}
}

// NewGetPipelineSonarStatusHandlerOK creates a GetPipelineSonarStatusHandlerOK with default headers values
func NewGetPipelineSonarStatusHandlerOK() *GetPipelineSonarStatusHandlerOK {
	return &GetPipelineSonarStatusHandlerOK{}
}

/*
GetPipelineSonarStatusHandlerOK describes a response with status code 200, with default header values.

ok
*/
type GetPipelineSonarStatusHandlerOK struct {
	Payload []*models.SonarqubeSonarStatus
}

// IsSuccess returns true when this get pipeline sonar status handler o k response has a 2xx status code
func (o *GetPipelineSonarStatusHandlerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get pipeline sonar status handler o k response has a 3xx status code
func (o *GetPipelineSonarStatusHandlerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get pipeline sonar status handler o k response has a 4xx status code
func (o *GetPipelineSonarStatusHandlerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get pipeline sonar status handler o k response has a 5xx status code
func (o *GetPipelineSonarStatusHandlerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get pipeline sonar status handler o k response a status code equal to that given
func (o *GetPipelineSonarStatusHandlerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get pipeline sonar status handler o k response
func (o *GetPipelineSonarStatusHandlerOK) Code() int {
	return 200
}

func (o *GetPipelineSonarStatusHandlerOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/sonarstatus][%d] getPipelineSonarStatusHandlerOK %s", 200, payload)
}

func (o *GetPipelineSonarStatusHandlerOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/sonarstatus][%d] getPipelineSonarStatusHandlerOK %s", 200, payload)
}

func (o *GetPipelineSonarStatusHandlerOK) GetPayload() []*models.SonarqubeSonarStatus {
	return o.Payload
}

func (o *GetPipelineSonarStatusHandlerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
