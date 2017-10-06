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

// NodesGetByIDReader is a Reader for the NodesGetByID structure.
type NodesGetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodesGetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNodesGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewNodesGetByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNodesGetByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodesGetByIDOK creates a NodesGetByIDOK with default headers values
func NewNodesGetByIDOK() *NodesGetByIDOK {
	return &NodesGetByIDOK{}
}

/*NodesGetByIDOK handles this case with default header values.

Successfully retrieved the specified node
*/
type NodesGetByIDOK struct {
	Payload models.NodesGetByIDOKBody
}

func (o *NodesGetByIDOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}][%d] nodesGetByIdOK  %+v", 200, o.Payload)
}

func (o *NodesGetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesGetByIDNotFound creates a NodesGetByIDNotFound with default headers values
func NewNodesGetByIDNotFound() *NodesGetByIDNotFound {
	return &NodesGetByIDNotFound{}
}

/*NodesGetByIDNotFound handles this case with default header values.

The specified node was not found
*/
type NodesGetByIDNotFound struct {
	Payload *models.Error
}

func (o *NodesGetByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}][%d] nodesGetByIdNotFound  %+v", 404, o.Payload)
}

func (o *NodesGetByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesGetByIDDefault creates a NodesGetByIDDefault with default headers values
func NewNodesGetByIDDefault(code int) *NodesGetByIDDefault {
	return &NodesGetByIDDefault{
		_statusCode: code,
	}
}

/*NodesGetByIDDefault handles this case with default header values.

Unexpected error
*/
type NodesGetByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the nodes get by Id default response
func (o *NodesGetByIDDefault) Code() int {
	return o._statusCode
}

func (o *NodesGetByIDDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}][%d] nodesGetById default  %+v", o._statusCode, o.Payload)
}

func (o *NodesGetByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}