package reflect_test

import (
	"fmt"
	. "github.com/mswatermelon/golang2/reflect"
	"testing"
)

type TestParents struct {
	Mom string
	Dad string
}

var personData = map[string]interface{}{
	"Name":    "Ivan",
	"Surname": "Ivanov",
	"Phone":   "Some international number",
	"Age":     45,
	"Parents": TestParents{Mom: "Anna", Dad: "Vladimir"},
	"Height":  172.5,
	"Merried": true,
}

func TestAssignToStringStruct(t *testing.T) {
	type TestPersonStruct struct {
		Name    string
		Surname string
		Phone   string
	}
	newPerson := TestPersonStruct{}
	if err := AssignToStruct(&newPerson, personData); err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(newPerson)
}

func TestAssignToStructWithVarFields(t *testing.T) {
	type TestPersonStruct struct {
		Name    string
		Surname string
		Phone   string
		Age     int
		Parents struct{}
		Height  float64
	}
	newPerson := TestPersonStruct{}
	if err := AssignToStruct(&newPerson, personData); err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(newPerson)
}

func TestAssignToUnexpectedField(t *testing.T) {
	type TestPersonStruct struct {
		Merried bool
	}
	newPerson := TestPersonStruct{}
	if err := AssignToStruct(&newPerson, personData); err == nil {
		t.Errorf(err.Error())
	}
}

func TestAssignToString(t *testing.T) {
	newStr := "newStr"
	if err := AssignToStruct(&newStr, personData); err == nil {
		t.Errorf(err.Error())
	}
}

func TestAssignToCustomStruct(t *testing.T) {
	type TestPersonStruct struct {
		Name    string
		Surname string
		Phone   string
		Age     int
		Parents TestParents
		Height  float64
	}
	newPerson := TestPersonStruct{}
	if err := AssignToStruct(&newPerson, personData); err == nil {
		t.Errorf(err.Error())
	}
}

func TestAssignToNil(t *testing.T) {
	if err := AssignToStruct(nil, personData); err == nil {
		t.Errorf(err.Error())
	}
}
