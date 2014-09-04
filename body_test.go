package body

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var params = map[string][]string{
		"user[email]":      {"luis@vega.com"},
		"user[first_name]": {"Luis"},
		"user[last_name]":  {"Vega"},
		"user[age]":        {"27"},
		"user[cities][]":   {"Quezon City", "Chicago"},
		"user[ids][]":      {"1", "2", "3", "4"}}

	type SignupForm struct {
		Email     string   `name:"user[email]"`
		FirstName string   `name:"user[first_name]"`
		LastName  string   `name:"user[last_name]"`
		Age       int      `name:"user[age]"`
		Cities    []string `name:"user[cities][]"`
		Ids       []int64  `name:"user[ids][]"`
	}

	var signupForm = new(SignupForm)

	if err := Parse(params, signupForm); err != nil {
		t.Error("Failed!")
	}

	if signupForm.Email != "luis@vega.com" {
		t.Error("Failed!")
	}

	if signupForm.FirstName != "Luis" {
		t.Error("Failed!")
	}

	if signupForm.LastName != "Vega" {
		t.Error("Failed!")
	}

	if signupForm.Age != 27 {
		t.Error("Failed!")
	}

	if !reflect.DeepEqual(signupForm.Cities, []string{"Quezon City", "Chicago"}) {
		t.Error("Failed!")
	}

	if !reflect.DeepEqual(signupForm.Ids, []int64{1, 2, 3, 4}) {
		t.Error("Failed!")
	}
}
