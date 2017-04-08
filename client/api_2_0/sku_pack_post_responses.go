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

// SkuPackPostReader is a Reader for the SkuPackPost structure.
type SkuPackPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *SkuPackPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewSkuPackPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewSkuPackPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewSkuPackPostCreated creates a SkuPackPostCreated with default headers values
func NewSkuPackPostCreated() *SkuPackPostCreated {
	return &SkuPackPostCreated{}
}

/*SkuPackPostCreated handles this case with default header values.

Successfully created the SKU Pack
*/
type SkuPackPostCreated struct {
	Payload SkuPackPostCreatedBodyBody
}

func (o *SkuPackPostCreated) Error() string {
	return fmt.Sprintf("[POST /skus/pack][%d] skuPackPostCreated  %+v", 201, o.Payload)
}

func (o *SkuPackPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSkuPackPostDefault creates a SkuPackPostDefault with default headers values
func NewSkuPackPostDefault(code int) *SkuPackPostDefault {
	return &SkuPackPostDefault{
		_statusCode: code,
	}
}

/*SkuPackPostDefault handles this case with default header values.

Unexpected error
*/
type SkuPackPostDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the sku pack post default response
func (o *SkuPackPostDefault) Code() int {
	return o._statusCode
}

func (o *SkuPackPostDefault) Error() string {
	return fmt.Sprintf("[POST /skus/pack][%d] skuPackPost default  %+v", o._statusCode, o.Payload)
}

func (o *SkuPackPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SkuPackPostCreatedBodyBody sku pack post created body body

swagger:model SkuPackPostCreatedBodyBody
*/
type SkuPackPostCreatedBodyBody interface{}
