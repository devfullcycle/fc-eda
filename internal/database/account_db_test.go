package database

import (
	"database/sql"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
    suite.Suite
    db *sql.DB
    accountDB *AccountDB
    client *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
    db, err := sql.Open("sqlite3", ":memory:")
    s.Nil(err)
    s.db = db
    db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
    s.accountDB = NewAccountDB(db)
    s.client, _ = entity.NewClient("John", "j@j.com")    
}

func (s *AccountDBTestSuite) TearDownSuite() {
    defer s.db.Close()
    s.db.Exec("DROP TABLE clients")
    s.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
    suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
    account := entity.NewAccount(s.client)
    err := s.accountDB.Save(account)
    s.Nil(err)
}

func (s *AccountDBTestSuite) TestFindByID() {
    s.db.Exec("Insert into clients (id, name, email, created_at) values (?, ?, ?, ?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt,
	)
    account := entity.NewAccount(s.client)
    err := s.accountDB.Save(account)
    s.Nil(err)
    accountDB, err := s.accountDB.FindByID(account.ID)
    s.Nil(err)
    s.Equal(account.ID, accountDB.ID)
    s.Equal(account.ClientID, accountDB.ClientID)
    s.Equal(account.Balance, accountDB.Balance)
    s.Equal(account.Client.ID, accountDB.Client.ID)
    s.Equal(account.Client.Name, accountDB.Client.Name)
    s.Equal(account.Client.Email, accountDB.Client.Email)
}