package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuit struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

// Esse cara vai ser executado toda vez que rodarmos um teste
func (s *ClientDBTestSuit) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuit) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDbTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuit))
	/*
		Quando esse cara rodar, ele executara todas as "funções"
		que estão incubidas na nossa suite de DB
	*/
}
