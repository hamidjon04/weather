package storage

import (
	"database/sql"
	"log/slog"
	"weather/storage/postgres"
)

type Storage interface{
	Users()postgres.UsersRepo
}

type storageImpl struct{
	DB *sql.DB
	Log *slog.Logger
}

func NewStorage(db *sql.DB, log *slog.Logger)Storage{
	return &storageImpl{
		DB: db,
		Log: log,
	}
}

func(S *storageImpl) Users()postgres.UsersRepo{
	return postgres.NewUsersRepo(S.DB, S.Log)
}