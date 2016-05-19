package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// GetWorkflowsTasksReader is a Reader for the GetWorkflowsTasks structure.
type GetWorkflowsTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetWorkflowsTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWorkflowsTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetWorkflowsTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetWorkflowsTasksOK creates a GetWorkflowsTasksOK with default headers values
func NewGetWorkflowsTasksOK() *GetWorkflowsTasksOK {
	return &GetWorkflowsTasksOK{}
}

/*GetWorkflowsTasksOK handles this case with default header values.

Fetch tasks from task library

*/
type GetWorkflowsTasksOK struct {
	Payload []interface{}
}

func (o *GetWorkflowsTasksOK) Error() string {
	return fmt.Sprintf("[GET /workflows/tasks][%d] getWorkflowsTasksOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowsTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowsTasksDefault creates a GetWorkflowsTasksDefault with default headers values
func NewGetWorkflowsTasksDefault(code int) *GetWorkflowsTasksDefault {
	return &GetWorkflowsTasksDefault{
		_statusCode: code,
	}
}

/*GetWorkflowsTasksDefault handles this case with default header values.

Unexpected error
*/
type GetWorkflowsTasksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get workflows tasks default response
func (o *GetWorkflowsTasksDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowsTasksDefault) Error() string {
	return fmt.Sprintf("[GET /workflows/tasks][%d] GetWorkflowsTasks default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowsTasksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
