// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// NodesDelRelationsReader is a Reader for the NodesDelRelations structure.
type NodesDelRelationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodesDelRelationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewNodesDelRelationsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewNodesDelRelationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNodesDelRelationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodesDelRelationsNoContent creates a NodesDelRelationsNoContent with default headers values
func NewNodesDelRelationsNoContent() *NodesDelRelationsNoContent {
	return &NodesDelRelationsNoContent{}
}

/*NodesDelRelationsNoContent handles this case with default header values.

Successfully removed the relations
*/
type NodesDelRelationsNoContent struct {
}

func (o *NodesDelRelationsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{identifier}/relations][%d] nodesDelRelationsNoContent ", 204)
}

func (o *NodesDelRelationsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNodesDelRelationsNotFound creates a NodesDelRelationsNotFound with default headers values
func NewNodesDelRelationsNotFound() *NodesDelRelationsNotFound {
	return &NodesDelRelationsNotFound{}
}

/*NodesDelRelationsNotFound handles this case with default header values.

The specified node was not found.
*/
type NodesDelRelationsNotFound struct {
	Payload *models.Error
}

func (o *NodesDelRelationsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{identifier}/relations][%d] nodesDelRelationsNotFound  %+v", 404, o.Payload)
}

func (o *NodesDelRelationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesDelRelationsDefault creates a NodesDelRelationsDefault with default headers values
func NewNodesDelRelationsDefault(code int) *NodesDelRelationsDefault {
	return &NodesDelRelationsDefault{
		_statusCode: code,
	}
}

/*NodesDelRelationsDefault handles this case with default header values.

Unexpected error
*/
type NodesDelRelationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the nodes del relations default response
func (o *NodesDelRelationsDefault) Code() int {
	return o._statusCode
}

func (o *NodesDelRelationsDefault) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{identifier}/relations][%d] nodesDelRelations default  %+v", o._statusCode, o.Payload)
}

func (o *NodesDelRelationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}