package userdb

import (
	"fmt"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	userTest := struct {
		username string
		password string
		email    string
	}{
		"username",
		"password",
		"email",
	}

	InsertUser(userTest.username, userTest.password, userTest.email)

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
	DeleteUserbByUsername(userTest.username)
}

func TestGetUserByID(t *testing.T) {
	userTest := struct {
		username string
		password string
		email    string
	}{
		"username",
		"password",
		"email",
	}

	user, err := GetUserByID(1)
	if err != nil {
		t.Error("fail to get user")
	}
	if user != nil {
		t.Error("user expected was nil")
	}

	//without error
	InsertUser(userTest.username, userTest.password, userTest.email)

	u, _ := GetUserByUsernameAndPassword(userTest.username, userTest.password)

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
	userTest := struct {
		username string
		password string
		email    string
	}{
		"username",
		"password",
		"email",
	}

	user, err := GetUserByUsernameAndPassword(userTest.username, userTest.password)
	if err != nil {
		t.Error("fail to get user")
	}
	if user != nil {
		t.Error("user expected was nil")
	}

	//without error
	InsertUser(userTest.username, userTest.password, userTest.email)

	user, err = GetUserByUsernameAndPassword(userTest.username, userTest.password)
	if err != nil {
		t.Error("fail to get user")
	}
	if user == nil {
		fmt.Println(user)
		t.Error("fail to get user")
	}

	DeleteUserbByUsername(userTest.username)
}

func TestCheckIfUserAlreadyExist(t *testing.T) {
	userTest := struct {
		username string
		password string
		email    string
	}{
		"username",
		"password",
		"email",
	}

	check, err := CheckIfUserAlreadyExist(userTest.username)
	if err != nil {
		t.Error("fail to get user")
	}
	if check {
		t.Error("check expected was false")
	}

	//without error
	InsertUser(userTest.username, userTest.password, userTest.email)

	check, err = CheckIfUserAlreadyExist(userTest.username)
	if err != nil {
		t.Error("fail to get user")
	}
	if !check {
		t.Error("check expected was true")
	}
	DeleteUserbByUsername(userTest.username)
}

func TestInsertUser(t *testing.T) {
	userTest := struct {
		username string
		password string
		email    string
	}{
		"username",
		"password",
		"email",
	}

	err := InsertUser(userTest.username, userTest.password, userTest.email)
	if err != nil {
		t.Error(err)
	}

	//with error
	err = InsertUser(userTest.username, userTest.password, userTest.email)
	if err == nil {
		t.Error("error was expected")
	}

	//with error
	Close()
	err = InsertUser(userTest.username, userTest.password, userTest.email)
	if err == nil {
		t.Error("error was expected")
	}
	Open()

	DeleteUserbByUsername(userTest.username)
}

func TestDeleteUserbByID(t *testing.T) {
	userTest := struct {
		username string
		password string
		email    string
	}{
		"username",
		"password",
		"email",
	}

	InsertUser(userTest.username, userTest.password, userTest.email)

	user, _ := GetUserByUsernameAndPassword(userTest.username, userTest.password)

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
	userTest := struct {
		username string
		password string
		email    string
	}{
		"username",
		"password",
		"email",
	}

	InsertUser(userTest.username, userTest.password, userTest.email)

	count, err := DeleteUserbByUsername(userTest.username)
	if err != nil {
		t.Error("error to delete user")
	}
	if count != 1 {
		t.Error("count expected was 1")
	}

	//with error
	count, err = DeleteUserbByUsername(userTest.username)
	if err != nil {
		t.Error("error to delete user")
	}
	if count != 0 {
		t.Error("count expected was 0")
	}

	//with error
	Close()
	count, err = DeleteUserbByUsername(userTest.username)
	if err == nil {
		t.Error("error was expected")
	}
	if count != 0 {
		t.Error("count expected was 0")
	}
	Open()
}
