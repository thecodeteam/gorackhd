package put

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// DeleteSkusIdentifierReader is a Reader for the DeleteSkusIdentifier structure.
type DeleteSkusIdentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteSkusIdentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSkusIdentifierNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteSkusIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteSkusIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteSkusIdentifierNoContent creates a DeleteSkusIdentifierNoContent with default headers values
func NewDeleteSkusIdentifierNoContent() *DeleteSkusIdentifierNoContent {
	return &DeleteSkusIdentifierNoContent{}
}

/*DeleteSkusIdentifierNoContent handles this case with default header values.

return all skus

*/
type DeleteSkusIdentifierNoContent struct {
	Payload DeleteSkusIdentifierNoContentBodyBody
}

func (o *DeleteSkusIdentifierNoContent) Error() string {
	return fmt.Sprintf("[DELETE /skus/{identifier}][%d] deleteSkusIdentifierNoContent  %+v", 204, o.Payload)
}

func (o *DeleteSkusIdentifierNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSkusIdentifierNotFound creates a DeleteSkusIdentifierNotFound with default headers values
func NewDeleteSkusIdentifierNotFound() *DeleteSkusIdentifierNotFound {
	return &DeleteSkusIdentifierNotFound{}
}

/*DeleteSkusIdentifierNotFound handles this case with default header values.

sku with identifier not found, failed.

*/
type DeleteSkusIdentifierNotFound struct {
	Payload *models.Error
}

func (o *DeleteSkusIdentifierNotFound) Error() string {
	return fmt.Sprintf("[DELETE /skus/{identifier}][%d] deleteSkusIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSkusIdentifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSkusIdentifierDefault creates a DeleteSkusIdentifierDefault with default headers values
func NewDeleteSkusIdentifierDefault(code int) *DeleteSkusIdentifierDefault {
	return &DeleteSkusIdentifierDefault{
		_statusCode: code,
	}
}

/*DeleteSkusIdentifierDefault handles this case with default header values.

Unexpected error
*/
type DeleteSkusIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete skus identifier default response
func (o *DeleteSkusIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSkusIdentifierDefault) Error() string {
	return fmt.Sprintf("[DELETE /skus/{identifier}][%d] DeleteSkusIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSkusIdentifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteSkusIdentifierNoContentBodyBody delete skus identifier no content body body

swagger:model DeleteSkusIdentifierNoContentBodyBody
*/
type DeleteSkusIdentifierNoContentBodyBody interface{}
