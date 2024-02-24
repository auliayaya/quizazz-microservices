package grpc

import (
	"context"
	"database/sql"
	"google.golang.org/grpc"
	"quizazz/internal/di"
	"quizazz/user-service/internal/application"
	"quizazz/user-service/internal/constants"
	"quizazz/user-service/userspb"
)

type serverTx struct {
	c di.Container
	userspb.UnimplementedUsersServiceServer
}

var _ userspb.UsersServiceServer = (*serverTx)(nil)

func RegisterServerTx(container di.Container, registrar grpc.ServiceRegistrar) error {
	userspb.RegisterUsersServiceServer(registrar, serverTx{
		c: container,
	})
	return nil
}

func (s serverTx) RegisterUser(ctx context.Context, request *userspb.RegisterUserRequest) (resp *userspb.RegisterUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}

	return next.RegisterUser(ctx, request)
}

func (s serverTx) AuthorizeUser(ctx context.Context, request *userspb.AuthorizeUserRequest) (resp *userspb.AuthorizeUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}

	return next.AuthorizeUser(ctx, request)
}

func (s serverTx) GetUser(ctx context.Context, request *userspb.GetUserRequest) (resp *userspb.GetUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}

	return next.GetUser(ctx, request)
}

func (s serverTx) EnableUser(ctx context.Context, request *userspb.EnableUserRequest) (resp *userspb.EnableUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}

	return next.EnableUser(ctx, request)
}

func (s serverTx) DisableUser(ctx context.Context, request *userspb.DisableUserRequest) (resp *userspb.DisableUserResponse, err error) {
	ctx = s.c.Scoped(ctx)
	defer func(tx *sql.Tx) {
		err = s.closeTx(tx, err)
	}(di.Get(ctx, constants.DatabaseTransactionKey).(*sql.Tx))

	next := server{app: di.Get(ctx, constants.ApplicationKey).(application.App)}

	return next.DisableUser(ctx, request)
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
