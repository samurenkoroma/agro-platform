package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/samurenkoroma/agro-platform/internal/bootstrap"
	configs "github.com/samurenkoroma/agro-platform/internal/shared/config"
	"github.com/samurenkoroma/agro-platform/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	ctx := context.Background()
	pool, err := db.NewPool(conf.Db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer pool.Close()
	//
	app, err := bootstrap.Build(ctx, pool, conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server started on :8080")

	if err := http.ListenAndServe(":8080", app.HTTPHandler); err != nil {
		log.Fatal(err)
	}

}
