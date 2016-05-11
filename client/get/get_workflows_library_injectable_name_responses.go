package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetWorkflowsLibraryInjectableNameReader is a Reader for the GetWorkflowsLibraryInjectableName structure.
type GetWorkflowsLibraryInjectableNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetWorkflowsLibraryInjectableNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWorkflowsLibraryInjectableNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetWorkflowsLibraryInjectableNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetWorkflowsLibraryInjectableNameOK creates a GetWorkflowsLibraryInjectableNameOK with default headers values
func NewGetWorkflowsLibraryInjectableNameOK() *GetWorkflowsLibraryInjectableNameOK {
	return &GetWorkflowsLibraryInjectableNameOK{}
}

/*GetWorkflowsLibraryInjectableNameOK handles this case with default header values.

List all workflows available to run

*/
type GetWorkflowsLibraryInjectableNameOK struct {
	Payload GetWorkflowsLibraryInjectableNameOKBodyBody
}

func (o *GetWorkflowsLibraryInjectableNameOK) Error() string {
	return fmt.Sprintf("[GET /workflows/library/{injectableName}][%d] getWorkflowsLibraryInjectableNameOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowsLibraryInjectableNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowsLibraryInjectableNameDefault creates a GetWorkflowsLibraryInjectableNameDefault with default headers values
func NewGetWorkflowsLibraryInjectableNameDefault(code int) *GetWorkflowsLibraryInjectableNameDefault {
	return &GetWorkflowsLibraryInjectableNameDefault{
		_statusCode: code,
	}
}

/*GetWorkflowsLibraryInjectableNameDefault handles this case with default header values.

Unexpected error
*/
type GetWorkflowsLibraryInjectableNameDefault struct {
	_statusCode int

	Payload GetWorkflowsLibraryInjectableNameDefaultBodyBody
}

// Code gets the status code for the get workflows library injectable name default response
func (o *GetWorkflowsLibraryInjectableNameDefault) Code() int {
	return o._statusCode
}

func (o *GetWorkflowsLibraryInjectableNameDefault) Error() string {
	return fmt.Sprintf("[GET /workflows/library/{injectableName}][%d] GetWorkflowsLibraryInjectableName default  %+v", o._statusCode, o.Payload)
}

func (o *GetWorkflowsLibraryInjectableNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWorkflowsLibraryInjectableNameOKBodyBody get workflows library injectable name o k body body

swagger:model GetWorkflowsLibraryInjectableNameOKBodyBody
*/
type GetWorkflowsLibraryInjectableNameOKBodyBody interface{}

/*GetWorkflowsLibraryInjectableNameDefaultBodyBody get workflows library injectable name default body body

swagger:model GetWorkflowsLibraryInjectableNameDefaultBodyBody
*/
type GetWorkflowsLibraryInjectableNameDefaultBodyBody interface{}
