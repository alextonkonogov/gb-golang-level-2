package funcs_test

import (
	"testing"

	"github.com/alextonkonogov/gb-golang-level-2/homework7/task1/funcs"
	"github.com/alextonkonogov/gb-golang-level-2/homework7/task1/persons"
)

var tests = []struct {
	input map[string]interface{}
	want  persons.Person
}{
	{
		input: map[string]interface{}{
			"Name": "",
		},
		want: persons.Person{"", 0, false, 0},
	},
	{
		input: map[string]interface{}{
			"Name": "Michael",
		},
		want: persons.Person{"Michael", 0, false, 0},
	},
	{
		input: map[string]interface{}{
			"Name": "Ivan",
			"Age":  18,
		},
		want: persons.Person{"Ivan", 18, false, 0},
	},
	{
		input: map[string]interface{}{
			"Name":    "Susan",
			"Age":     25,
			"Married": false,
		},
		want: persons.Person{"Susan", 25, false, 0},
	},
	{
		input: map[string]interface{}{
			"Name":        "Alex",
			"Age":         33,
			"Married":     true,
			"Temperature": 36.6,
		},
		want: persons.Person{"Alex", 33, true, 36.6},
	},
}

func TestChangeStructField(t *testing.T) {
	for i, tt := range tests {
		person := persons.Person{}
		_ = funcs.ChangeStructField(&person, tt.input)
		if person != tt.want {
			t.Errorf("on index %v got %v want %v", i, person, tt.want)
		}
	}
}

func BenchmarkChangeStructField(b *testing.B) {
	m := map[string]interface{}{
		"Name":        "Alex",
		"Age":         33,
		"Married":     true,
		"Temperature": 36.6,
	}

	person := persons.Person{}

	for i := 0; i < b.N; i++ {
		funcs.ChangeStructField(&person, m)
	}
}
