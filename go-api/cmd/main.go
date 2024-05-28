package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/marumaro-git/study-go-api/go-api/domain/repository"
	"github.com/marumaro-git/study-go-api/go-api/handler"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// DB接続情報取得
	db := repository.SetUpDb()

	defer db.Close()

	srv := &handler.Server{
		DBAccessor: repository.DBAccessor{
			Db: db,
		},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/study", srv.StudyGetHandler)

	server := &http.Server{
		Addr:         ":8089",
		Handler:      mux,
		WriteTimeout: 2 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("server exited with error: %v", err)
		}
	}()

	<-ctx.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("server shutdown failed: %v", err)
	}

	log.Println("server exited properly")
}
