package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// GetNodesIdentifierReader is a Reader for the GetNodesIdentifier structure.
type GetNodesIdentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNodesIdentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNodesIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNodesIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetNodesIdentifierOK creates a GetNodesIdentifierOK with default headers values
func NewGetNodesIdentifierOK() *GetNodesIdentifierOK {
	return &GetNodesIdentifierOK{}
}

/*GetNodesIdentifierOK handles this case with default header values.

array of all
*/
type GetNodesIdentifierOK struct {
	Payload []interface{}
}

func (o *GetNodesIdentifierOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}][%d] getNodesIdentifierOK  %+v", 200, o.Payload)
}

func (o *GetNodesIdentifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierNotFound creates a GetNodesIdentifierNotFound with default headers values
func NewGetNodesIdentifierNotFound() *GetNodesIdentifierNotFound {
	return &GetNodesIdentifierNotFound{}
}

/*GetNodesIdentifierNotFound handles this case with default header values.

The node with the identifier was not found.
*/
type GetNodesIdentifierNotFound struct {
	Payload *models.Error
}

func (o *GetNodesIdentifierNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}][%d] getNodesIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *GetNodesIdentifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierDefault creates a GetNodesIdentifierDefault with default headers values
func NewGetNodesIdentifierDefault(code int) *GetNodesIdentifierDefault {
	return &GetNodesIdentifierDefault{
		_statusCode: code,
	}
}

/*GetNodesIdentifierDefault handles this case with default header values.

Unexpected error
*/
type GetNodesIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nodes identifier default response
func (o *GetNodesIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *GetNodesIdentifierDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}][%d] GetNodesIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodesIdentifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
