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

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/*GetUserOK handles this case with default header values.

Successfully retrieved the specified user
*/
type GetUserOK struct {
	Payload *models.UserObj
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{name}][%d] getUserOK  %+v", 200, o.Payload)
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserObj)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserUnauthorized creates a GetUserUnauthorized with default headers values
func NewGetUserUnauthorized() *GetUserUnauthorized {
	return &GetUserUnauthorized{}
}

/*GetUserUnauthorized handles this case with default header values.

Unauthorized
*/
type GetUserUnauthorized struct {
	Payload GetUserUnauthorizedBodyBody
}

func (o *GetUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{name}][%d] getUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserForbidden creates a GetUserForbidden with default headers values
func NewGetUserForbidden() *GetUserForbidden {
	return &GetUserForbidden{}
}

/*GetUserForbidden handles this case with default header values.

Forbidden
*/
type GetUserForbidden struct {
	Payload GetUserForbiddenBodyBody
}

func (o *GetUserForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{name}][%d] getUserForbidden  %+v", 403, o.Payload)
}

func (o *GetUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserDefault creates a GetUserDefault with default headers values
func NewGetUserDefault(code int) *GetUserDefault {
	return &GetUserDefault{
		_statusCode: code,
	}
}

/*GetUserDefault handles this case with default header values.

Unexpected error
*/
type GetUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get user default response
func (o *GetUserDefault) Code() int {
	return o._statusCode
}

func (o *GetUserDefault) Error() string {
	return fmt.Sprintf("[GET /users/{name}][%d] getUser default  %+v", o._statusCode, o.Payload)
}

func (o *GetUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetUserForbiddenBodyBody get user forbidden body body

swagger:model GetUserForbiddenBodyBody
*/
type GetUserForbiddenBodyBody interface{}

/*GetUserUnauthorizedBodyBody get user unauthorized body body

swagger:model GetUserUnauthorizedBodyBody
*/
type GetUserUnauthorizedBodyBody interface{}
