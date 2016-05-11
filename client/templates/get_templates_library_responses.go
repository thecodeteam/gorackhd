package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// GetTemplatesLibraryReader is a Reader for the GetTemplatesLibrary structure.
type GetTemplatesLibraryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetTemplatesLibraryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTemplatesLibraryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetTemplatesLibraryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetTemplatesLibraryOK creates a GetTemplatesLibraryOK with default headers values
func NewGetTemplatesLibraryOK() *GetTemplatesLibraryOK {
	return &GetTemplatesLibraryOK{}
}

/*GetTemplatesLibraryOK handles this case with default header values.

list of possible templates

*/
type GetTemplatesLibraryOK struct {
	Payload GetTemplatesLibraryOKBodyBody
}

func (o *GetTemplatesLibraryOK) Error() string {
	return fmt.Sprintf("[GET /templates/library][%d] getTemplatesLibraryOK  %+v", 200, o.Payload)
}

func (o *GetTemplatesLibraryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTemplatesLibraryDefault creates a GetTemplatesLibraryDefault with default headers values
func NewGetTemplatesLibraryDefault(code int) *GetTemplatesLibraryDefault {
	return &GetTemplatesLibraryDefault{
		_statusCode: code,
	}
}

/*GetTemplatesLibraryDefault handles this case with default header values.

Unexpected error
*/
type GetTemplatesLibraryDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get templates library default response
func (o *GetTemplatesLibraryDefault) Code() int {
	return o._statusCode
}

func (o *GetTemplatesLibraryDefault) Error() string {
	return fmt.Sprintf("[GET /templates/library][%d] GetTemplatesLibrary default  %+v", o._statusCode, o.Payload)
}

func (o *GetTemplatesLibraryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetTemplatesLibraryOKBodyBody get templates library o k body body

swagger:model GetTemplatesLibraryOKBodyBody
*/
type GetTemplatesLibraryOKBodyBody interface{}
