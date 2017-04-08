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

// NodesGetCatalogSourceByIDReader is a Reader for the NodesGetCatalogSourceByID structure.
type NodesGetCatalogSourceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *NodesGetCatalogSourceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNodesGetCatalogSourceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewNodesGetCatalogSourceByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNodesGetCatalogSourceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewNodesGetCatalogSourceByIDOK creates a NodesGetCatalogSourceByIDOK with default headers values
func NewNodesGetCatalogSourceByIDOK() *NodesGetCatalogSourceByIDOK {
	return &NodesGetCatalogSourceByIDOK{}
}

/*NodesGetCatalogSourceByIDOK handles this case with default header values.

Successfully retrieved specific source catalog of specified node
*/
type NodesGetCatalogSourceByIDOK struct {
	Payload NodesGetCatalogSourceByIDOKBodyBody
}

func (o *NodesGetCatalogSourceByIDOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/catalogs/{source}][%d] nodesGetCatalogSourceByIdOK  %+v", 200, o.Payload)
}

func (o *NodesGetCatalogSourceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesGetCatalogSourceByIDNotFound creates a NodesGetCatalogSourceByIDNotFound with default headers values
func NewNodesGetCatalogSourceByIDNotFound() *NodesGetCatalogSourceByIDNotFound {
	return &NodesGetCatalogSourceByIDNotFound{}
}

/*NodesGetCatalogSourceByIDNotFound handles this case with default header values.

The specified node was not found
*/
type NodesGetCatalogSourceByIDNotFound struct {
	Payload *models.Error
}

func (o *NodesGetCatalogSourceByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/catalogs/{source}][%d] nodesGetCatalogSourceByIdNotFound  %+v", 404, o.Payload)
}

func (o *NodesGetCatalogSourceByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesGetCatalogSourceByIDDefault creates a NodesGetCatalogSourceByIDDefault with default headers values
func NewNodesGetCatalogSourceByIDDefault(code int) *NodesGetCatalogSourceByIDDefault {
	return &NodesGetCatalogSourceByIDDefault{
		_statusCode: code,
	}
}

/*NodesGetCatalogSourceByIDDefault handles this case with default header values.

Unexpected error
*/
type NodesGetCatalogSourceByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the nodes get catalog source by Id default response
func (o *NodesGetCatalogSourceByIDDefault) Code() int {
	return o._statusCode
}

func (o *NodesGetCatalogSourceByIDDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/catalogs/{source}][%d] nodesGetCatalogSourceById default  %+v", o._statusCode, o.Payload)
}

func (o *NodesGetCatalogSourceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*NodesGetCatalogSourceByIDOKBodyBody nodes get catalog source by ID o k body body

swagger:model NodesGetCatalogSourceByIDOKBodyBody
*/
type NodesGetCatalogSourceByIDOKBodyBody interface{}
