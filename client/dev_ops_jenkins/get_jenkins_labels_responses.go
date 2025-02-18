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

// GetJenkinsLabelsReader is a Reader for the GetJenkinsLabels structure.
type GetJenkinsLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJenkinsLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJenkinsLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha3/ci/nodelabels] getJenkinsLabels", response, response.Code())
	}
}

// NewGetJenkinsLabelsOK creates a GetJenkinsLabelsOK with default headers values
func NewGetJenkinsLabelsOK() *GetJenkinsLabelsOK {
	return &GetJenkinsLabelsOK{}
}

/*
GetJenkinsLabelsOK describes a response with status code 200, with default header values.

ok
*/
type GetJenkinsLabelsOK struct {
	Payload *models.V1alpha3GenericArrayResponse
}

// IsSuccess returns true when this get jenkins labels o k response has a 2xx status code
func (o *GetJenkinsLabelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get jenkins labels o k response has a 3xx status code
func (o *GetJenkinsLabelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get jenkins labels o k response has a 4xx status code
func (o *GetJenkinsLabelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get jenkins labels o k response has a 5xx status code
func (o *GetJenkinsLabelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get jenkins labels o k response a status code equal to that given
func (o *GetJenkinsLabelsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get jenkins labels o k response
func (o *GetJenkinsLabelsOK) Code() int {
	return 200
}

func (o *GetJenkinsLabelsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/ci/nodelabels][%d] getJenkinsLabelsOK %s", 200, payload)
}

func (o *GetJenkinsLabelsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/ci/nodelabels][%d] getJenkinsLabelsOK %s", 200, payload)
}

func (o *GetJenkinsLabelsOK) GetPayload() *models.V1alpha3GenericArrayResponse {
	return o.Payload
}

func (o *GetJenkinsLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha3GenericArrayResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
