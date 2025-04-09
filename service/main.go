package service

import "golang-tutorial/contract"

func New(repo *contract.Repository) *contract.Service {
	return &contract.Service{
		// Code here
		// Example:
		// Example: implExampleService(repo),
		Intro: implIntroService(repo),
		ToDo: implToDoService(repo),
	}
}
