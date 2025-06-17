package main

import (
	"database/sql"
	"net"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/14-grpc/internal/database"
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/14-grpc/internal/pb"
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/14-grpc/internal/service"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Abre a conexão com o banco de dados SQLite
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Cria a instância do banco de dados de categoria e o serviço de categoria
	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	// Cria o servidor gRPC e registra o serviço de categoria
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	// Permite que clientes descubram os serviços disponíveis
	// Em produção, isso pode ser desativado por questões de segurança
	reflection.Register(grpcServer)

	// Inicia o servidor gRPC na porta 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	// Serve o servidor gRPC
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}

}
