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
			if tt.in.ID == 0 {
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
	for i, tt := range []struct {
		in  structure.User
		out structure.User
	}{
		{structure.User{}, structure.User{}},
		{structure.User{Username: "username", Password: "password", Email: "email"}, structure.User{Username: "username", Password: "password", Email: "email"}},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if tt.in.Username == "username" {
				err := InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
				defer DeleteUserbByUsername(tt.in.Username)
			}

			user, err := GetUserByUsernameAndPassword(tt.in.Username, tt.in.Password)
			if err != nil {
				t.Error(err)
			}

			if tt.in.Username != "username" {
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

func TestGetIDByUsername(t *testing.T) {
	for i, tt := range []struct {
		in  structure.User
		out int
	}{
		{structure.User{}, 0},
		{structure.User{Username: "username", Password: "password", Email: "email"}, 1},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if tt.in.Username == "username" {
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

			if tt.in.Username != "username" {
				if id != tt.out {
					t.Errorf("want %v; got %v", tt.out, id)
				}
			} else {
				if id < tt.out {
					t.Errorf("want %v; got %v", tt.out, id)
				}
			}
		})
	}
}

func TestCheckIfUserAlreadyExist(t *testing.T) {
	for i, tt := range []struct {
		in  structure.User
		out bool
	}{
		{structure.User{}, false},
		{structure.User{Username: "username", Password: "password", Email: "email"}, true},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if tt.in.Username == "username" {
				err := InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
				defer DeleteUserbByUsername(tt.in.Username)
			}

			check, err := CheckIfUserAlreadyExist(tt.in.Username)
			if err != nil {
				t.Error("fail to get user")
			}

			if check != tt.out {
				t.Errorf("want %v; got %v", tt.out, check)
			}
		})
	}
}

func TestInsertUser(t *testing.T) {
	for i, tt := range []struct {
		in  structure.User
		out string
	}{
		{structure.User{}, "database is closed"},
		{structure.User{Username: "username", Password: "password", Email: "email"}, ""},
		{structure.User{}, "duplicate key value"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if tt.out == "database is closed" {
				err := Close()
				if err != nil {
					t.Error(err)
				}
			}

			err := InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
			if err != nil {
				if !strings.Contains(err.Error(), tt.out) {
					t.Errorf("want %v; got %v", tt.out, err)
				}
			}

			if tt.out == "database is closed" {
				Open()
			}

			if tt.out != "" {
				_, err = DeleteUserbByUsername(tt.in.Username)
				if err != nil {
					t.Error(err)
				}
				_, err = DeleteUserbByUsername("username")
				if err != nil {
					t.Error(err)
				}
			}

		})
	}
}

func TestDeleteUserbByID(t *testing.T) {
	for i, tt := range []struct {
		in    structure.User
		out   int
		error string
	}{
		{structure.User{Username: "username", Password: "password", Email: "email"}, 1, ""},
		{structure.User{}, 0, ""},
		{structure.User{}, 0, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if tt.error == "database is closed" {
				err := Close()
				if err != nil {
					t.Error(err)
				}
				defer Open()
			}

			var id int

			if tt.in.Username == "username" {
				err := InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}

				id, err = GetIDByUsername(tt.in.Username)
				if err != nil {
					t.Error(err)
				}
			}

			count, err := DeleteUserbByID(id)
			if tt.out == 1 {
				if err != nil {
					t.Error(err)
				}
				if count != 1 {
					t.Errorf("want %v; got %v", tt.out, err)
				}
			}
		})
	}
}

func TestDeleteUserbByUsername(t *testing.T) {
	for i, tt := range []struct {
		in    structure.User
		out   int
		error string
	}{
		{structure.User{Username: "username", Password: "password", Email: "email"}, 1, ""},
		{structure.User{}, 0, ""},
		{structure.User{}, 0, "database is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if tt.error == "database is closed" {
				err := Close()
				if err != nil {
					t.Error(err)
				}
				defer Open()
			}

			if tt.in.Username == "username" {
				err := InsertUser(tt.in.Username, tt.in.Password, tt.in.Email)
				if err != nil {
					t.Error(err)
				}
			}

			count, err := DeleteUserbByUsername(tt.in.Username)
			if tt.out == 1 {
				if err != nil {
					t.Error(err)
				}
				if count != 1 {
					t.Errorf("want %v; got %v", tt.out, err)
				}
			}
		})
	}
}
