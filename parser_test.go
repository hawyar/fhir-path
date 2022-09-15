package main

import "testing"

func TestParser(t *testing.T) {
	want := "Patient.name.give"
	//expr := "Patient.name.exists()"
	//compelx := "Message.segment.where(code = 'PID').field[3].element.first().simple()"
	got, _ := Parse(want)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
