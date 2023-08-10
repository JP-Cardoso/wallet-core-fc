package create_client

import (
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// type ClientGatewayMock struct {
// 	mock.Mock
// }

// func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
// 	args := m.Called(id) //verifica se foi chamadoo id
// 	return args.Get(0).(*entity.Client), args.Error(1)
// }

// func (m *ClientGatewayMock) Save(client *entity.Client) error {
// 	args := m.Called(client) //verifica se foi chamadoo id
// 	return args.Error(0)
// }

// /*
// estamos fazendo um teste de unidade, mocando o DB com o ClientGatewayMock,
// para testar o método save da interface
// */
// func TestCreateClientUseCase_Execute(t *testing.T) {
// 	m := &ClientGatewayMock{}
// 	m.On("Save", mock.Anything).Return(nil) //Ele vai chamar o save e retornar nil
// 	uc := NewCreateClientUseCase(m)
	
// 	output, err := uc.Execute(CreateClientInputDto{
// 		Name: "Joao",
// 		Email: "test@teste.com",
// 	})
// 	assert.Nil(t, err)
// 	assert.NotNil(t, output)
// 	assert.NotEmpty(t, output.ID)
// 	assert.Equal(t, "Joao", output.Name)
// 	assert.Equal(t, "test@teste.com", output.Email)
// 	m.AssertExpectations(t) //garante que o save foi chamado
// 	m.AssertNumberOfCalls(t, "Save", 1) //Faz a contagem de quantas vezes o método foi chamado
// }

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMock{}
	m.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(CreateClientInputDto{
		Name:  "John Doe",
		Email: "j@j",
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "j@j", output.Email)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}