// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// GetAllTagsReader is a Reader for the GetAllTags structure.
type GetAllTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetAllTagsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetAllTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAllTagsOK creates a GetAllTagsOK with default headers values
func NewGetAllTagsOK() *GetAllTagsOK {
	return &GetAllTagsOK{}
}

/*GetAllTagsOK handles this case with default header values.

Successfully retrieved all tags.
*/
type GetAllTagsOK struct {
	Payload []interface{}
}

func (o *GetAllTagsOK) Error() string {
	return fmt.Sprintf("[GET /tags][%d] getAllTagsOK  %+v", 200, o.Payload)
}

func (o *GetAllTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTagsNotFound creates a GetAllTagsNotFound with default headers values
func NewGetAllTagsNotFound() *GetAllTagsNotFound {
	return &GetAllTagsNotFound{}
}

/*GetAllTagsNotFound handles this case with default header values.

No tags found.
*/
type GetAllTagsNotFound struct {
	Payload *models.Error
}

func (o *GetAllTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /tags][%d] getAllTagsNotFound  %+v", 404, o.Payload)
}

func (o *GetAllTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTagsDefault creates a GetAllTagsDefault with default headers values
func NewGetAllTagsDefault(code int) *GetAllTagsDefault {
	return &GetAllTagsDefault{
		_statusCode: code,
	}
}

/*GetAllTagsDefault handles this case with default header values.

Unexpected error
*/
type GetAllTagsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get all tags default response
func (o *GetAllTagsDefault) Code() int {
	return o._statusCode
}

func (o *GetAllTagsDefault) Error() string {
	return fmt.Sprintf("[GET /tags][%d] getAllTags default  %+v", o._statusCode, o.Payload)
}

func (o *GetAllTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}