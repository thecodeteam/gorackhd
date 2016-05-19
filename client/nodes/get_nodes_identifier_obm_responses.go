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

// GetNodesIdentifierObmReader is a Reader for the GetNodesIdentifierObm structure.
type GetNodesIdentifierObmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNodesIdentifierObmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesIdentifierObmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNodesIdentifierObmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNodesIdentifierObmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetNodesIdentifierObmOK creates a GetNodesIdentifierObmOK with default headers values
func NewGetNodesIdentifierObmOK() *GetNodesIdentifierObmOK {
	return &GetNodesIdentifierObmOK{}
}

/*GetNodesIdentifierObmOK handles this case with default header values.

obm settings
*/
type GetNodesIdentifierObmOK struct {
	Payload []interface{}
}

func (o *GetNodesIdentifierObmOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/obm][%d] getNodesIdentifierObmOK  %+v", 200, o.Payload)
}

func (o *GetNodesIdentifierObmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierObmNotFound creates a GetNodesIdentifierObmNotFound with default headers values
func NewGetNodesIdentifierObmNotFound() *GetNodesIdentifierObmNotFound {
	return &GetNodesIdentifierObmNotFound{}
}

/*GetNodesIdentifierObmNotFound handles this case with default header values.

The node with the identifier was not found or has no obm settings.

*/
type GetNodesIdentifierObmNotFound struct {
	Payload *models.Error
}

func (o *GetNodesIdentifierObmNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/obm][%d] getNodesIdentifierObmNotFound  %+v", 404, o.Payload)
}

func (o *GetNodesIdentifierObmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierObmDefault creates a GetNodesIdentifierObmDefault with default headers values
func NewGetNodesIdentifierObmDefault(code int) *GetNodesIdentifierObmDefault {
	return &GetNodesIdentifierObmDefault{
		_statusCode: code,
	}
}

/*GetNodesIdentifierObmDefault handles this case with default header values.

Unexpected error
*/
type GetNodesIdentifierObmDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nodes identifier obm default response
func (o *GetNodesIdentifierObmDefault) Code() int {
	return o._statusCode
}

func (o *GetNodesIdentifierObmDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/obm][%d] GetNodesIdentifierObm default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodesIdentifierObmDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
