package main

import (
	"testing"
)

func TestCleanPath(t *testing.T) {
	for _, tt := range []struct {
		p    string
		want string
	}{
		{p: "foo/../bar/", want: "bar/index.html"},
		{p: "foo/../bar", want: "bar"},
		{p: "../../bar", want: "bar"},
		{p: "bar", want: "bar"},
		{p: "bar/", want: "bar/index.html"},
		{p: "/", want: "index.html"},
		{p: "", want: "index.html"},
		{p: "bar/../../", want: "index.html"},
		{p: "//bar/", want: "bar/index.html"},
	} {
		got := cleanPath(tt.p)
		if got != tt.want {
			t.Errorf("cleanPath(%s) = %s, want %s", tt.p, got, tt.want)
		}
	}
}
