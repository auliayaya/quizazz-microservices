package application

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
)

type instrumentedApp struct {
	App
	usersRegistered prometheus.Counter
}

var _ App = (*instrumentedApp)(nil)

func NewInstrumentedApp(app App, usersRegistered prometheus.Counter) App {
	return instrumentedApp{
		App:             app,
		usersRegistered: usersRegistered,
	}
}

func (a instrumentedApp) RegisterUser(ctx context.Context, register RegisterUser) error {
	err := a.App.RegisterUser(ctx, register)
	if err != nil {
		return err
	}
	a.usersRegistered.Inc()
	return nil
}
