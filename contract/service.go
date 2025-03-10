package contract

import "golang-tutorial/dto"

type Service struct {
	Intro IntroService
}

// type exampleService interface {
// Code here
// }

type IntroService interface {
	GetIntro(introID int) (*dto.IntroResponse, error)
	CreateIntro(payload *dto.IntroRequest) (*dto.IntroResponse, error)
}
