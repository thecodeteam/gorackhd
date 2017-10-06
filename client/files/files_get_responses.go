// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// FilesGetReader is a Reader for the FilesGet structure.
type FilesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFilesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewFilesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewFilesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewFilesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFilesGetOK creates a FilesGetOK with default headers values
func NewFilesGetOK() *FilesGetOK {
	return &FilesGetOK{}
}

/*FilesGetOK handles this case with default header values.

Successfully retrieved specified file
*/
type FilesGetOK struct {
	Payload models.FilesGetOKBody
}

func (o *FilesGetOK) Error() string {
	return fmt.Sprintf("[GET /files/{fileidentifier}][%d] filesGetOK  %+v", 200, o.Payload)
}

func (o *FilesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilesGetNotFound creates a FilesGetNotFound with default headers values
func NewFilesGetNotFound() *FilesGetNotFound {
	return &FilesGetNotFound{}
}

/*FilesGetNotFound handles this case with default header values.

File not found
*/
type FilesGetNotFound struct {
	Payload *models.Error
}

func (o *FilesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /files/{fileidentifier}][%d] filesGetNotFound  %+v", 404, o.Payload)
}

func (o *FilesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilesGetInternalServerError creates a FilesGetInternalServerError with default headers values
func NewFilesGetInternalServerError() *FilesGetInternalServerError {
	return &FilesGetInternalServerError{}
}

/*FilesGetInternalServerError handles this case with default header values.

Failed to serve file request
*/
type FilesGetInternalServerError struct {
	Payload *models.Error
}

func (o *FilesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /files/{fileidentifier}][%d] filesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *FilesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilesGetDefault creates a FilesGetDefault with default headers values
func NewFilesGetDefault(code int) *FilesGetDefault {
	return &FilesGetDefault{
		_statusCode: code,
	}
}

/*FilesGetDefault handles this case with default header values.

Unexpected error
*/
type FilesGetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the files get default response
func (o *FilesGetDefault) Code() int {
	return o._statusCode
}

func (o *FilesGetDefault) Error() string {
	return fmt.Sprintf("[GET /files/{fileidentifier}][%d] filesGet default  %+v", o._statusCode, o.Payload)
}

func (o *FilesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}