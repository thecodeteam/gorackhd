package workflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// PostNodesIdentifierWorkflowsReader is a Reader for the PostNodesIdentifierWorkflows structure.
type PostNodesIdentifierWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostNodesIdentifierWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostNodesIdentifierWorkflowsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostNodesIdentifierWorkflowsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostNodesIdentifierWorkflowsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostNodesIdentifierWorkflowsCreated creates a PostNodesIdentifierWorkflowsCreated with default headers values
func NewPostNodesIdentifierWorkflowsCreated() *PostNodesIdentifierWorkflowsCreated {
	return &PostNodesIdentifierWorkflowsCreated{}
}

/*PostNodesIdentifierWorkflowsCreated handles this case with default header values.

the workflow that was created

*/
type PostNodesIdentifierWorkflowsCreated struct {
	Payload PostNodesIdentifierWorkflowsCreatedBodyBody
}

func (o *PostNodesIdentifierWorkflowsCreated) Error() string {
	return fmt.Sprintf("[POST /nodes/{identifier}/workflows][%d] postNodesIdentifierWorkflowsCreated  %+v", 201, o.Payload)
}

func (o *PostNodesIdentifierWorkflowsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNodesIdentifierWorkflowsNotFound creates a PostNodesIdentifierWorkflowsNotFound with default headers values
func NewPostNodesIdentifierWorkflowsNotFound() *PostNodesIdentifierWorkflowsNotFound {
	return &PostNodesIdentifierWorkflowsNotFound{}
}

/*PostNodesIdentifierWorkflowsNotFound handles this case with default header values.

The node with the identifier was not found

*/
type PostNodesIdentifierWorkflowsNotFound struct {
	Payload *models.Error
}

func (o *PostNodesIdentifierWorkflowsNotFound) Error() string {
	return fmt.Sprintf("[POST /nodes/{identifier}/workflows][%d] postNodesIdentifierWorkflowsNotFound  %+v", 404, o.Payload)
}

func (o *PostNodesIdentifierWorkflowsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNodesIdentifierWorkflowsDefault creates a PostNodesIdentifierWorkflowsDefault with default headers values
func NewPostNodesIdentifierWorkflowsDefault(code int) *PostNodesIdentifierWorkflowsDefault {
	return &PostNodesIdentifierWorkflowsDefault{
		_statusCode: code,
	}
}

/*PostNodesIdentifierWorkflowsDefault handles this case with default header values.

Unexpected error
*/
type PostNodesIdentifierWorkflowsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post nodes identifier workflows default response
func (o *PostNodesIdentifierWorkflowsDefault) Code() int {
	return o._statusCode
}

func (o *PostNodesIdentifierWorkflowsDefault) Error() string {
	return fmt.Sprintf("[POST /nodes/{identifier}/workflows][%d] PostNodesIdentifierWorkflows default  %+v", o._statusCode, o.Payload)
}

func (o *PostNodesIdentifierWorkflowsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostNodesIdentifierWorkflowsCreatedBodyBody post nodes identifier workflows created body body

swagger:model PostNodesIdentifierWorkflowsCreatedBodyBody
*/
type PostNodesIdentifierWorkflowsCreatedBodyBody interface{}
