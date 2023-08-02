package database

import (
	"database/sql"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
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

func (s *ClientDBTestSuit) TestSave() {
	client := &entity.Client{
		ID:    "1",
		Name:  "John",
		Email: "abc@abcd.com",
	}
	err := s.clientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDBTestSuit) TestGet() {
	client, _ := entity.NewClient("John Doe", "j@j.com")
	s.clientDB.Save(client)

	clientDb, err := s.clientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDb.ID)
	s.Equal(client.Name, clientDb.Name)
	s.Equal(client.Email, clientDb.Email)
}
