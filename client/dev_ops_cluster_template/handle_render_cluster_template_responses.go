// Code generated by go-swagger; DO NOT EDIT.

package dev_ops_cluster_template

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

// HandleRenderClusterTemplateReader is a Reader for the HandleRenderClusterTemplate structure.
type HandleRenderClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleRenderClusterTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandleRenderClusterTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /kapis/devops.kubesphere.io/v1alpha3/clustertemplates/{clustertemplate}/render] handleRenderClusterTemplate", response, response.Code())
	}
}

// NewHandleRenderClusterTemplateOK creates a HandleRenderClusterTemplateOK with default headers values
func NewHandleRenderClusterTemplateOK() *HandleRenderClusterTemplateOK {
	return &HandleRenderClusterTemplateOK{}
}

/*
HandleRenderClusterTemplateOK describes a response with status code 200, with default header values.

ok
*/
type HandleRenderClusterTemplateOK struct {
	Payload *models.V1alpha3ClusterTemplate
}

// IsSuccess returns true when this handle render cluster template o k response has a 2xx status code
func (o *HandleRenderClusterTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this handle render cluster template o k response has a 3xx status code
func (o *HandleRenderClusterTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this handle render cluster template o k response has a 4xx status code
func (o *HandleRenderClusterTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this handle render cluster template o k response has a 5xx status code
func (o *HandleRenderClusterTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this handle render cluster template o k response a status code equal to that given
func (o *HandleRenderClusterTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the handle render cluster template o k response
func (o *HandleRenderClusterTemplateOK) Code() int {
	return 200
}

func (o *HandleRenderClusterTemplateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha3/clustertemplates/{clustertemplate}/render][%d] handleRenderClusterTemplateOK %s", 200, payload)
}

func (o *HandleRenderClusterTemplateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kapis/devops.kubesphere.io/v1alpha3/clustertemplates/{clustertemplate}/render][%d] handleRenderClusterTemplateOK %s", 200, payload)
}

func (o *HandleRenderClusterTemplateOK) GetPayload() *models.V1alpha3ClusterTemplate {
	return o.Payload
}

func (o *HandleRenderClusterTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1alpha3ClusterTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
