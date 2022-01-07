package cache

import "testing"

/* func TestOpen(t *testing.T) {
} */

func TestClose(t *testing.T) {
	err := Close()
	if err != nil {
		t.Error("error to close cache")
	}
	err = Close()
	if err == nil {
		t.Error("want err; got nil")
	}

	Open()
}

func TestSetToken(t *testing.T) {
	valueTest := "test"

	err := SetToken(valueTest)
	if err != nil {
		t.Error("error to set token")
	}

	//with error
	err = Close()
	if err != nil {
		t.Fatal(err)
	}

	err = SetToken(valueTest)
	if err == nil {
		t.Error("want error; got nil")
	}

	Open()
}

func TestDeleteTokenUsingvalue(t *testing.T) {
	valueTest := "test"

	err := DeleteTokenUsingValue(valueTest)
	if err != nil {
		t.Error("error to delete token")
	}
}

func TestTokenIsValid(t *testing.T) {
	valueTest := "test"

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

	Open()
}
