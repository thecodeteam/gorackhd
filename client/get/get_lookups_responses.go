package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// GetLookupsReader is a Reader for the GetLookups structure.
type GetLookupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetLookupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLookupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetLookupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetLookupsOK creates a GetLookupsOK with default headers values
func NewGetLookupsOK() *GetLookupsOK {
	return &GetLookupsOK{}
}

/*GetLookupsOK handles this case with default header values.

array of all
*/
type GetLookupsOK struct {
	Payload []interface{}
}

func (o *GetLookupsOK) Error() string {
	return fmt.Sprintf("[GET /lookups][%d] getLookupsOK  %+v", 200, o.Payload)
}

func (o *GetLookupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLookupsDefault creates a GetLookupsDefault with default headers values
func NewGetLookupsDefault(code int) *GetLookupsDefault {
	return &GetLookupsDefault{
		_statusCode: code,
	}
}

/*GetLookupsDefault handles this case with default header values.

Unexpected error
*/
type GetLookupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get lookups default response
func (o *GetLookupsDefault) Code() int {
	return o._statusCode
}

func (o *GetLookupsDefault) Error() string {
	return fmt.Sprintf("[GET /lookups][%d] GetLookups default  %+v", o._statusCode, o.Payload)
}

func (o *GetLookupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}