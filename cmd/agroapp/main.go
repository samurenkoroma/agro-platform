package main

import (
	"context"
	"log"
	"net/http"

	configs "github.com/samurenkoroma/agro-platform/internal/shared/config"
	"github.com/samurenkoroma/agro-platform/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	ctx := context.Background()

	conn, err := db.NewDB(ctx, conf.Db.Dsn)
	if err != nil {
		return
	}
	defer conn.Close()

	app, err := c.Build(ctx, conn, conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server started on :8080")

	if err := http.ListenAndServe(":8080", app.HTTPHandler); err != nil {
		log.Fatal(err)
	}

}
