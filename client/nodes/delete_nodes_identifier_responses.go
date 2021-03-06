package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// DeleteNodesIdentifierReader is a Reader for the DeleteNodesIdentifier structure.
type DeleteNodesIdentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteNodesIdentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteNodesIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteNodesIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteNodesIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteNodesIdentifierOK creates a DeleteNodesIdentifierOK with default headers values
func NewDeleteNodesIdentifierOK() *DeleteNodesIdentifierOK {
	return &DeleteNodesIdentifierOK{}
}

/*DeleteNodesIdentifierOK handles this case with default header values.

Delete successful
*/
type DeleteNodesIdentifierOK struct {
}

func (o *DeleteNodesIdentifierOK) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{identifier}][%d] deleteNodesIdentifierOK ", 200)
}

func (o *DeleteNodesIdentifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNodesIdentifierNotFound creates a DeleteNodesIdentifierNotFound with default headers values
func NewDeleteNodesIdentifierNotFound() *DeleteNodesIdentifierNotFound {
	return &DeleteNodesIdentifierNotFound{}
}

/*DeleteNodesIdentifierNotFound handles this case with default header values.

The node with the identifier was not found.
*/
type DeleteNodesIdentifierNotFound struct {
	Payload *models.Error
}

func (o *DeleteNodesIdentifierNotFound) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{identifier}][%d] deleteNodesIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *DeleteNodesIdentifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodesIdentifierDefault creates a DeleteNodesIdentifierDefault with default headers values
func NewDeleteNodesIdentifierDefault(code int) *DeleteNodesIdentifierDefault {
	return &DeleteNodesIdentifierDefault{
		_statusCode: code,
	}
}

/*DeleteNodesIdentifierDefault handles this case with default header values.

Unexpected error
*/
type DeleteNodesIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete nodes identifier default response
func (o *DeleteNodesIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNodesIdentifierDefault) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{identifier}][%d] DeleteNodesIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNodesIdentifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
