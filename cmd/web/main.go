package main

import (
	"github.com/bmerchant22/pkg/store"
	"github.com/bmerchant22/pkg/web"
	"log"
)

func main() {
	mongoStore := new(store.MongoStore)
	mongoStore.OpenConnectionWithMongoDB()
	srv := web.CreateWebServer(mongoStore)
	port := ":8080"
	if err := srv.StartListeningRequests(port); err != nil {
		log.Printf("Error occurred while starting the server: %v", err)
	}
	//worker.PerformWork(mongoStore)
}
