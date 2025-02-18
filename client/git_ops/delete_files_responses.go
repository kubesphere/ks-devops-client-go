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

// DeleteFilesReader is a Reader for the DeleteFiles structure.
type DeleteFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/files] DeleteFiles", response, response.Code())
	}
}

// NewDeleteFilesOK creates a DeleteFilesOK with default headers values
func NewDeleteFilesOK() *DeleteFilesOK {
	return &DeleteFilesOK{}
}

/*
DeleteFilesOK describes a response with status code 200, with default header values.

ok
*/
type DeleteFilesOK struct {
	Payload *models.GitopsDeleteFilesOutput
}

// IsSuccess returns true when this delete files o k response has a 2xx status code
func (o *DeleteFilesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete files o k response has a 3xx status code
func (o *DeleteFilesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete files o k response has a 4xx status code
func (o *DeleteFilesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete files o k response has a 5xx status code
func (o *DeleteFilesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete files o k response a status code equal to that given
func (o *DeleteFilesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete files o k response
func (o *DeleteFilesOK) Code() int {
	return 200
}

func (o *DeleteFilesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/files][%d] deleteFilesOK %s", 200, payload)
}

func (o *DeleteFilesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}/branches/{branch}/files][%d] deleteFilesOK %s", 200, payload)
}

func (o *DeleteFilesOK) GetPayload() *models.GitopsDeleteFilesOutput {
	return o.Payload
}

func (o *DeleteFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GitopsDeleteFilesOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
