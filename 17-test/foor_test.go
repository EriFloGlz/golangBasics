package test

import "testing"

func TestFooer(t *testing.T) {
	result := Foo(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s", result, "Foo")
	}
}

func TestFooTableDrive(t *testing.T) {
	//Defining the columns of the table
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		// the tabl itself
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 shoul be Foo", 0, "Foo"},
	}
	//the execution
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Foo(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
