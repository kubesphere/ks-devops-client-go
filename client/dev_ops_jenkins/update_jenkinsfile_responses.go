// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_jenkins

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

// UpdateJenkinsfileReader is a Reader for the UpdateJenkinsfile structure.
type UpdateJenkinsfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateJenkinsfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateJenkinsfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/pipelines/{pipeline}/jenkinsfile] UpdateJenkinsfile", response, response.Code())
	}
}

// NewUpdateJenkinsfileOK creates a UpdateJenkinsfileOK with default headers values
func NewUpdateJenkinsfileOK() *UpdateJenkinsfileOK {
	return &UpdateJenkinsfileOK{}
}

/*
UpdateJenkinsfileOK describes a response with status code 200, with default header values.

ok
*/
type UpdateJenkinsfileOK struct {
	Payload *models.V1alpha3GenericResponse
}

// IsSuccess returns true when this update jenkinsfile o k response has a 2xx status code
func (o *UpdateJenkinsfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update jenkinsfile o k response has a 3xx status code
func (o *UpdateJenkinsfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update jenkinsfile o k response has a 4xx status code
func (o *UpdateJenkinsfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update jenkinsfile o k response has a 5xx status code
func (o *UpdateJenkinsfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update jenkinsfile o k response a status code equal to that given
func (o *UpdateJenkinsfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update jenkinsfile o k response
func (o *UpdateJenkinsfileOK) Code() int {
	return 200
}

func (o *UpdateJenkinsfileOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/pipelines/{pipeline}/jenkinsfile][%d] updateJenkinsfileOK %s", 200, payload)
}

func (o *UpdateJenkinsfileOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /kapis/devops.kubesphere.io/v1alpha3/namespaces/{devops}/pipelines/{pipeline}/jenkinsfile][%d] updateJenkinsfileOK %s", 200, payload)
}

func (o *UpdateJenkinsfileOK) GetPayload() *models.V1alpha3GenericResponse {
	return o.Payload
}

func (o *UpdateJenkinsfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha3GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
