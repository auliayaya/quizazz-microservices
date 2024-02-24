package grpc

import (
	"context"
	"database/sql"
	"google.golang.org/grpc"
	"quizazz/internal/di"
	"quizazz/quiz-service/internal/application"
	"quizazz/quiz-service/internal/constants"
	"quizazz/quiz-service/quizspb"
)

type serverTx struct {
	c di.Container
	quizspb.UnimplementedQuizsServiceServer
}

var _ quizspb.QuizsServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	quizspb.RegisterQuizsServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) CreateQuiz(ctx context.Context, request *quizspb.CreateQuizRequest) (resp *quizspb.CreateQuizResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}

	return next.CreateQuiz(ctx, request)
}

func (s serverTx) AuthorizeQuiz(ctx context.Context, request *quizspb.AuthorizeQuizRequest) (resp *quizspb.AuthorizeQuizResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}

	return next.AuthorizeQuiz(ctx, request)
}

func (s serverTx) GetQuiz(ctx context.Context, request *quizspb.GetQuizRequest) (resp *quizspb.GetQuizResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}

	return next.GetQuiz(ctx, request)
}

func (s serverTx) EnableQuiz(ctx context.Context, request *quizspb.EnableQuizRequest) (resp *quizspb.EnableQuizResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}

	return next.EnableQuiz(ctx, request)
}

func (s serverTx) DisableQuiz(ctx context.Context, request *quizspb.DisableQuizRequest) (resp *quizspb.DisableQuizResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}

	return next.DisableQuiz(ctx, request)
}

func (s serverTx) closeTx(tx *sql.Tx, err error) error {
	if p := recover(); p != nil {
		_ = tx.Rollback()
		panic(p)
	} else if err != nil {
		_ = tx.Rollback()
		return err
	} else {
		return tx.Commit()
	}
}
