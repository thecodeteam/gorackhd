package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// GetPollersReader is a Reader for the GetPollers structure.
type GetPollersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPollersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPollersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetPollersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetPollersOK creates a GetPollersOK with default headers values
func NewGetPollersOK() *GetPollersOK {
	return &GetPollersOK{}
}

/*GetPollersOK handles this case with default header values.

list of all pollers

*/
type GetPollersOK struct {
	Payload []interface{}
}

func (o *GetPollersOK) Error() string {
	return fmt.Sprintf("[GET /pollers][%d] getPollersOK  %+v", 200, o.Payload)
}

func (o *GetPollersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPollersDefault creates a GetPollersDefault with default headers values
func NewGetPollersDefault(code int) *GetPollersDefault {
	return &GetPollersDefault{
		_statusCode: code,
	}
}

/*GetPollersDefault handles this case with default header values.

Unexpected error
*/
type GetPollersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get pollers default response
func (o *GetPollersDefault) Code() int {
	return o._statusCode
}

func (o *GetPollersDefault) Error() string {
	return fmt.Sprintf("[GET /pollers][%d] GetPollers default  %+v", o._statusCode, o.Payload)
}

func (o *GetPollersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
