package service

import (
	"database/sql"
	"log"
)

// DbService uses interface UserAccess to CRUD new users
type DbService struct {
	Db             *sql.DB
	UserService    UserAccess
	SessionService SessAccess
	l              *log.Logger
}

func NewDbService(Db *sql.DB, UserService UserAccess, SessionService SessAccess, logger *log.Logger) DbService {
	return DbService{Db, UserService, SessionService, logger}
}

// Get from the database
type Retrieval interface {
}
