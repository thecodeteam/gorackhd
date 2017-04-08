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

// ViewsGetByIDReader is a Reader for the ViewsGetByID structure.
type ViewsGetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ViewsGetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewViewsGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewViewsGetByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewViewsGetByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewViewsGetByIDOK creates a ViewsGetByIDOK with default headers values
func NewViewsGetByIDOK() *ViewsGetByIDOK {
	return &ViewsGetByIDOK{}
}

/*ViewsGetByIDOK handles this case with default header values.

Successfully retrieved the specified view
*/
type ViewsGetByIDOK struct {
	Payload ViewsGetByIDOKBodyBody
}

func (o *ViewsGetByIDOK) Error() string {
	return fmt.Sprintf("[GET /views/{identifier}][%d] viewsGetByIdOK  %+v", 200, o.Payload)
}

func (o *ViewsGetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewsGetByIDNotFound creates a ViewsGetByIDNotFound with default headers values
func NewViewsGetByIDNotFound() *ViewsGetByIDNotFound {
	return &ViewsGetByIDNotFound{}
}

/*ViewsGetByIDNotFound handles this case with default header values.

The view with specified name was not found
*/
type ViewsGetByIDNotFound struct {
	Payload *models.Error
}

func (o *ViewsGetByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /views/{identifier}][%d] viewsGetByIdNotFound  %+v", 404, o.Payload)
}

func (o *ViewsGetByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewsGetByIDDefault creates a ViewsGetByIDDefault with default headers values
func NewViewsGetByIDDefault(code int) *ViewsGetByIDDefault {
	return &ViewsGetByIDDefault{
		_statusCode: code,
	}
}

/*ViewsGetByIDDefault handles this case with default header values.

Unexpected error
*/
type ViewsGetByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the views get by Id default response
func (o *ViewsGetByIDDefault) Code() int {
	return o._statusCode
}

func (o *ViewsGetByIDDefault) Error() string {
	return fmt.Sprintf("[GET /views/{identifier}][%d] viewsGetById default  %+v", o._statusCode, o.Payload)
}

func (o *ViewsGetByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ViewsGetByIDOKBodyBody views get by ID o k body body

swagger:model ViewsGetByIDOKBodyBody
*/
type ViewsGetByIDOKBodyBody interface{}
