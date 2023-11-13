// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mihailtudos/microservices/sdk/models"
)

// ListSingleProductReader is a Reader for the ListSingleProduct structure.
type ListSingleProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSingleProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSingleProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListSingleProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /products/{id}] listSingleProduct", response, response.Code())
	}
}

// NewListSingleProductOK creates a ListSingleProductOK with default headers values
func NewListSingleProductOK() *ListSingleProductOK {
	return &ListSingleProductOK{}
}

/*
ListSingleProductOK describes a response with status code 200, with default header values.

Data structure representing a single product
*/
type ListSingleProductOK struct {
	Payload *models.Product
}

// IsSuccess returns true when this list single product o k response has a 2xx status code
func (o *ListSingleProductOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list single product o k response has a 3xx status code
func (o *ListSingleProductOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list single product o k response has a 4xx status code
func (o *ListSingleProductOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list single product o k response has a 5xx status code
func (o *ListSingleProductOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list single product o k response a status code equal to that given
func (o *ListSingleProductOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list single product o k response
func (o *ListSingleProductOK) Code() int {
	return 200
}

func (o *ListSingleProductOK) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] listSingleProductOK  %+v", 200, o.Payload)
}

func (o *ListSingleProductOK) String() string {
	return fmt.Sprintf("[GET /products/{id}][%d] listSingleProductOK  %+v", 200, o.Payload)
}

func (o *ListSingleProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *ListSingleProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSingleProductNotFound creates a ListSingleProductNotFound with default headers values
func NewListSingleProductNotFound() *ListSingleProductNotFound {
	return &ListSingleProductNotFound{}
}

/*
ListSingleProductNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type ListSingleProductNotFound struct {
	Payload *models.GenericError
}

// IsSuccess returns true when this list single product not found response has a 2xx status code
func (o *ListSingleProductNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list single product not found response has a 3xx status code
func (o *ListSingleProductNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list single product not found response has a 4xx status code
func (o *ListSingleProductNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list single product not found response has a 5xx status code
func (o *ListSingleProductNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list single product not found response a status code equal to that given
func (o *ListSingleProductNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list single product not found response
func (o *ListSingleProductNotFound) Code() int {
	return 404
}

func (o *ListSingleProductNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{id}][%d] listSingleProductNotFound  %+v", 404, o.Payload)
}

func (o *ListSingleProductNotFound) String() string {
	return fmt.Sprintf("[GET /products/{id}][%d] listSingleProductNotFound  %+v", 404, o.Payload)
}

func (o *ListSingleProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *ListSingleProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
