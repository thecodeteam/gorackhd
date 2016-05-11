package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// GetPollersIdentifierReader is a Reader for the GetPollersIdentifier structure.
type GetPollersIdentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPollersIdentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPollersIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetPollersIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPollersIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetPollersIdentifierOK creates a GetPollersIdentifierOK with default headers values
func NewGetPollersIdentifierOK() *GetPollersIdentifierOK {
	return &GetPollersIdentifierOK{}
}

/*GetPollersIdentifierOK handles this case with default header values.

Specifics of the specified poller

*/
type GetPollersIdentifierOK struct {
	Payload GetPollersIdentifierOKBodyBody
}

func (o *GetPollersIdentifierOK) Error() string {
	return fmt.Sprintf("[GET /pollers/{identifier}][%d] getPollersIdentifierOK  %+v", 200, o.Payload)
}

func (o *GetPollersIdentifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPollersIdentifierNotFound creates a GetPollersIdentifierNotFound with default headers values
func NewGetPollersIdentifierNotFound() *GetPollersIdentifierNotFound {
	return &GetPollersIdentifierNotFound{}
}

/*GetPollersIdentifierNotFound handles this case with default header values.

There is no  poller with specified identifier.

*/
type GetPollersIdentifierNotFound struct {
	Payload *models.Error
}

func (o *GetPollersIdentifierNotFound) Error() string {
	return fmt.Sprintf("[GET /pollers/{identifier}][%d] getPollersIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *GetPollersIdentifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPollersIdentifierDefault creates a GetPollersIdentifierDefault with default headers values
func NewGetPollersIdentifierDefault(code int) *GetPollersIdentifierDefault {
	return &GetPollersIdentifierDefault{
		_statusCode: code,
	}
}

/*GetPollersIdentifierDefault handles this case with default header values.

Unexpected error
*/
type GetPollersIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get pollers identifier default response
func (o *GetPollersIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *GetPollersIdentifierDefault) Error() string {
	return fmt.Sprintf("[GET /pollers/{identifier}][%d] GetPollersIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *GetPollersIdentifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetPollersIdentifierOKBodyBody get pollers identifier o k body body

swagger:model GetPollersIdentifierOKBodyBody
*/
type GetPollersIdentifierOKBodyBody interface{}
