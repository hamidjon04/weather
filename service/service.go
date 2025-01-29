package service

import (
	"fmt"
	"log/slog"
	"weather/api/token"
	"weather/pkg/model"
	"weather/storage"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Storage storage.Storage
	Log     *slog.Logger
}

func NewService(storage storage.Storage, log *slog.Logger)Service{
	return Service{
		Storage: storage,
		Log: log,
	}
}

func (S *Service) Register(req *model.RegisterReq)(*model.RegisterResp, error){
	id, err := S.Storage.Users().Register(req)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is registration at service: %v", err))
		return nil, err
	}
	token, err := token.GenerateToken(id)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is generate token: %v", err))
		return nil, err
	}
	cheack, err := S.Storage.Users().CreateToken(&model.CreateTokenReq{
		UserId: id,
		Token: token.Token,
		ExpiresAt: token.ExpiresAt,
	})
	if err != nil || !cheack {
		S.Log.Error(fmt.Sprintf("Error is save token: %v", err))
		return nil, err
	}
	return &model.RegisterResp{
		AccessToken: token.Token,
	}, nil
}

func (S *Service) Login(req *model.LoginReq) (*model.RegisterResp, error){
	user, err := S.Storage.Users().GetUser(req.Username)
	if err != nil || len(user.Id) == 0{
		S.Log.Error(fmt.Sprintf("User is not registration: %v", err))
		return nil, fmt.Errorf("user is not registration")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil{
		S.Log.Error(fmt.Sprintf("Wrong password: %v", err))
		return nil, fmt.Errorf("wrong password")
	}
	token, err := token.GenerateToken(user.Id)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is generate token: %v", err))
		return nil, err
	}
	cheack, err := S.Storage.Users().CreateToken(&model.CreateTokenReq{
		UserId: user.Id,
		Token: token.Token,
		ExpiresAt: token.ExpiresAt,
	})
	if err != nil || !cheack {
		S.Log.Error(fmt.Sprintf("Error is save token: %v", err))
		return nil, err
	}
	return &model.RegisterResp{
		AccessToken: token.Token,
	}, nil
}