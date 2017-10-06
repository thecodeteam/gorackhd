// Code generated by go-swagger; DO NOT EDIT.

package skus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// SkusIDGetNodesReader is a Reader for the SkusIDGetNodes structure.
type SkusIDGetNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SkusIDGetNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSkusIDGetNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewSkusIDGetNodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewSkusIDGetNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSkusIDGetNodesOK creates a SkusIDGetNodesOK with default headers values
func NewSkusIDGetNodesOK() *SkusIDGetNodesOK {
	return &SkusIDGetNodesOK{}
}

/*SkusIDGetNodesOK handles this case with default header values.

Successfully retrieved the nodes associated with the specified SKU
*/
type SkusIDGetNodesOK struct {
	Payload models.SkusIDGetNodesOKBody
}

func (o *SkusIDGetNodesOK) Error() string {
	return fmt.Sprintf("[GET /skus/{identifier}/nodes][%d] skusIdGetNodesOK  %+v", 200, o.Payload)
}

func (o *SkusIDGetNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusIDGetNodesNotFound creates a SkusIDGetNodesNotFound with default headers values
func NewSkusIDGetNodesNotFound() *SkusIDGetNodesNotFound {
	return &SkusIDGetNodesNotFound{}
}

/*SkusIDGetNodesNotFound handles this case with default header values.

The SKU with the specified identifier was not found
*/
type SkusIDGetNodesNotFound struct {
	Payload *models.Error
}

func (o *SkusIDGetNodesNotFound) Error() string {
	return fmt.Sprintf("[GET /skus/{identifier}/nodes][%d] skusIdGetNodesNotFound  %+v", 404, o.Payload)
}

func (o *SkusIDGetNodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkusIDGetNodesDefault creates a SkusIDGetNodesDefault with default headers values
func NewSkusIDGetNodesDefault(code int) *SkusIDGetNodesDefault {
	return &SkusIDGetNodesDefault{
		_statusCode: code,
	}
}

/*SkusIDGetNodesDefault handles this case with default header values.

Unexpected error
*/
type SkusIDGetNodesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the skus Id get nodes default response
func (o *SkusIDGetNodesDefault) Code() int {
	return o._statusCode
}

func (o *SkusIDGetNodesDefault) Error() string {
	return fmt.Sprintf("[GET /skus/{identifier}/nodes][%d] skusIdGetNodes default  %+v", o._statusCode, o.Payload)
}

func (o *SkusIDGetNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}