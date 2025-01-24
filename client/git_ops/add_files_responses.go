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

// AddFilesReader is a Reader for the AddFiles structure.
type AddFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/files] AddFiles", response, response.Code())
	}
}

// NewAddFilesOK creates a AddFilesOK with default headers values
func NewAddFilesOK() *AddFilesOK {
	return &AddFilesOK{}
}

/*
AddFilesOK describes a response with status code 200, with default header values.

ok
*/
type AddFilesOK struct {
	Payload *models.GitopsAddFilesOutput
}

// IsSuccess returns true when this add files o k response has a 2xx status code
func (o *AddFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add files o k response has a 3xx status code
func (o *AddFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add files o k response has a 4xx status code
func (o *AddFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add files o k response has a 5xx status code
func (o *AddFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add files o k response a status code equal to that given
func (o *AddFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add files o k response
func (o *AddFilesOK) Code() int {
	return 200
}

func (o *AddFilesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/files][%d] addFilesOK %s", 200, payload)
}

func (o *AddFilesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/files][%d] addFilesOK %s", 200, payload)
}

func (o *AddFilesOK) GetPayload() *models.GitopsAddFilesOutput {
	return o.Payload
}

func (o *AddFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitopsAddFilesOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
