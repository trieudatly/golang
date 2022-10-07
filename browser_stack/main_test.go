package main

import (
	"fmt"
	"strconv"
	"testing"
)
var testStack Stack
func TestPushTableDriven(t *testing.T) {
	var tests = []struct {
		input  string
		want int
	}{
		{"facebook.com", 1},
		{"tuoitre.com", 2},
		{"thanhnien.com", 3},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.input)
		t.Run(testname, func(t *testing.T) {
			testStack.Push(tt.input)
			ans := testStack.Length()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	ans := testStack.Peek()
	if ans != "thanhnien.com" {
		t.Error("", ans)
	}
}
func TestLenght(t *testing.T) {
	ans := testStack.Length()
	if ans != 3 {
		t.Error("", ans)
	}
}
func TestPopTableDriven(t *testing.T) {
	var tests = []struct {
		wantUrl    string
		wantLength string
	}{
		{"thanhnien.com", "2"},
		{"tuoitre.com", "1"},
		{"facebook.com", "0"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.wantUrl)
		t.Run(testname, func(t *testing.T) {
			pop := testStack.Pop()
			lenght := testStack.Length()
			ans := pop + strconv.Itoa(lenght)
			want := tt.wantUrl + tt.wantLength
			if ans != want {
				t.Errorf("got %v, want %v", ans, want)
			}
		})
	}
}
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
