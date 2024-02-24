package grpc

import (
	"context"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"quizazz/internal/errorsotel"
	"quizazz/quiz-service/internal/application"
	"quizazz/quiz-service/internal/domain"
	"quizazz/quiz-service/quizspb"
)

type server struct {
	app application.App
	quizspb.UnimplementedQuizsServiceServer
}

var _ quizspb.QuizsServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	quizspb.RegisterQuizsServiceServer(registrar, server{
		app: app,
	})
	return nil
}

func (s server) RegisterQuiz(ctx context.Context, request *quizspb.CreateQuizRequest) (resp *quizspb.CreateQuizResponse, err error) {
	span := trace.SpanFromContext(ctx)

	id := uuid.New().String()

	span.SetAttributes(
		attribute.String("QuizID", id),
	)

	err = s.app.CreateQuiz(ctx, application.CreateQuiz{
		ID:       id,
		QuizName: request.GetName(),
		QuizType: request.GetQuizType(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &quizspb.CreateQuizResponse{Id: id}, err
}

func (s server) AuthorizeQuiz(ctx context.Context, request *quizspb.AuthorizeQuizRequest) (resp *quizspb.AuthorizeQuizResponse, err error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("QuizID", request.GetId()),
	)

	err = s.app.AuthorizeQuiz(ctx, application.AuthorizeQuiz{
		ID: request.GetId(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &quizspb.AuthorizeQuizResponse{}, err
}

func (s server) GetQuiz(ctx context.Context, request *quizspb.GetQuizRequest) (resp *quizspb.GetQuizResponse, err error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("QuizID", request.GetId()),
	)

	quiz, err := s.app.GetQuiz(ctx, application.GetQuiz{
		ID: request.GetId(),
	})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	return &quizspb.GetQuizResponse{
		Quiz: s.quizFromDomain(quiz),
	}, nil
}

func (s server) EnableQuiz(ctx context.Context, request *quizspb.EnableQuizRequest) (resp *quizspb.EnableQuizResponse, err error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("QuizID", request.GetId()),
	)

	err = s.app.EnableQuiz(ctx, application.EnableQuiz{ID: request.GetId()})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &quizspb.EnableQuizResponse{}, err
}

func (s server) DisableQuiz(ctx context.Context, request *quizspb.DisableQuizRequest) (resp *quizspb.DisableQuizResponse, err error) {
	span := trace.SpanFromContext(ctx)

	span.SetAttributes(
		attribute.String("QuizID", request.GetId()),
	)

	err = s.app.DisableQuiz(ctx, application.DisableQuiz{ID: request.GetId()})
	if err != nil {
		span.RecordError(err, trace.WithAttributes(errorsotel.ErrAttrs(err)...))
		span.SetStatus(codes.Error, err.Error())
	}

	return &quizspb.DisableQuizResponse{}, err
}

func (s server) quizFromDomain(quiz *domain.Quiz) *quizspb.Quiz {
	return &quizspb.Quiz{
		Id:       quiz.ID(),
		Name:     quiz.QuizName,
		QuizType: quiz.QuizType,
	}
}
