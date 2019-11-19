package util

import (
	"testing"
)

func TestParseToken(t *testing.T) {

	t.Run("test Generate ParseToken", func(t *testing.T){
		username := "wang"
		password := "123"
		token, err := GenerateToken(username, password)
		if err != nil {
			t.Errorf("%+v", err)
		}
		claims, err := ParseToken(token)
		if err != nil {
			t.Errorf("%+v", err)
		}
		if (claims.Username + "1" != username) {
			t.Errorf("get error username or password")
		}

		if (claims.Password !=  password) {
			t.Errorf("get error username or password")
		}
		
	})
}