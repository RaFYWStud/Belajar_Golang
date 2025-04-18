package service

import (
	"errors"
	"golang-tutorial/contract"
	"golang-tutorial/dto"
	"golang-tutorial/entity"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	AccountRepository contract.AccountRepository
}

func implAccountService(repo *contract.Repository) *AccountService {
	return &AccountService{
		AccountRepository: repo.Account,
	}
}

func (s *AccountService) GetAccount(accountID int) (*dto.AccountResponse, error) {
	account, err := s.AccountRepository.GetAccount(accountID)
	if err != nil {
		return nil, err
	}

	response := &dto.AccountResponse{
		StatusCode : http.StatusOK,
		Massage : "Berhasil mendapatkan data",
		Data : dto.AccountData{
			ID:       account.ID,
			Email:    account.Email,
		},
	}
	return response, nil
}

func (s *AccountService) CreateAccount(payload *dto.AccountRequest) (*dto.AccountResponse, error) {
	if !isValidEmail(payload.Email) {
		return nil, errors.New("Email tidak valid (Gunakan format email @unity.com)")
	}
	if !isValidPassword(payload.Password) {
		return nil, errors.New("Password tidak valid (Harus ada minimal 1 huruf besar, 1 angka, 1 simbol)")
	}

	emailExists, err := s.AccountRepository.CheckEmail(payload.Email)
	if err != nil {
		return nil, err
	}
	if emailExists {
		return nil, errors.New("Email sudah terdaftar")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	account := &entity.Account{
		Email:    payload.Email,
		Password: string(hashPassword),
	}

    err = s.AccountRepository.CreateAccount(account)
	if err != nil {
		return nil, err
	}

	response := &dto.AccountResponse{
		StatusCode: http.StatusCreated,
		Massage:	"Berhasil register akun",
		Data: dto.AccountData{
			ID:     account.ID,
			Email:  account.Email,
		},
	}

	return response, nil
}

func isValidEmail (email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@unity\.com$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),._?":{}|<>]`).MatchString(password)
	allowedChar := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*(),._?":{}|<>]+$`).MatchString(password)

	return hasUpper && hasDigit && hasSpecial && allowedChar
}

func (s *AccountService) Login(payload *dto.AccountRequest) (*dto.AccountResponse, error) {
	account, err := s.AccountRepository.GetAccountByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(payload.Password))
	if err != nil {
		return nil, errors.New("Password salah")
	}

	response := &dto.AccountResponse{
		StatusCode: http.StatusOK,
		Massage :   "Berhasil Login",
		Data: dto.AccountData{
			ID:       account.ID,
			Email:    account.Email,
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		},
	}

	return response, nil
}