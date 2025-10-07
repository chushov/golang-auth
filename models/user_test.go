package models

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestHashPassword_Success(t *testing.T) {
    user := &User{}
    err := user.HashPassword("testPassword123")
    assert.NoError(t, err)
    assert.NotEmpty(t, user.Password)
    assert.NotEqual(t, "testPassword123", user.Password)
}

func TestHashPassword_EmptyPassword(t *testing.T) {
    user := &User{}
    err := user.HashPassword("")
    assert.Error(t, err)
}

func TestCheckPassword_Success(t *testing.T) {
    user := &User{}
    user.HashPassword("rightPass")
    err := user.CheckPassword("rightPass")
    assert.NoError(t, err)
}

func TestCheckPassword_WrongPassword(t *testing.T) {
    user := &User{}
    user.HashPassword("rightPass")
    err := user.CheckPassword("wrongPass")
    assert.Error(t, err)
}

