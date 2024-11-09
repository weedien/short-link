package main

import (
	"github.com/jinzhu/copier"
	"shortlink/internal/base/types"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

type People struct {
	Locale  string
	Persons []Person
}

type Student struct {
	Name string
	Age  int
}

type Class struct {
	Locale   string
	Students []Student `copier:"Persons"`
}

func TestDeepCopyArray(t *testing.T) {
	people := People{
		Locale: "zh-CN",
		Persons: []Person{
			{
				Name: "Alice",
				Age:  20,
			},
			{
				Name: "Bob",
				Age:  21,
			},
		},
	}
	class := Class{}
	if err := copier.Copy(&class, &people); err != nil {
		t.Errorf("copy error: %v", err)
	}
	t.Logf("people: %+v", people)
	t.Logf("class: %+v", class)
}

func TestDeepCopyPage(t *testing.T) {
	personPage := types.PageResp[Person]{
		Current: 1,
		Size:    10,
		Total:   10,
		Records: []Person{
			{
				Name: "Alice",
				Age:  20,
			},
			{
				Name: "Bob",
				Age:  21,
			},
		},
	}
	studentPage := types.PageResp[Student]{}
	if err := copier.Copy(&studentPage, &personPage); err != nil {
		t.Errorf("copy error: %v", err)
	}
	t.Logf("personPage: %+v", personPage)
	t.Logf("studentPage: %+v", studentPage)
}
