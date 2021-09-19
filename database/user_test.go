package database

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func Test_createUser(t *testing.T) {
	password, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	user := &User{
		Username: "admin",
		Password: string(password),
	}
	t.Log(InsertUser(user))
}

func Test_login(t *testing.T) {
	user, err := FindUserByUserName("admin")
	if err != nil || user.ID == 0 {
		t.Fatal(err)
	}
	t.Log(bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("admin")))
}

func Test_update(t *testing.T) {
	user, err := FindUserByUserName("admin")
	if err != nil || user.ID == 0 {
		t.Fatal(err)
	}
	password, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	user.Password = string(password)
	t.Log(UpdateUser(user))
}
