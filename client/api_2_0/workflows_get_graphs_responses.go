package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// WorkflowsGetGraphsReader is a Reader for the WorkflowsGetGraphs structure.
type WorkflowsGetGraphsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *WorkflowsGetGraphsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWorkflowsGetGraphsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewWorkflowsGetGraphsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewWorkflowsGetGraphsOK creates a WorkflowsGetGraphsOK with default headers values
func NewWorkflowsGetGraphsOK() *WorkflowsGetGraphsOK {
	return &WorkflowsGetGraphsOK{}
}

/*WorkflowsGetGraphsOK handles this case with default header values.

Successfully retrieved all workflow graphs
*/
type WorkflowsGetGraphsOK struct {
	Payload WorkflowsGetGraphsOKBodyBody
}

func (o *WorkflowsGetGraphsOK) Error() string {
	return fmt.Sprintf("[GET /workflows/graphs][%d] workflowsGetGraphsOK  %+v", 200, o.Payload)
}

func (o *WorkflowsGetGraphsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowsGetGraphsDefault creates a WorkflowsGetGraphsDefault with default headers values
func NewWorkflowsGetGraphsDefault(code int) *WorkflowsGetGraphsDefault {
	return &WorkflowsGetGraphsDefault{
		_statusCode: code,
	}
}

/*WorkflowsGetGraphsDefault handles this case with default header values.

Unexpected error
*/
type WorkflowsGetGraphsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the workflows get graphs default response
func (o *WorkflowsGetGraphsDefault) Code() int {
	return o._statusCode
}

func (o *WorkflowsGetGraphsDefault) Error() string {
	return fmt.Sprintf("[GET /workflows/graphs][%d] workflowsGetGraphs default  %+v", o._statusCode, o.Payload)
}

func (o *WorkflowsGetGraphsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WorkflowsGetGraphsOKBodyBody workflows get graphs o k body body

swagger:model WorkflowsGetGraphsOKBodyBody
*/
type WorkflowsGetGraphsOKBodyBody interface{}
