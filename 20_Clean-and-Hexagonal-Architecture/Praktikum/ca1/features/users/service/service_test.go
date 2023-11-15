package service

import (
	"errors"
	"restEcho1/features/users"
	"restEcho1/features/users/mocks"
	helper "restEcho1/helper/mocks"

	"github.com/stretchr/testify/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
    generator := helper.NewGeneratorInterface(t)
    jwt := helper.NewJWTInterface(t)
    data := mocks.NewUserDataInterface(t)
    service := New(data, generator, jwt)
    newUser := users.User{
        Nama:     "yusuf",
        HP:       "12345",
        Password: "mantul123",
    }

    t.Run("Success insert", func(t *testing.T) {
        generator.On("GenerateUUID").Return("randomUUID", nil).Once()
        newUser.ID = "randomUUID"
        data.On("Insert", newUser).Return(&newUser, nil).Once()

        result, err := service.Register(newUser)
        assert.Nil(t, err)
        assert.Equal(t, newUser.ID, result.ID)
        assert.Equal(t, newUser.Nama, result.Nama)
        generator.AssertExpectations(t)
        data.AssertExpectations(t)
    })

    t.Run("Generate failed", func(t *testing.T) {
        generator.On("GenerateUUID").Return("", errors.New("some error on generator")).Once()

        result, err := service.Register(newUser)
        assert.Error(t, err)
        assert.EqualError(t, err, "id generator failed")
        assert.Nil(t, result)
        generator.AssertExpectations(t)
    })

    t.Run("Insert failed", func(t *testing.T) {
        generator.On("GenerateUUID").Return("randomUUID", nil).Once()
        newUser.ID = "randomUUID"
        data.On("Insert", newUser).Return(nil, errors.New("some error on insert")).Once()

        result, err := service.Register(newUser)
        assert.Error(t, err)
        assert.EqualError(t, err, "insert process failed")
        assert.Nil(t, result)
        generator.AssertExpectations(t)
        data.AssertExpectations(t)
    })
}


func TestLogin(t *testing.T) {
	generator := helper.NewGeneratorInterface(t)
	j := helper.NewJWTInterface(t)
	data := mocks.NewUserDataInterface(t)
	service := New(data, generator, j)
	userData := users.User{
		ID:       "randomUserID",
		Nama:     "yusuf",
		HP:       "12345",
		Password: "mantul123",
	}

	t.Run("success login", func(t *testing.T) {
		jwtResult := map[string]any{"access_token": "randomAccessToken"}
		data.On("Login", userData.HP, userData.Password).Return(&userData, nil)
		j.On("GenerateJWT", mock.Anything).Return(jwtResult)
		result, err := service.Login(userData.HP, userData.Password)

		data.AssertExpectations(t)
		j.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "yusuf", result.Nama)
		assert.Equal(t, jwtResult, result.Access)
	})
}

func TestGetAllUsers(t *testing.T) {
    data := mocks.NewUserDataInterface(t)
    generator := helper.NewGeneratorInterface(t)
    jwt := helper.NewJWTInterface(t)
    service := New(data, generator, jwt)

    t.Run("Success - Data Found", func(t *testing.T) {
        usersData := []users.User{
            {ID: "1", Nama: "User1", HP: "12345", Password: "pass1"},
            {ID: "2", Nama: "User2", HP: "67890", Password: "pass2"},
        }

        data.On("GetAllUsers").Return(usersData, nil).Once()

        result, err := service.GetAllUsers()

        assert.Nil(t, err)
        assert.Equal(t, usersData, result)
        data.AssertExpectations(t)
    })

    t.Run("Data Not Found", func(t *testing.T) {
        data.On("GetAllUsers").Return(nil, errors.New("not found")).Once()

        result, err := service.GetAllUsers()

        assert.Error(t, err)
        assert.EqualError(t, err, "data not found")
        assert.Nil(t, result)
        data.AssertExpectations(t)
    })

    t.Run("Process Failed", func(t *testing.T) {
        data.On("GetAllUsers").Return(nil, errors.New("some error")).Once()

        result, err := service.GetAllUsers()

        assert.Error(t, err)
        assert.EqualError(t, err, "process failed")
        assert.Nil(t, result)
        data.AssertExpectations(t)
    })
}
