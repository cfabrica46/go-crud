package userdb

import (
	"fmt"
	"testing"

	"github.com/cfabrica46/go-crud/structure"
)

func TestGetAllUsers(t *testing.T) {
	userTest := structure.User{
		Username: "username",
		Password: "password",
		Email:    "email",
	}

	InsertUser(userTest.Username, userTest.Password, userTest.Email)

	users, err := GetAllUsers()
	if err != nil {
		t.Error("Fail to get all users")
	}
	if len(users) < 1 {
		t.Error("Fail to get all users")
	}

	//error
	Close()
	_, err = GetAllUsers()
	if err == nil {
		t.Error("error was expected")
	}
	Open()
	DeleteUserbByUsername(userTest.Username)
}

func TestGetUserByID(t *testing.T) {
	userTest := structure.User{
		Username: "username",
		Password: "password",
		Email:    "email",
	}

	user, err := GetUserByID(1)
	if err != nil {
		t.Error("fail to get user")
	}
	if user != nil {
		t.Error("user expected was nil")
	}

	//without error
	InsertUser(userTest.Username, userTest.Password, userTest.Email)

	u, _ := GetUserByUsernameAndPassword(userTest.Username, userTest.Password)

	user, err = GetUserByID(u.ID)
	if err != nil {
		t.Error("fail to get user")
	}
	if user == nil {
		fmt.Println(user)
		t.Error("fail to get user")
	}
	DeleteUserbByID(u.ID)
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
