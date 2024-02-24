package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"net/http"
	"os"
	"quizazz/internal/config"
	"quizazz/internal/system"
	"quizazz/internal/web"
	"quizazz/migrations"
	quiz_service "quizazz/quiz-service"
	"quizazz/user-service"
)

type monolith struct {
	*system.System
	modules []system.Module
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("mallbots exitted abnormally: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}
	s, err := system.NewSystem(cfg)
	if err != nil {
		return err
	}
	m := monolith{
		System: s,
		modules: []system.Module{
			&user_service.Module{},
			&quiz_service.Module{},
		},
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(m.DB())
	err = m.MigrateDB(migrations.FS)
	if err != nil {
		return err
	}

	if err = m.startupModules(); err != nil {
		return err
	}

	// Mount general web resources
	m.Mux().Mount("/", http.FileServer(http.FS(web.WebUI)))

	fmt.Println("started mallbots application")
	defer fmt.Println("stopped mallbots application")

	m.Waiter().Add(
		m.WaitForWeb,
		m.WaitForRPC,
		m.WaitForStream,
	)

	// go func() {
	// 	for {
	// 		var mem runtime.MemStats
	// 		runtime.ReadMemStats(&mem)
	// 		m.logger.Debug().Msgf("Alloc = %v  TotalAlloc = %v  Sys = %v  NumGC = %v", mem.Alloc/1024, mem.TotalAlloc/1024, mem.Sys/1024, mem.NumGC)
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }()

	return m.Waiter().Wait()
}

func (m *monolith) startupModules() error {
	for _, module := range m.modules {
		ctx := m.Waiter().Context()
		if err := module.Startup(ctx, m); err != nil {
			return err
		}
	}

	return nil
}
