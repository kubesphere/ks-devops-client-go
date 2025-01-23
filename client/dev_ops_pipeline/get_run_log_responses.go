// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_pipeline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRunLogReader is a Reader for the GetRunLog structure.
type GetRunLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/runs/{run}/log] GetRunLog", response, response.Code())
	}
}

// NewGetRunLogOK creates a GetRunLogOK with default headers values
func NewGetRunLogOK() *GetRunLogOK {
	return &GetRunLogOK{}
}

/*
GetRunLogOK describes a response with status code 200, with default header values.

OK
*/
type GetRunLogOK struct {
}

// IsSuccess returns true when this get run log o k response has a 2xx status code
func (o *GetRunLogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get run log o k response has a 3xx status code
func (o *GetRunLogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get run log o k response has a 4xx status code
func (o *GetRunLogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get run log o k response has a 5xx status code
func (o *GetRunLogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get run log o k response a status code equal to that given
func (o *GetRunLogOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get run log o k response
func (o *GetRunLogOK) Code() int {
	return 200
}

func (o *GetRunLogOK) Error() string {
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/runs/{run}/log][%d] getRunLogOK", 200)
}

func (o *GetRunLogOK) String() string {
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/runs/{run}/log][%d] getRunLogOK", 200)
}

func (o *GetRunLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
