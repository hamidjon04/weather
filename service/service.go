package service

import (
	"fmt"
	"log/slog"
	"weather/api/token"
	"weather/pkg/model"
	"weather/storage"
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
	accessToken, err := token.GenerateToken(id)
	if err != nil{
		S.Log.Error(fmt.Sprintf("Error is generate token: %v", err))
		return nil, err
	}
	return &model.RegisterResp{
		AccessToken: accessToken,
	}, nil
}