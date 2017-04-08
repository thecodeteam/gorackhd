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

// TemplatesGetByNameReader is a Reader for the TemplatesGetByName structure.
type TemplatesGetByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *TemplatesGetByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTemplatesGetByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewTemplatesGetByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewTemplatesGetByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewTemplatesGetByNameOK creates a TemplatesGetByNameOK with default headers values
func NewTemplatesGetByNameOK() *TemplatesGetByNameOK {
	return &TemplatesGetByNameOK{}
}

/*TemplatesGetByNameOK handles this case with default header values.

Successfuly retrieved the specified template
*/
type TemplatesGetByNameOK struct {
	Payload TemplatesGetByNameOKBodyBody
}

func (o *TemplatesGetByNameOK) Error() string {
	return fmt.Sprintf("[GET /templates/{name}][%d] templatesGetByNameOK  %+v", 200, o.Payload)
}

func (o *TemplatesGetByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplatesGetByNameNotFound creates a TemplatesGetByNameNotFound with default headers values
func NewTemplatesGetByNameNotFound() *TemplatesGetByNameNotFound {
	return &TemplatesGetByNameNotFound{}
}

/*TemplatesGetByNameNotFound handles this case with default header values.

The template with specified identifier was not found
*/
type TemplatesGetByNameNotFound struct {
	Payload *models.Error
}

func (o *TemplatesGetByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /templates/{name}][%d] templatesGetByNameNotFound  %+v", 404, o.Payload)
}

func (o *TemplatesGetByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplatesGetByNameDefault creates a TemplatesGetByNameDefault with default headers values
func NewTemplatesGetByNameDefault(code int) *TemplatesGetByNameDefault {
	return &TemplatesGetByNameDefault{
		_statusCode: code,
	}
}

/*TemplatesGetByNameDefault handles this case with default header values.

Unexpected error
*/
type TemplatesGetByNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the templates get by name default response
func (o *TemplatesGetByNameDefault) Code() int {
	return o._statusCode
}

func (o *TemplatesGetByNameDefault) Error() string {
	return fmt.Sprintf("[GET /templates/{name}][%d] templatesGetByName default  %+v", o._statusCode, o.Payload)
}

func (o *TemplatesGetByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TemplatesGetByNameOKBodyBody templates get by name o k body body

swagger:model TemplatesGetByNameOKBodyBody
*/
type TemplatesGetByNameOKBodyBody interface{}
