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

// GetFileFromBranchReader is a Reader for the GetFileFromBranch structure.
type GetFileFromBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileFromBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileFromBranchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/files/{file}] GetFileFromBranch", response, response.Code())
	}
}

// NewGetFileFromBranchOK creates a GetFileFromBranchOK with default headers values
func NewGetFileFromBranchOK() *GetFileFromBranchOK {
	return &GetFileFromBranchOK{}
}

/*
GetFileFromBranchOK describes a response with status code 200, with default header values.

ok
*/
type GetFileFromBranchOK struct {
	Payload *models.GitopsFileInfo
}

// IsSuccess returns true when this get file from branch o k response has a 2xx status code
func (o *GetFileFromBranchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file from branch o k response has a 3xx status code
func (o *GetFileFromBranchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file from branch o k response has a 4xx status code
func (o *GetFileFromBranchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file from branch o k response has a 5xx status code
func (o *GetFileFromBranchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file from branch o k response a status code equal to that given
func (o *GetFileFromBranchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file from branch o k response
func (o *GetFileFromBranchOK) Code() int {
	return 200
}

func (o *GetFileFromBranchOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/files/{file}][%d] getFileFromBranchOK %s", 200, payload)
}

func (o *GetFileFromBranchOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/files/{file}][%d] getFileFromBranchOK %s", 200, payload)
}

func (o *GetFileFromBranchOK) GetPayload() *models.GitopsFileInfo {
	return o.Payload
}

func (o *GetFileFromBranchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitopsFileInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
