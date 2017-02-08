package bootenvs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/models"
)

// PatchBootenvReader is a Reader for the PatchBootenv structure.
type PatchBootenvReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchBootenvReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewPatchBootenvAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchBootenvUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchBootenvNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 417:
		result := NewPatchBootenvExpectationFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchBootenvInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchBootenvAccepted creates a PatchBootenvAccepted with default headers values
func NewPatchBootenvAccepted() *PatchBootenvAccepted {
	return &PatchBootenvAccepted{}
}

/*PatchBootenvAccepted handles this case with default header values.

PatchBootenvAccepted patch bootenv accepted
*/
type PatchBootenvAccepted struct {
	Payload *models.BootenvInput
}

func (o *PatchBootenvAccepted) Error() string {
	return fmt.Sprintf("[PATCH /bootenvs/{name}][%d] patchBootenvAccepted  %+v", 202, o.Payload)
}

func (o *PatchBootenvAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BootenvInput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBootenvUnauthorized creates a PatchBootenvUnauthorized with default headers values
func NewPatchBootenvUnauthorized() *PatchBootenvUnauthorized {
	return &PatchBootenvUnauthorized{}
}

/*PatchBootenvUnauthorized handles this case with default header values.

PatchBootenvUnauthorized patch bootenv unauthorized
*/
type PatchBootenvUnauthorized struct {
	Payload *models.Result
}

func (o *PatchBootenvUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /bootenvs/{name}][%d] patchBootenvUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchBootenvUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBootenvNotFound creates a PatchBootenvNotFound with default headers values
func NewPatchBootenvNotFound() *PatchBootenvNotFound {
	return &PatchBootenvNotFound{}
}

/*PatchBootenvNotFound handles this case with default header values.

PatchBootenvNotFound patch bootenv not found
*/
type PatchBootenvNotFound struct {
	Payload *models.Result
}

func (o *PatchBootenvNotFound) Error() string {
	return fmt.Sprintf("[PATCH /bootenvs/{name}][%d] patchBootenvNotFound  %+v", 404, o.Payload)
}

func (o *PatchBootenvNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBootenvExpectationFailed creates a PatchBootenvExpectationFailed with default headers values
func NewPatchBootenvExpectationFailed() *PatchBootenvExpectationFailed {
	return &PatchBootenvExpectationFailed{}
}

/*PatchBootenvExpectationFailed handles this case with default header values.

PatchBootenvExpectationFailed patch bootenv expectation failed
*/
type PatchBootenvExpectationFailed struct {
	Payload *models.Result
}

func (o *PatchBootenvExpectationFailed) Error() string {
	return fmt.Sprintf("[PATCH /bootenvs/{name}][%d] patchBootenvExpectationFailed  %+v", 417, o.Payload)
}

func (o *PatchBootenvExpectationFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBootenvInternalServerError creates a PatchBootenvInternalServerError with default headers values
func NewPatchBootenvInternalServerError() *PatchBootenvInternalServerError {
	return &PatchBootenvInternalServerError{}
}

/*PatchBootenvInternalServerError handles this case with default header values.

PatchBootenvInternalServerError patch bootenv internal server error
*/
type PatchBootenvInternalServerError struct {
	Payload *models.Result
}

func (o *PatchBootenvInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /bootenvs/{name}][%d] patchBootenvInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchBootenvInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
