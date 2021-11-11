package main

import (
	"esme_team/src/db/mongodb"
	"esme_team/src/db/postgresdb"
	"esme_team/src/model"
	"esme_team/src/pbs/reportpb"
	"esme_team/src/pbs/userpb"
	"esme_team/src/server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	postgresdb.PostgresInit()
	if err := postgresdb.DB.AutoMigrate(&model.UserModel{},&model.TokenModel{}); err != nil {
		log.Fatalf("error while AutoMigrate : %v" , err)
	}
	mongodb.Connect()

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("error while listening to host: %v\n", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)

	serviceRegistry(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error while serving the server : %v", err)
	}
}

func serviceRegistry(s *grpc.Server) {
	reportpb.RegisterReportServiceServer(s , &server.ReportServer{
		Collection: mongodb.MongoClient.Database("server_db").Collection("reports"),
	})
	userpb.RegisterUserServiceServer(s , &server.UserServer{})
}
