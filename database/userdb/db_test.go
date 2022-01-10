package userdb

import (
	"fmt"
	"strings"
	"testing"
)

func TestOpen(t *testing.T) {
	for i, tt := range []struct {
		out string
	}{
		{""},
		{"unknown driver"},
		{"connection refused"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			var aux string
			switch tt.out {
			case "unknown driver":
				aux = dbDriver
				dbDriver = ""
				defer func() { dbDriver = aux }()
			case "connection refused":
				aux = psqlInfo
				psqlInfo = ""
				defer func() { psqlInfo = aux }()
			}

			_, err := Open()
			if err != nil {
				if !strings.Contains(err.Error(), tt.out) {
					t.Errorf("want %s; got %s", tt.out, err)
				}
			}
		})
	}
}

func TestClose(t *testing.T) {
	for i, tt := range []struct {
		out string
	}{
		{""},
		{"database already close"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			switch tt.out {
			case "database already close":
				db = nil
				defer func() {
					_, err := Open()
					if err != nil {
						t.Error(err)
					}
				}()
			}

			err := Close()
			if err != nil {
				if !strings.Contains(err.Error(), tt.out) {
					t.Errorf("want %v; got %v", tt.out, err)
				}
			}
		})
	}
}
