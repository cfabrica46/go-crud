package cache

import (
	"fmt"
	"strings"
	"testing"
)

func TestClose(t *testing.T) {
	for i, tt := range []struct {
		out string
	}{
		{""},
		{"client is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			err := Close()
			if err != nil {
				if !strings.Contains(err.Error(), tt.out) {
					t.Errorf("want %v; got %v", tt.out, err)
				}
				Open()
			}
		})
	}
}

func TestSetToken(t *testing.T) {
	for i, tt := range []struct {
		in  string
		out string
	}{
		{"token", ""},
		{"", "client is closed"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if tt.in == "" {
				err := Close()
				if err != nil {
					t.Error(err)
				}
				defer Open()
			}

			err := SetToken(tt.in)
			if err != nil {
				if !strings.Contains(err.Error(), tt.out) {
					t.Errorf("want %v; got %v", tt.out, err)
				}
			}
		})
	}
}

func TestDeleteTokenUsingvalue(t *testing.T) {
	valueTest := "test"

	err := DeleteTokenUsingValue(valueTest)
	if err != nil {
		t.Error("error to delete token")
	}
}

func TestTokenIsValid(t *testing.T) {
	for i, tt := range []struct {
		in  string
		out bool
	}{
		{"", false},
		{"close", false},
		{"token", true},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			if tt.in == "close" {
				err := Close()
				if err != nil {
					t.Error(err)
				}
				defer Open()
			}

			if tt.in == "token" {
				err := SetToken(tt.in)
				if err != nil {
					t.Error(err)
				}
			}

			check, err := TokenIsValid(tt.in)
			if tt.in != "close" {
				if err != nil {
					t.Error(err)
				}
			}
			if check != tt.out {
				t.Errorf("want %v; got %v", !check, check)
			}
		})
	}
	/* valueTest := "test"

	check, err := TokenIsValid(valueTest)
	if err != nil {
		t.Error("error to check token")
	}
	if check {
		t.Errorf("want %v; got %v", !check, check)
	}

	//without error
	err = SetToken(valueTest)
	if err != nil {
		t.Fatal(err)
	}

	check, err = TokenIsValid(valueTest)
	if err != nil {
		t.Error("error to check token")
	}
	if !check {
		t.Errorf("want %v; got %v", !check, check)
	}

	//with error
	err = Close()
	if err != nil {
		t.Fatal(err)
	}

	check, err = TokenIsValid(valueTest)
	if err == nil {
		t.Error("want error; got nil")
	}
	if check {
		t.Errorf("want %v; got %v", !check, check)
	}

	Open() */
}
