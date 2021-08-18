package mycachedb

import (
	"testing"
)

func TestDB_GET(t *testing.T) {
	db := New()
	want := "b"
	db.SET("a", want)
	got := db.GET("a")
	if got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
	db.GET("12")
}
