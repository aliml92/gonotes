package main

import "fmt"

type model struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type Service interface {
	add(name string) error
	remove(id int) error
	getAll() ([]model, error)
}

func main(){
	fmt.Println("Hello World")
}

type svc struct{}

func NewService() Service {
	return &svc{}
}

func (s *svc) add(name string) error {
	return nil
}

func (s *svc) remove(id int) error {	
	return nil
}

func (s *svc) getAll() ([]model, error) {
	return []model{}, nil
}

type addRequest struct {
	Name string `json:"name"`
 }
 
 type removeRequest struct {
	ID int `json:"id"`
 }