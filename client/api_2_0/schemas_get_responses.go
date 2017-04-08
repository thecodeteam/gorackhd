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

// SchemasGetReader is a Reader for the SchemasGet structure.
type SchemasGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *SchemasGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSchemasGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewSchemasGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewSchemasGetOK creates a SchemasGetOK with default headers values
func NewSchemasGetOK() *SchemasGetOK {
	return &SchemasGetOK{}
}

/*SchemasGetOK handles this case with default header values.

Successfully retrieved the list of schemas
*/
type SchemasGetOK struct {
	Payload SchemasGetOKBodyBody
}

func (o *SchemasGetOK) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] schemasGetOK  %+v", 200, o.Payload)
}

func (o *SchemasGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemasGetDefault creates a SchemasGetDefault with default headers values
func NewSchemasGetDefault(code int) *SchemasGetDefault {
	return &SchemasGetDefault{
		_statusCode: code,
	}
}

/*SchemasGetDefault handles this case with default header values.

Error
*/
type SchemasGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the schemas get default response
func (o *SchemasGetDefault) Code() int {
	return o._statusCode
}

func (o *SchemasGetDefault) Error() string {
	return fmt.Sprintf("[GET /schemas][%d] schemasGet default  %+v", o._statusCode, o.Payload)
}

func (o *SchemasGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SchemasGetOKBodyBody schemas get o k body body

swagger:model SchemasGetOKBodyBody
*/
type SchemasGetOKBodyBody interface{}
