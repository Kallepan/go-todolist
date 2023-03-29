package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kallepan/go-backend/db"
	"github.com/kallepan/go-backend/handler"
	"github.com/kallepan/go-backend/utils"
)

func Run() {
	utils.LoadEnv()
	dbInfo := utils.GetDbInfo()

	address := ":8080"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err.Error())
	}

	database, err := db.InitDB(dbInfo.User, dbInfo.Password, dbInfo.DbName, dbInfo.Host, dbInfo.Port)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err.Error())
	}
	defer database.CON.Close()

	httpHandler := handler.NewHandler(database)
	server := &http.Server{
		Handler: httpHandler,
	}

	go func() {
		server.Serve(listener)
	}()

	defer Stop(server)

	log.Printf("Server is running on %s", address)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprintf("Got signal: %s", <-ch))
	log.Println("Shutting down server...")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err.Error())
		os.Exit(1)
	}
}
