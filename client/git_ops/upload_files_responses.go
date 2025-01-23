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

	"kubesphere.io/devops-client/models"
)

// UploadFilesReader is a Reader for the UploadFiles structure.
type UploadFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/uploads] UploadFiles", response, response.Code())
	}
}

// NewUploadFilesOK creates a UploadFilesOK with default headers values
func NewUploadFilesOK() *UploadFilesOK {
	return &UploadFilesOK{}
}

/*
UploadFilesOK describes a response with status code 200, with default header values.

ok
*/
type UploadFilesOK struct {
	Payload models.GitopsUploadFilesOutput
}

// IsSuccess returns true when this upload files o k response has a 2xx status code
func (o *UploadFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload files o k response has a 3xx status code
func (o *UploadFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload files o k response has a 4xx status code
func (o *UploadFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload files o k response has a 5xx status code
func (o *UploadFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload files o k response a status code equal to that given
func (o *UploadFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upload files o k response
func (o *UploadFilesOK) Code() int {
	return 200
}

func (o *UploadFilesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/uploads][%d] uploadFilesOK %s", 200, payload)
}

func (o *UploadFilesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/uploads][%d] uploadFilesOK %s", 200, payload)
}

func (o *UploadFilesOK) GetPayload() models.GitopsUploadFilesOutput {
	return o.Payload
}

func (o *UploadFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
