// Code generated by go-swagger; DO NOT EDIT.

package workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// WorkflowsGetGraphsByNameReader is a Reader for the WorkflowsGetGraphsByName structure.
type WorkflowsGetGraphsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowsGetGraphsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWorkflowsGetGraphsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewWorkflowsGetGraphsByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWorkflowsGetGraphsByNameOK creates a WorkflowsGetGraphsByNameOK with default headers values
func NewWorkflowsGetGraphsByNameOK() *WorkflowsGetGraphsByNameOK {
	return &WorkflowsGetGraphsByNameOK{}
}

/*WorkflowsGetGraphsByNameOK handles this case with default header values.

Successfully retrieved the workflow graph with the specified injectable name
*/
type WorkflowsGetGraphsByNameOK struct {
	Payload WorkflowsGetGraphsByNameOKBody
}

func (o *WorkflowsGetGraphsByNameOK) Error() string {
	return fmt.Sprintf("[GET /workflows/graphs/{injectableName}][%d] workflowsGetGraphsByNameOK  %+v", 200, o.Payload)
}

func (o *WorkflowsGetGraphsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowsGetGraphsByNameDefault creates a WorkflowsGetGraphsByNameDefault with default headers values
func NewWorkflowsGetGraphsByNameDefault(code int) *WorkflowsGetGraphsByNameDefault {
	return &WorkflowsGetGraphsByNameDefault{
		_statusCode: code,
	}
}

/*WorkflowsGetGraphsByNameDefault handles this case with default header values.

Unexpected error
*/
type WorkflowsGetGraphsByNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the workflows get graphs by name default response
func (o *WorkflowsGetGraphsByNameDefault) Code() int {
	return o._statusCode
}

func (o *WorkflowsGetGraphsByNameDefault) Error() string {
	return fmt.Sprintf("[GET /workflows/graphs/{injectableName}][%d] workflowsGetGraphsByName default  %+v", o._statusCode, o.Payload)
}

func (o *WorkflowsGetGraphsByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WorkflowsGetGraphsByNameOKBody workflows get graphs by name o k body
swagger:model WorkflowsGetGraphsByNameOKBody
*/

type WorkflowsGetGraphsByNameOKBody interface{}