package main

import (
	"errors"
	"fmt"
)

type ActionStrategy interface {
	Eat()
	Drink()
}

var actionStrategy = map[string]ActionStrategy{
	"ben":   NewBen(),
	"alice": NewAlice(),
}

func GetActionStrategy(name string) (ActionStrategy, error) {
	v, ok := actionStrategy[name]
	if !ok {
		return nil, errors.New("not found actionStrategy")
	}
	return v, nil
}

type Ben struct {
	Name string
}

func NewBen() *Ben {
	return &Ben{
		Name: "ben",
	}
}

func (e Ben) Eat() {
	fmt.Printf("%s is Eat", e.Name)
}

func (e Ben) Drink() {
	fmt.Printf("%s is Drink", e.Name)
}

type Alice struct {
	Name string
}

func NewAlice() *Alice {
	return &Alice{
		Name: "alice",
	}
}

func (a Alice) Eat() {
	fmt.Printf("%s is Eat", a.Name)
}

func (a Alice) Drink() {
	fmt.Printf("%s is Drink", a.Name)
}

func main() {
	strategy, err := GetActionStrategy("alice")
	if err != nil {
		return
	}
	strategy.Drink()
}
