package main

import "testing"

func TestFooer(t *testing.T) {
	result := Fooer(5)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}

func TestFooerParallel(t *testing.T) {
	t.Run("Test 3 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(3)
		if result != "Foo" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
		}
	})
	t.Run("Test 7 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(7)
		if result != "7" {
			t.Errorf("Result was incorrect, got: %s, want: %s.", result, "7")
		}
	})
}

func TestFooer2(t *testing.T) {
	input := 3
	result := Fooer(5)
	t.Logf("The input was %d", input)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
	t.Fatalf("Stop the test")
	t.Error("won't be executed")
}

func TestFooerTableDriven(t *testing.T) {

	var tests = []struct {
		name  string
		input int
		want  string
	}{
		{"1 should be Foo", 1, "Foo"},
		{"2 should be Foo", 2, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"4 should be Foo", 4, "Foo"},
		{"5 should be Foo", 5, "Foo"},
		{"6 should be Foo", 6, "Foo"},
		{"9 should be Foo", 7, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 0, "1"},
		{"0 should be Foo", 0, "Foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Fooer(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
