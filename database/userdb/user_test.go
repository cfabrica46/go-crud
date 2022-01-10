package userdb

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cfabrica46/go-crud/structure"
)

func TestGetAllUsers(t *testing.T) {
	for i, tt := range []struct {
		in  structure.User
		out string
	}{
		{structure.User{Username: "username", Password: "password", Email: "email"}, ""},
		{structure.User{Username: "username", Password: "password", Email: "email"}, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if tt.out == "database is closed" {
				err := Close()
				if err != nil {
					t.Error(err)
				}
				defer Open()
			} else if tt.out == "" {
				err := InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
				defer DeleteUserbByUsername(tt.in.Username)
			}

			_, err := GetAllUsers()
			if err != nil {
				if !strings.Contains(err.Error(), tt.out) {
					t.Errorf("want %v; got %v", tt.out, err)
				}
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	for i, tt := range []struct {
		in  structure.User
		out structure.User
	}{
		{structure.User{ID: -1}, structure.User{}},
		{structure.User{Username: "username", Password: "password", Email: "email"}, structure.User{Username: "username", Password: "password", Email: "email"}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if tt.in.ID != -1 {
				err := InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
				defer DeleteUserbByUsername(tt.in.Username)
			}

			id, err := GetIDByUsername(tt.in.Username)
			if err != nil {
				t.Error(err)
			}

			user, err := GetUserByID(id)
			if err != nil {
				t.Error(err)
			}
			if tt.in.ID == -1 {
				if user != nil {
					t.Errorf("want %v; got %v", tt.out, user)
				}
			} else {
				if user == nil {
					t.Errorf("want %v; got %v", tt.out, user)
				}
			}
		})
	}
}

func TestGetUserByUsernameAndPassword(t *testing.T) {
	userTest := structure.User{
		Username: "username",
		Password: "password",
		Email:    "email",
	}

	user, err := GetUserByUsernameAndPassword(userTest.Username, userTest.Password)
	if err != nil {
		t.Error("fail to get user")
	}
	if user != nil {
		t.Error("user expected was nil")
	}

	//without error
	InsertUser(userTest.Username, userTest.Password, userTest.Email)

	user, err = GetUserByUsernameAndPassword(userTest.Username, userTest.Password)
	if err != nil {
		t.Error("fail to get user")
	}
	if user == nil {
		t.Error("fail to get user")
	}

	DeleteUserbByUsername(userTest.Username)
}

func TestGetIDByUsername(t *testing.T) {
	userTest := structure.User{
		Username: "username",
		Password: "password",
		Email:    "email",
	}

	InsertUser(userTest.Username, userTest.Password, userTest.Email)

	id, err := GetIDByUsername(userTest.Username)
	if err != nil {
		fmt.Println(err)
		t.Error("fail to get ID")
	}
	if id <= 0 {
		t.Error("fail to get ID")
	}
	DeleteUserbByUsername(userTest.Username)

	//without results
	id, err = GetIDByUsername(userTest.Username)
	if err != nil {
		fmt.Println(err)
		t.Error("fail to get ID")
	}
	if id != 0 {
		t.Errorf("sant 0; got %d", id)
	}
}

func TestCheckIfUserAlreadyExist(t *testing.T) {
	userTest := structure.User{
		Username: "username",
		Password: "password",
		Email:    "email",
	}

	check, err := CheckIfUserAlreadyExist(userTest.Username)
	if err != nil {
		t.Error("fail to get user")
	}
	if check {
		t.Error("check expected was false")
	}

	//without error
	InsertUser(userTest.Username, userTest.Password, userTest.Email)

	check, err = CheckIfUserAlreadyExist(userTest.Username)
	if err != nil {
		t.Error("fail to get user")
	}
	if !check {
		t.Error("check expected was true")
	}
	DeleteUserbByUsername(userTest.Username)
}

func TestInsertUser(t *testing.T) {
	userTest := structure.User{
		Username: "username",
		Password: "password",
		Email:    "email",
	}

	err := InsertUser(userTest.Username, userTest.Password, userTest.Email)
	if err != nil {
		t.Error(err)
	}

	//with error
	err = InsertUser(userTest.Username, userTest.Password, userTest.Email)
	if err == nil {
		t.Error("error was expected")
	}

	//with error
	Close()
	err = InsertUser(userTest.Username, userTest.Password, userTest.Email)
	if err == nil {
		t.Error("error was expected")
	}
	Open()

	DeleteUserbByUsername(userTest.Username)
}

func TestDeleteUserbByID(t *testing.T) {
	userTest := structure.User{
		Username: "username",
		Password: "password",
		Email:    "email",
	}

	InsertUser(userTest.Username, userTest.Password, userTest.Email)

	user, _ := GetUserByUsernameAndPassword(userTest.Username, userTest.Password)

	count, err := DeleteUserbByID(user.ID)
	if err != nil {
		t.Error("error to delete user")
	}
	if count != 1 {
		t.Error("count expected was 1")
	}

	//with error
	count, err = DeleteUserbByID(user.ID)
	if err != nil {
		t.Error("error to delete user")
	}
	if count != 0 {
		t.Error("count expected was 0")
	}

	//with error
	Close()
	count, err = DeleteUserbByID(user.ID)
	if err == nil {
		t.Error("error was expected")
	}
	if count != 0 {
		t.Error("count expected was 0")
	}
	Open()
}

func TestDeleteUserbByUsername(t *testing.T) {
	userTest := structure.User{
		Username: "username",
		Password: "password",
		Email:    "email",
	}

	InsertUser(userTest.Username, userTest.Password, userTest.Email)

	count, err := DeleteUserbByUsername(userTest.Username)
	if err != nil {
		t.Error("error to delete user")
	}
	if count != 1 {
		t.Error("count expected was 1")
	}

	//with error
	count, err = DeleteUserbByUsername(userTest.Username)
	if err != nil {
		t.Error("error to delete user")
	}
	if count != 0 {
		t.Error("count expected was 0")
	}

	//with error
	Close()
	count, err = DeleteUserbByUsername(userTest.Username)
	if err == nil {
		t.Error("error was expected")
	}
	if count != 0 {
		t.Error("count expected was 0")
	}
	Open()
}
