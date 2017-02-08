package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/rackn/rocket-skates/models"
)

// DeleteFileReader is a Reader for the DeleteFile structure.
type DeleteFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteFileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFileNoContent creates a DeleteFileNoContent with default headers values
func NewDeleteFileNoContent() *DeleteFileNoContent {
	return &DeleteFileNoContent{}
}

/*DeleteFileNoContent handles this case with default header values.

DeleteFileNoContent delete file no content
*/
type DeleteFileNoContent struct {
}

func (o *DeleteFileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /files/{path}][%d] deleteFileNoContent ", 204)
}

func (o *DeleteFileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFileForbidden creates a DeleteFileForbidden with default headers values
func NewDeleteFileForbidden() *DeleteFileForbidden {
	return &DeleteFileForbidden{}
}

/*DeleteFileForbidden handles this case with default header values.

DeleteFileForbidden delete file forbidden
*/
type DeleteFileForbidden struct {
	Payload *models.Result
}

func (o *DeleteFileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /files/{path}][%d] deleteFileForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFileNotFound creates a DeleteFileNotFound with default headers values
func NewDeleteFileNotFound() *DeleteFileNotFound {
	return &DeleteFileNotFound{}
}

/*DeleteFileNotFound handles this case with default header values.

DeleteFileNotFound delete file not found
*/
type DeleteFileNotFound struct {
	Payload *models.Result
}

func (o *DeleteFileNotFound) Error() string {
	return fmt.Sprintf("[DELETE /files/{path}][%d] deleteFileNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFileConflict creates a DeleteFileConflict with default headers values
func NewDeleteFileConflict() *DeleteFileConflict {
	return &DeleteFileConflict{}
}

/*DeleteFileConflict handles this case with default header values.

DeleteFileConflict delete file conflict
*/
type DeleteFileConflict struct {
	Payload *models.Result
}

func (o *DeleteFileConflict) Error() string {
	return fmt.Sprintf("[DELETE /files/{path}][%d] deleteFileConflict  %+v", 409, o.Payload)
}

func (o *DeleteFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFileInternalServerError creates a DeleteFileInternalServerError with default headers values
func NewDeleteFileInternalServerError() *DeleteFileInternalServerError {
	return &DeleteFileInternalServerError{}
}

/*DeleteFileInternalServerError handles this case with default header values.

DeleteFileInternalServerError delete file internal server error
*/
type DeleteFileInternalServerError struct {
	Payload *models.Result
}

func (o *DeleteFileInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /files/{path}][%d] deleteFileInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
