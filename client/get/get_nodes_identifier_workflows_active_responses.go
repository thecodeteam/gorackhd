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

// GetNodesIdentifierWorkflowsActiveReader is a Reader for the GetNodesIdentifierWorkflowsActive structure.
type GetNodesIdentifierWorkflowsActiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNodesIdentifierWorkflowsActiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesIdentifierWorkflowsActiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNodesIdentifierWorkflowsActiveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNodesIdentifierWorkflowsActiveDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetNodesIdentifierWorkflowsActiveOK creates a GetNodesIdentifierWorkflowsActiveOK with default headers values
func NewGetNodesIdentifierWorkflowsActiveOK() *GetNodesIdentifierWorkflowsActiveOK {
	return &GetNodesIdentifierWorkflowsActiveOK{}
}

/*GetNodesIdentifierWorkflowsActiveOK handles this case with default header values.

Fetch currently running workflows for specified node

*/
type GetNodesIdentifierWorkflowsActiveOK struct {
	Payload GetNodesIdentifierWorkflowsActiveOKBodyBody
}

func (o *GetNodesIdentifierWorkflowsActiveOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/workflows/active][%d] getNodesIdentifierWorkflowsActiveOK  %+v", 200, o.Payload)
}

func (o *GetNodesIdentifierWorkflowsActiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierWorkflowsActiveNotFound creates a GetNodesIdentifierWorkflowsActiveNotFound with default headers values
func NewGetNodesIdentifierWorkflowsActiveNotFound() *GetNodesIdentifierWorkflowsActiveNotFound {
	return &GetNodesIdentifierWorkflowsActiveNotFound{}
}

/*GetNodesIdentifierWorkflowsActiveNotFound handles this case with default header values.

The node with the identifier was not found

*/
type GetNodesIdentifierWorkflowsActiveNotFound struct {
	Payload *models.Error
}

func (o *GetNodesIdentifierWorkflowsActiveNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/workflows/active][%d] getNodesIdentifierWorkflowsActiveNotFound  %+v", 404, o.Payload)
}

func (o *GetNodesIdentifierWorkflowsActiveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierWorkflowsActiveDefault creates a GetNodesIdentifierWorkflowsActiveDefault with default headers values
func NewGetNodesIdentifierWorkflowsActiveDefault(code int) *GetNodesIdentifierWorkflowsActiveDefault {
	return &GetNodesIdentifierWorkflowsActiveDefault{
		_statusCode: code,
	}
}

/*GetNodesIdentifierWorkflowsActiveDefault handles this case with default header values.

Unexpected error
*/
type GetNodesIdentifierWorkflowsActiveDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nodes identifier workflows active default response
func (o *GetNodesIdentifierWorkflowsActiveDefault) Code() int {
	return o._statusCode
}

func (o *GetNodesIdentifierWorkflowsActiveDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/workflows/active][%d] GetNodesIdentifierWorkflowsActive default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodesIdentifierWorkflowsActiveDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetNodesIdentifierWorkflowsActiveOKBodyBody get nodes identifier workflows active o k body body

swagger:model GetNodesIdentifierWorkflowsActiveOKBodyBody
*/
type GetNodesIdentifierWorkflowsActiveOKBodyBody interface{}
