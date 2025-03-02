package designauthenticationmanager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthManager_1(t *testing.T) {
	obj := Constructor(5)
	obj.Renew("aaa", 1)
	obj.Generate("aaa", 2)
	assert.Equal(t, 1, obj.CountUnexpiredTokens(6))
	obj.Generate("bbb", 7)
	obj.Renew("aaa", 8)
	obj.Renew("bbb", 10)
	assert.Equal(t, 1, obj.CountUnexpiredTokens(11))
}

func TestAuthManager_2(t *testing.T) {
	obj := Constructor(5)
	obj.Renew("aaa", 1)
	obj.Generate("aaa", 2)
	assert.Equal(t, 1, obj.CountUnexpiredTokens(6))
	obj.Generate("bbb", 7)
	obj.Renew("aaa", 8)
	obj.Renew("bbb", 10)
	assert.Equal(t, 0, obj.CountUnexpiredTokens(15))
}
