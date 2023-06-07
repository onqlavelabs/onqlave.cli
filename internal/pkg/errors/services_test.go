package errors

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestDuplicateErr(t *testing.T) {
	ID := "duplicate-id"
	got := NewDuplicateRecord("user", ID)
	want := NewDuplicateRecord("user", ID)

	assert.Equal(t, got, want)
}
