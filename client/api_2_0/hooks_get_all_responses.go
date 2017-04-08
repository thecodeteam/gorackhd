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

// HooksGetAllReader is a Reader for the HooksGetAll structure.
type HooksGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *HooksGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHooksGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewHooksGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewHooksGetAllOK creates a HooksGetAllOK with default headers values
func NewHooksGetAllOK() *HooksGetAllOK {
	return &HooksGetAllOK{}
}

/*HooksGetAllOK handles this case with default header values.

Successfully retrieved the hook configuration
*/
type HooksGetAllOK struct {
	Payload []*models.Hooks20HookBase
}

func (o *HooksGetAllOK) Error() string {
	return fmt.Sprintf("[GET /hooks][%d] hooksGetAllOK  %+v", 200, o.Payload)
}

func (o *HooksGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHooksGetAllDefault creates a HooksGetAllDefault with default headers values
func NewHooksGetAllDefault(code int) *HooksGetAllDefault {
	return &HooksGetAllDefault{
		_statusCode: code,
	}
}

/*HooksGetAllDefault handles this case with default header values.

Unexpected error
*/
type HooksGetAllDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the hooks get all default response
func (o *HooksGetAllDefault) Code() int {
	return o._statusCode
}

func (o *HooksGetAllDefault) Error() string {
	return fmt.Sprintf("[GET /hooks][%d] hooksGetAll default  %+v", o._statusCode, o.Payload)
}

func (o *HooksGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
