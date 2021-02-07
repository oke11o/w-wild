package main

import (
	"context"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/go-chi/chi"

	grpcClient "github.com/oke11o/w-wild/internal/common/client"
	"github.com/oke11o/w-wild/internal/common/logs"
	"github.com/oke11o/w-wild/internal/common/server"
)

func main() {
	logs.Init()

	ctx := context.Background()
	fsClient, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	if err != nil {
		panic(err)
	}

	trainerClient, closeTrainerClient, err := grpcClient.NewTrainerClient()
	if err != nil {
		panic(err)
	}
	defer closeTrainerClient()

	usersClient, closeUsersClient, err := grpcClient.NewUsersClient()
	if err != nil {
		panic(err)
	}
	defer closeUsersClient()

	firebaseDB := db{fsClient}

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return HandlerFromMux(HttpServer{firebaseDB, trainerClient, usersClient}, router)
	})
}
