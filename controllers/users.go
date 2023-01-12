package controllers

import (
	"context"
	db "users/grpc/database"
	"users/grpc/models"
	"users/grpc/pb"
)

type Server struct {
	*pb.UnimplementedUserServerServer
}

func (s *Server) GetAllUsers(in *pb.Empty,
	stream pb.UserServer_GetAllUsersServer) error {

	var usuarios []models.Users
	db.Database.Find(&usuarios)

	var users []*pb.Users

	for _, userModels := range usuarios {
		users = append(users, &pb.Users{Id: int64(userModels.Id), Username: userModels.Username, Passwd: userModels.Passwd, Email: userModels.Email})
	}

	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) GetUserById(ctx context.Context,
	in *pb.Id) (*pb.Users, error) {

	var usuario models.Users
	db.Database.First(&usuario, in.Id)

	res := &pb.Users{}

	var users []*pb.Users
	users = append(users, &pb.Users{Id: int64(usuario.Id), Username: usuario.Username, Passwd: usuario.Passwd, Email: usuario.Email})

	for _, users := range users {
		if users.GetId() == in.GetId() {
			res = users
			break
		}
	}
	return res, nil
}

func (s *Server) UpdateUser(ctx context.Context,
	in *pb.UpdateRequest) (*pb.Users, error) {

	var usuarioUpdate models.Users

	db.Database.First(&usuarioUpdate, in.Id)

	if usuarioUpdate.Id == 0 {
		return &pb.Users{Id: 0, Username: "", Passwd: "", Email: ""}, nil
	}

	if in.User.Username != "" {
		usuarioUpdate.Username = in.User.Username
	}

	if in.User.Passwd != "" {
		usuarioUpdate.Passwd = in.User.Passwd
	}

	if in.User.Email != "" {
		usuarioUpdate.Email = in.User.Email
	}

	db.Database.Save(&usuarioUpdate)
	res := pb.Users{Id: int64(usuarioUpdate.Id), Username: usuarioUpdate.Username, Passwd: usuarioUpdate.Passwd, Email: usuarioUpdate.Email}

	return &res, nil
}

func (s *Server) InsertUser(ctx context.Context,
	in *pb.Users) (*pb.Users, error) {

	var usuario models.Users
	usuario.Username = in.Username
	usuario.Passwd = in.Passwd
	usuario.Email = in.Email

	db.Database.Create(&usuario)
	res := pb.Users{Id: int64(usuario.Id), Username: usuario.Username, Passwd: usuario.Passwd, Email: usuario.Email}

	return &res, nil
}

func (s *Server) DeleteUser(ctx context.Context,
	in *pb.Id) (*pb.DeleteResponse, error) {

	var usuario models.Users
	db.Database.Delete(&usuario, in.Id)

	return &pb.DeleteResponse{Message: "Usuário deletado"}, nil
}
