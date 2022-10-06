package main

import (
	"context"
	"os"

	"github.com/haleyrc/cheevos/auth"
	"github.com/haleyrc/cheevos/cheevos"
	"github.com/haleyrc/cheevos/lib/db/pg"
	"github.com/haleyrc/cheevos/lib/logger"
	"github.com/haleyrc/cheevos/roster"
	"github.com/haleyrc/cheevos/server"
)

func main() {
	ctx := context.Background()

	log := &logger.JSONLogger{
		EnableDebug: true,
		Output:      os.Stdout,
	}

	databaseURL := os.Getenv("DATABASE_URL")
	db, err := pg.Connect(ctx, databaseURL)
	if err != nil {
		panic(err)
	}

	srv := server.Server{
		Auth: server.AuthServer{
			Auth: &auth.Logger{
				Service: &auth.Service{
					DB:   db,
					Repo: &auth.Repository{},
				},
				Logger: log,
			},
		},
		Cheevos: server.CheevosServer{
			Cheevos: &cheevos.Logger{
				Service: &cheevos.Service{
					DB:   db,
					Repo: &cheevos.Repository{},
				},
				Logger: log,
			},
		},
		Roster: server.RosterServer{
			Roster: &roster.Logger{
				Service: &roster.Service{
					DB:   db,
					Repo: &roster.Repository{},
				},
				Logger: log,
			},
		},
	}

	if err := srv.Start(ctx); err != nil {
		log.Error(ctx, "server quit unexpectedly", err)
		os.Exit(1)
	}
}
