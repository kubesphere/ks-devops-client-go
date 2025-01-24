// Code generated by go-swagger; DO NOT EDIT.

package git_ops

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

// CheckOutBranchReader is a Reader for the CheckOutBranch structure.
type CheckOutBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckOutBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckOutBranchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/checkouts] CheckOutBranch", response, response.Code())
	}
}

// NewCheckOutBranchOK creates a CheckOutBranchOK with default headers values
func NewCheckOutBranchOK() *CheckOutBranchOK {
	return &CheckOutBranchOK{}
}

/*
CheckOutBranchOK describes a response with status code 200, with default header values.

ok
*/
type CheckOutBranchOK struct {
	Payload models.GitopsCheckOutBranchOutput
}

// IsSuccess returns true when this check out branch o k response has a 2xx status code
func (o *CheckOutBranchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this check out branch o k response has a 3xx status code
func (o *CheckOutBranchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check out branch o k response has a 4xx status code
func (o *CheckOutBranchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this check out branch o k response has a 5xx status code
func (o *CheckOutBranchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this check out branch o k response a status code equal to that given
func (o *CheckOutBranchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the check out branch o k response
func (o *CheckOutBranchOK) Code() int {
	return 200
}

func (o *CheckOutBranchOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/checkouts][%d] checkOutBranchOK %s", 200, payload)
}

func (o *CheckOutBranchOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/checkouts][%d] checkOutBranchOK %s", 200, payload)
}

func (o *CheckOutBranchOK) GetPayload() models.GitopsCheckOutBranchOutput {
	return o.Payload
}

func (o *CheckOutBranchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
