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
)

// ScanBranchReader is a Reader for the ScanBranch structure.
type ScanBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScanBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScanBranchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/scan] ScanBranch", response, response.Code())
	}
}

// NewScanBranchOK creates a ScanBranchOK with default headers values
func NewScanBranchOK() *ScanBranchOK {
	return &ScanBranchOK{}
}

/*
ScanBranchOK describes a response with status code 200, with default header values.

ok
*/
type ScanBranchOK struct {
	Payload []int64
}

// IsSuccess returns true when this scan branch o k response has a 2xx status code
func (o *ScanBranchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this scan branch o k response has a 3xx status code
func (o *ScanBranchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this scan branch o k response has a 4xx status code
func (o *ScanBranchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this scan branch o k response has a 5xx status code
func (o *ScanBranchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this scan branch o k response a status code equal to that given
func (o *ScanBranchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the scan branch o k response
func (o *ScanBranchOK) Code() int {
	return 200
}

func (o *ScanBranchOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/scan][%d] scanBranchOK %s", 200, payload)
}

func (o *ScanBranchOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/scan][%d] scanBranchOK %s", 200, payload)
}

func (o *ScanBranchOK) GetPayload() []int64 {
	return o.Payload
}

func (o *ScanBranchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
