package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/jmoiron/sqlx"
	"github.com/tappn/grpc-chat/gen/go/room/v1"
	"github.com/tappn/grpc-chat/handler"
	"github.com/tappn/grpc-chat/repository"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// todo server option
	grpcServer := grpc.NewServer(opts...)
	db := Connect()
	roomRepo := repository.NewRoomRepository(db)
	room.RegisterRoomServiceServer(grpcServer, handler.NewRoomHandler(roomRepo))
	log.Println("------ start grpc server ------")
	grpcServer.Serve(listen)
}

func Connect() *sqlx.DB {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("fatal load for .env: %v", err.Error())
	}

	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	name := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		user, pass, host, name)

	for {
		db, err := sqlx.Connect("mysql", dsn)
		if err == nil {
			return db
		}
		log.Println(err)
		time.Sleep(1 * time.Second)
	}
}
