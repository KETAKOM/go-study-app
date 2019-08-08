package main

import (
	"testing"

	"github.com/KETAKOM/go-study-app/application/usecase"
)

func TestBookIndex(t *testing.T) {

	b := usecase.BookUsecase{}
	result, err := b.GetAll()
	if err != nil {
		t.Fatalf("failed test%#v", err)
	}

	if result[0].Id != 1 {
		t.Fatal("failed test Id")
	}

	if result[0].Title != "Test1" {
		t.Fatal("failed test Title")
	}

	if result[0].Author != "Tester1" {
		t.Fatal("failed test Author")
	}
}
