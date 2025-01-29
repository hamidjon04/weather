package postgres

import (
	"database/sql"
	"fmt"
	"log/slog"
	"weather/pkg/model"

	"github.com/google/uuid"
)

type UsersRepo interface {
	Register(req *model.RegisterReq) (string, error)
	CreateToken(req *model.CreateTokenReq) (bool, error)
	GetToken(userId string) (*model.GetTokenResp, error)
}

type usersImpl struct {
	DB  *sql.DB
	Log *slog.Logger
}

func NewUsersRepo(db *sql.DB, log *slog.Logger) UsersRepo {
	return &usersImpl{
		DB:  db,
		Log: log,
	}
}

func (U *usersImpl) Register(req *model.RegisterReq) (string, error) {
	id := uuid.NewString()
	query := `
				INSERT INTO users(
					id, name, surname, username, password)
				VALUES
					($1, $2, $3, $4, $5)`
	_, err := U.DB.Exec(query, id, req.Name, req.Surname, req.Username, req.Password)
	if err != nil {
		U.Log.Error(fmt.Sprintf("Error is doing registration: %v", err))
		return "", err
	}
	return id, nil
}

func (U *usersImpl) CreateToken(req *model.CreateTokenReq) (bool, error) {
	query := `
				INSERT INTO tokens(
					user_id, token, expires_at)
				VALUES
					($1, $2, $3)`
	_, err := U.DB.Exec(query, req.UserId, req.Token, req.ExpiresAt)
	if err != nil {
		U.Log.Error(fmt.Sprintf("Error is saving token: %v", err))
		return false, err
	}
	return true, nil
}

func (U *usersImpl) GetToken(userId string) (*model.GetTokenResp, error) {
	resp := model.GetTokenResp{}
	query := `
				SELECT 
					token, expires_at
				FROM
					tokens
				WHERE 
					user_id = $1`
	err := U.DB.QueryRow(query, userId).Scan(&resp.Token, &resp.ExpiresAt)
	if err != nil{
		U.Log.Error(fmt.Sprintf("Error is getting token: %v", err))
		return nil, err
	}
	return &resp, nil
}
