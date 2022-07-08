package db

import (
	"context"
	"testing"

	"github.com/hmhuan/backend-go-master-class/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashPassword, err := util.HashPassword(util.RandomString(20))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username: util.RandomOwner(),
		Password: hashPassword,
		FullName: util.RandomOwner(),
		Email:    util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	user1, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user.Username, user1.Username)
	require.Equal(t, user.Email, user1.Email)
	require.Equal(t, user.FullName, user1.FullName)
	require.Equal(t, user.Password, user1.Password)
}
