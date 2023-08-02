package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Dow", "test@teste.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Dow", "test@teste.com")
	account := NewAccount(client)
	account.Credit(50)
	assert.NotNil(t, account)
	assert.Equal(t, float64(50), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Dow", "test@teste.com")
	account := NewAccount(client)
	account.Credit(50)
	account.Debit(20)
	assert.NotNil(t, account)
	assert.Equal(t, float64(30), account.Balance)
}