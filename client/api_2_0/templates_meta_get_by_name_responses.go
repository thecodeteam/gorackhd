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

// TemplatesMetaGetByNameReader is a Reader for the TemplatesMetaGetByName structure.
type TemplatesMetaGetByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *TemplatesMetaGetByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTemplatesMetaGetByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewTemplatesMetaGetByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewTemplatesMetaGetByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewTemplatesMetaGetByNameOK creates a TemplatesMetaGetByNameOK with default headers values
func NewTemplatesMetaGetByNameOK() *TemplatesMetaGetByNameOK {
	return &TemplatesMetaGetByNameOK{}
}

/*TemplatesMetaGetByNameOK handles this case with default header values.

Successfully retrieved the metadata of the specified template.
*/
type TemplatesMetaGetByNameOK struct {
	Payload TemplatesMetaGetByNameOKBodyBody
}

func (o *TemplatesMetaGetByNameOK) Error() string {
	return fmt.Sprintf("[GET /templates/metadata/{name}][%d] templatesMetaGetByNameOK  %+v", 200, o.Payload)
}

func (o *TemplatesMetaGetByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplatesMetaGetByNameNotFound creates a TemplatesMetaGetByNameNotFound with default headers values
func NewTemplatesMetaGetByNameNotFound() *TemplatesMetaGetByNameNotFound {
	return &TemplatesMetaGetByNameNotFound{}
}

/*TemplatesMetaGetByNameNotFound handles this case with default header values.

Template with specified identifier was not found
*/
type TemplatesMetaGetByNameNotFound struct {
	Payload *models.Error
}

func (o *TemplatesMetaGetByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /templates/metadata/{name}][%d] templatesMetaGetByNameNotFound  %+v", 404, o.Payload)
}

func (o *TemplatesMetaGetByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplatesMetaGetByNameDefault creates a TemplatesMetaGetByNameDefault with default headers values
func NewTemplatesMetaGetByNameDefault(code int) *TemplatesMetaGetByNameDefault {
	return &TemplatesMetaGetByNameDefault{
		_statusCode: code,
	}
}

/*TemplatesMetaGetByNameDefault handles this case with default header values.

Unexpected error
*/
type TemplatesMetaGetByNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the templates meta get by name default response
func (o *TemplatesMetaGetByNameDefault) Code() int {
	return o._statusCode
}

func (o *TemplatesMetaGetByNameDefault) Error() string {
	return fmt.Sprintf("[GET /templates/metadata/{name}][%d] templatesMetaGetByName default  %+v", o._statusCode, o.Payload)
}

func (o *TemplatesMetaGetByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TemplatesMetaGetByNameOKBodyBody templates meta get by name o k body body

swagger:model TemplatesMetaGetByNameOKBodyBody
*/
type TemplatesMetaGetByNameOKBodyBody interface{}
