// Code generated by go-swagger; DO NOT EDIT.

package quiz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"quizazz/quiz-service/quizsclient/models"
)

// RegisterQuizReader is a Reader for the RegisterQuiz structure.
type RegisterQuizReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterQuizReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterQuizOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRegisterQuizDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRegisterQuizOK creates a RegisterQuizOK with default headers values
func NewRegisterQuizOK() *RegisterQuizOK {
	return &RegisterQuizOK{}
}

/*
RegisterQuizOK describes a response with status code 200, with default header values.

A successful response.
*/
type RegisterQuizOK struct {
	Payload *models.QuizspbCreateQuizResponse
}

// IsSuccess returns true when this register quiz o k response has a 2xx status code
func (o *RegisterQuizOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this register quiz o k response has a 3xx status code
func (o *RegisterQuizOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register quiz o k response has a 4xx status code
func (o *RegisterQuizOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this register quiz o k response has a 5xx status code
func (o *RegisterQuizOK) IsServerError() bool {
	return false
}

// IsCode returns true when this register quiz o k response a status code equal to that given
func (o *RegisterQuizOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the register quiz o k response
func (o *RegisterQuizOK) Code() int {
	return 200
}

func (o *RegisterQuizOK) Error() string {
	return fmt.Sprintf("[POST /api/quizs][%d] registerQuizOK  %+v", 200, o.Payload)
}

func (o *RegisterQuizOK) String() string {
	return fmt.Sprintf("[POST /api/quizs][%d] registerQuizOK  %+v", 200, o.Payload)
}

func (o *RegisterQuizOK) GetPayload() *models.QuizspbCreateQuizResponse {
	return o.Payload
}

func (o *RegisterQuizOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuizspbCreateQuizResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterQuizDefault creates a RegisterQuizDefault with default headers values
func NewRegisterQuizDefault(code int) *RegisterQuizDefault {
	return &RegisterQuizDefault{
		_statusCode: code,
	}
}

/*
RegisterQuizDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RegisterQuizDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this register quiz default response has a 2xx status code
func (o *RegisterQuizDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this register quiz default response has a 3xx status code
func (o *RegisterQuizDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this register quiz default response has a 4xx status code
func (o *RegisterQuizDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this register quiz default response has a 5xx status code
func (o *RegisterQuizDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this register quiz default response a status code equal to that given
func (o *RegisterQuizDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the register quiz default response
func (o *RegisterQuizDefault) Code() int {
	return o._statusCode
}

func (o *RegisterQuizDefault) Error() string {
	return fmt.Sprintf("[POST /api/quizs][%d] registerQuiz default  %+v", o._statusCode, o.Payload)
}

func (o *RegisterQuizDefault) String() string {
	return fmt.Sprintf("[POST /api/quizs][%d] registerQuiz default  %+v", o._statusCode, o.Payload)
}

func (o *RegisterQuizDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *RegisterQuizDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
