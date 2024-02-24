package application

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
)

type instrumentedApp struct {
	App
	quizzesRegistered prometheus.Counter
}

var _ App = (*instrumentedApp)(nil)

func NewInstrumentedApp(app App, quiz prometheus.Counter) App {
	return instrumentedApp{
		App:               app,
		quizzesRegistered: quiz,
	}
}

func (a instrumentedApp) RegisterUser(ctx context.Context, register CreateQuiz) error {
	err := a.App.CreateQuiz(ctx, register)
	if err != nil {
		return err
	}
	a.quizzesRegistered.Inc()
	return nil
}
