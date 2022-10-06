package main

import (
	"fmt"
	"testing"
)

func TestCurrentPageTableDriven(t *testing.T) {
	var tests = []struct {
		url  string
		want string
	}{
		{"facebook.com", "facebook.com"},
		{"tuoitre.com", "tuoitre.com"},
		{"thanhnien.com", "thanhnien.com"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.url)
		t.Run(testname, func(t *testing.T) {
			Visit(tt.url)
			ans := CurrentPage()
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
func TestBack(t *testing.T) {
	ans := Back()
	if ans != "tuoitre.com" {
		t.Error("", ans)
	}
}
func TestForward(t *testing.T) {
	ans := Forward()
	if ans != "thanhnien.com" {
		t.Error("", ans)
	}
}
