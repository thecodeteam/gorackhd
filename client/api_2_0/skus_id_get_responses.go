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

// SkusIDGetReader is a Reader for the SkusIDGet structure.
type SkusIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *SkusIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSkusIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewSkusIDGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSkusIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewSkusIDGetOK creates a SkusIDGetOK with default headers values
func NewSkusIDGetOK() *SkusIDGetOK {
	return &SkusIDGetOK{}
}

/*SkusIDGetOK handles this case with default header values.

Successfull retrieved the specified SKU
*/
type SkusIDGetOK struct {
	Payload SkusIDGetOKBodyBody
}

func (o *SkusIDGetOK) Error() string {
	return fmt.Sprintf("[GET /skus/{identifier}][%d] skusIdGetOK  %+v", 200, o.Payload)
}

func (o *SkusIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusIDGetNotFound creates a SkusIDGetNotFound with default headers values
func NewSkusIDGetNotFound() *SkusIDGetNotFound {
	return &SkusIDGetNotFound{}
}

/*SkusIDGetNotFound handles this case with default header values.

The SKU with the specified identifier was not found
*/
type SkusIDGetNotFound struct {
	Payload *models.Error
}

func (o *SkusIDGetNotFound) Error() string {
	return fmt.Sprintf("[GET /skus/{identifier}][%d] skusIdGetNotFound  %+v", 404, o.Payload)
}

func (o *SkusIDGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusIDGetDefault creates a SkusIDGetDefault with default headers values
func NewSkusIDGetDefault(code int) *SkusIDGetDefault {
	return &SkusIDGetDefault{
		_statusCode: code,
	}
}

/*SkusIDGetDefault handles this case with default header values.

Unexpected error
*/
type SkusIDGetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the skus Id get default response
func (o *SkusIDGetDefault) Code() int {
	return o._statusCode
}

func (o *SkusIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /skus/{identifier}][%d] skusIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *SkusIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SkusIDGetOKBodyBody skus ID get o k body body

swagger:model SkusIDGetOKBodyBody
*/
type SkusIDGetOKBodyBody interface{}
