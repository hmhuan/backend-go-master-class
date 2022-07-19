package api

import (
	"testing"

	db "github.com/hmhuan/simple-bank/db/sqlc"
	"github.com/hmhuan/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func randomUser(t *testing.T) db.User {
	password := util.RandomString(20)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	return db.User{
		Username: util.RandomOwner(),
		Password: hashedPassword,
		FullName: util.RandomString(20),
		Email:    util.RandomEmail(),
	}
}
