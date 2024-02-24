package grpc

import (
	"context"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"quizazz/internal/errorsotel"
	"quizazz/user-service/internal/application"
	"quizazz/user-service/internal/domain"
	"quizazz/user-service/userspb"
)

type server struct {
	app application.App
	userspb.UnimplementedUsersServiceServer
}

var _ userspb.UsersServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	userspb.RegisterUsersServiceServer(registrar, server{
		app: app,
	})
	return nil
}

func (s server) RegisterUser(ctx context.Context, request *userspb.RegisterUserRequest) (resp *userspb.RegisterUserResponse, err error) {
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()

	span.SetAttributes(
		attribute.String("UserID", id),
	)

	err = s.app.RegisterUser(ctx, application.RegisterUser{
		ID:    id,
		Name:  request.GetName(),
		Email: request.GetEmail(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &userspb.RegisterUserResponse{Id: id}, err
}

func (s server) AuthorizeUser(ctx context.Context, request *userspb.AuthorizeUserRequest) (resp *userspb.AuthorizeUserResponse, err error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("UserID", request.GetId()),
	)

	err = s.app.AuthorizeUser(ctx, application.AuthorizeUser{
		ID: request.GetId(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &userspb.AuthorizeUserResponse{}, err
}

func (s server) GetUser(ctx context.Context, request *userspb.GetUserRequest) (resp *userspb.GetUserResponse, err error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("UserID", request.GetId()),
	)

	user, err := s.app.GetUser(ctx, application.GetUser{
		ID: request.GetId(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &userspb.GetUserResponse{
		User: s.userFromDomain(user),
	}, nil
}

func (s server) EnableUser(ctx context.Context, request *userspb.EnableUserRequest) (resp *userspb.EnableUserResponse, err error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("UserID", request.GetId()),
	)

	err = s.app.EnableUser(ctx, application.EnableUser{ID: request.GetId()})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &userspb.EnableUserResponse{}, err
}

func (s server) DisableUser(ctx context.Context, request *userspb.DisableUserRequest) (resp *userspb.DisableUserResponse, err error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("UserID", request.GetId()),
	)

	err = s.app.DisableUser(ctx, application.DisableUser{ID: request.GetId()})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &userspb.DisableUserResponse{}, err
}

func (s server) userFromDomain(user *domain.User) *userspb.User {
	return &userspb.User{
		Id:      user.ID(),
		Name:    user.Name,
		Email:   user.Email,
		Enabled: user.Enabled,
	}
}
