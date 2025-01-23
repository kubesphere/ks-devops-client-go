// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_scm

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

// DeleteGitRepositoriesReader is a Reader for the DeleteGitRepositories structure.
type DeleteGitRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGitRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGitRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}] deleteGitRepositories", response, response.Code())
	}
}

// NewDeleteGitRepositoriesOK creates a DeleteGitRepositoriesOK with default headers values
func NewDeleteGitRepositoriesOK() *DeleteGitRepositoriesOK {
	return &DeleteGitRepositoriesOK{}
}

/*
DeleteGitRepositoriesOK describes a response with status code 200, with default header values.

ok
*/
type DeleteGitRepositoriesOK struct {
	Payload *models.V1alpha3GitRepository
}

// IsSuccess returns true when this delete git repositories o k response has a 2xx status code
func (o *DeleteGitRepositoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete git repositories o k response has a 3xx status code
func (o *DeleteGitRepositoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete git repositories o k response has a 4xx status code
func (o *DeleteGitRepositoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete git repositories o k response has a 5xx status code
func (o *DeleteGitRepositoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete git repositories o k response a status code equal to that given
func (o *DeleteGitRepositoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete git repositories o k response
func (o *DeleteGitRepositoriesOK) Code() int {
	return 200
}

func (o *DeleteGitRepositoriesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}][%d] deleteGitRepositoriesOK %s", 200, payload)
}

func (o *DeleteGitRepositoriesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /kapis/devops.kubesphere.io/v1alpha3/namespaces/{namespace}/gitrepositories/{gitrepository}][%d] deleteGitRepositoriesOK %s", 200, payload)
}

func (o *DeleteGitRepositoriesOK) GetPayload() *models.V1alpha3GitRepository {
	return o.Payload
}

func (o *DeleteGitRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha3GitRepository)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
