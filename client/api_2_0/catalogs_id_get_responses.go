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

// CatalogsIDGetReader is a Reader for the CatalogsIDGet structure.
type CatalogsIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CatalogsIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCatalogsIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCatalogsIDGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewCatalogsIDGetOK creates a CatalogsIDGetOK with default headers values
func NewCatalogsIDGetOK() *CatalogsIDGetOK {
	return &CatalogsIDGetOK{}
}

/*CatalogsIDGetOK handles this case with default header values.

Successfully retrieved the catalog
*/
type CatalogsIDGetOK struct {
	Payload *models.VersionsResponse
}

func (o *CatalogsIDGetOK) Error() string {
	return fmt.Sprintf("[GET /catalogs/{identifier}][%d] catalogsIdGetOK  %+v", 200, o.Payload)
}

func (o *CatalogsIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogsIDGetDefault creates a CatalogsIDGetDefault with default headers values
func NewCatalogsIDGetDefault(code int) *CatalogsIDGetDefault {
	return &CatalogsIDGetDefault{
		_statusCode: code,
	}
}

/*CatalogsIDGetDefault handles this case with default header values.

Error
*/
type CatalogsIDGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the catalogs Id get default response
func (o *CatalogsIDGetDefault) Code() int {
	return o._statusCode
}

func (o *CatalogsIDGetDefault) Error() string {
	return fmt.Sprintf("[GET /catalogs/{identifier}][%d] catalogsIdGet default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogsIDGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
