package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"

	"gokit-example/account"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "anhtuan"
	psswd  = "abc13579"
	dbname = "gokit_example"
)

//const dbsource = "postgresql://postgres:postgres@localhost:8888/gokitexample?sslmode=disable"

func main() {

	httpAddr := flag.String("http", ":8080", "http listen address")
	// psswd := flag.String("password", "unknown", "Password of postgreSQL database")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, psswd, dbname)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", psqlconn)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}

	defer db.Close()

	flag.Parse()
	ctx := context.Background()

	// set channel for good exit and return errors
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	var srv account.Service
	{
		// initialize repository
		repository := account.NewRepository(db, logger)
		// initialize service
		srv = account.NewService(repository, logger)
	}

	// initialize Endpoints
	endpoints := account.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints, logger)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
