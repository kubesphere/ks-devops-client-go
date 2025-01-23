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

	"kubesphere.io/devops-client/models"
)

// GetNodesDetailReader is a Reader for the GetNodesDetail structure.
type GetNodesDetailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodesDetailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodesDetailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/runs/{run}/nodesdetail] GetNodesDetail", response, response.Code())
	}
}

// NewGetNodesDetailOK creates a GetNodesDetailOK with default headers values
func NewGetNodesDetailOK() *GetNodesDetailOK {
	return &GetNodesDetailOK{}
}

/*
GetNodesDetailOK describes a response with status code 200, with default header values.

ok
*/
type GetNodesDetailOK struct {
	Payload []*models.DevopsNodesDetail
}

// IsSuccess returns true when this get nodes detail o k response has a 2xx status code
func (o *GetNodesDetailOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get nodes detail o k response has a 3xx status code
func (o *GetNodesDetailOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nodes detail o k response has a 4xx status code
func (o *GetNodesDetailOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nodes detail o k response has a 5xx status code
func (o *GetNodesDetailOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get nodes detail o k response a status code equal to that given
func (o *GetNodesDetailOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get nodes detail o k response
func (o *GetNodesDetailOK) Code() int {
	return 200
}

func (o *GetNodesDetailOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/runs/{run}/nodesdetail][%d] getNodesDetailOK %s", 200, payload)
}

func (o *GetNodesDetailOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kapis/devops.kubesphere.io/v1alpha2/namespaces/{devops}/pipelines/{pipeline}/runs/{run}/nodesdetail][%d] getNodesDetailOK %s", 200, payload)
}

func (o *GetNodesDetailOK) GetPayload() []*models.DevopsNodesDetail {
	return o.Payload
}

func (o *GetNodesDetailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
