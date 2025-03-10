package service

import (
	"golang-tutorial/contract"
	"golang-tutorial/dto"
	"golang-tutorial/entity"
	"net/http"
)

type IntroService struct {
	introRepository contract.IntroRepository
}

func implIntroService(repo *contract.Repository) contract.IntroService {
	return &IntroService{
		introRepository: repo.Intro,
	}
}

func (s *IntroService) GetIntro(introID int) (*dto.IntroResponse, error) {
	intro, err := s.introRepository.GetIntro(introID)
	if err != nil {
		return nil, err
	}

	response := &dto.IntroResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan data",
		Data: dto.IntroData{
			ID:            intro.ID,
			Nama:          intro.Nama,
			NamaPanggilan: intro.NamaPanggilan,
			FunFact:       intro.FunFact,
			KeinginanBE:   intro.KeinginanBE,
		},
	}
	return response, nil
}

func (s *IntroService) CreateIntro(payload *dto.IntroRequest) (*dto.IntroResponse, error) {
	intro := &entity.Intro{
		Nama:          payload.Nama,
		NamaPanggilan: payload.NamaPanggilan,
		FunFact:       payload.FunFact,
		KeinginanBE:   payload.KeinginanBE,
	}

	err := s.introRepository.CreateIntro(intro)
	if err != nil {
		return nil, err
	}

	response := &dto.IntroResponse{
		StatusCode: http.StatusCreated,
		Message:    "Berhasil membuat data",
		Data: dto.IntroData{
			ID:            intro.ID,
			Nama:          intro.Nama,
			NamaPanggilan: intro.NamaPanggilan,
			FunFact:       intro.FunFact,
			KeinginanBE:   intro.KeinginanBE,
		},
	}

	return response, nil
}
