// Code generated by go-swagger; DO NOT EDIT.

package quiz

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new quiz API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quiz API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DisableQuiz(params *DisableQuizParams, opts ...ClientOption) (*DisableQuizOK, error)

	ChangeEmail(params *ChangeEmailParams, opts ...ClientOption) (*ChangeEmailOK, error)

	EnableQuiz(params *EnableQuizParams, opts ...ClientOption) (*EnableQuizOK, error)

	GetQuiz(params *GetQuizParams, opts ...ClientOption) (*GetQuizOK, error)

	RegisterQuiz(params *RegisterQuizParams, opts ...ClientOption) (*RegisterQuizOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DisableQuiz disables a quiz
*/
func (a *Client) DisableQuiz(params *DisableQuizParams, opts ...ClientOption) (*DisableQuizOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableQuizParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DisableQuiz",
		Method:             "PUT",
		PathPattern:        "/api/quizs/{id}/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisableQuizReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableQuizOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DisableQuizDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChangeEmail changes a quizs email
*/
func (a *Client) ChangeEmail(params *ChangeEmailParams, opts ...ClientOption) (*ChangeEmailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeEmailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "changeEmail",
		Method:             "PUT",
		PathPattern:        "/api/quizs/{id}/change-email",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChangeEmailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeEmailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeEmailDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EnableQuiz enables a quiz
*/
func (a *Client) EnableQuiz(params *EnableQuizParams, opts ...ClientOption) (*EnableQuizOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableQuizParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enableQuiz",
		Method:             "PUT",
		PathPattern:        "/api/quizs/{id}/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EnableQuizReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableQuizOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EnableQuizDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetQuiz gets a quiz
*/
func (a *Client) GetQuiz(params *GetQuizParams, opts ...ClientOption) (*GetQuizOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQuizParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getQuiz",
		Method:             "GET",
		PathPattern:        "/api/quizs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetQuizReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetQuizOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetQuizDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RegisterQuiz creates a new quiz
*/
func (a *Client) RegisterQuiz(params *RegisterQuizParams, opts ...ClientOption) (*RegisterQuizOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterQuizParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "registerQuiz",
		Method:             "POST",
		PathPattern:        "/api/quizs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterQuizReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegisterQuizOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RegisterQuizDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
