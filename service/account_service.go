package service

import (
	"errors"
	"fmt"
	"golang-tutorial/contract"
	"golang-tutorial/dto"
	"golang-tutorial/entity"
	"golang-tutorial/pkg/token"
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
			Username: account.Username,
			Email:    account.Email,
		},
	}
	return response, nil
}

func (s *AccountService) CreateAccount(payload *dto.AccountRequest) (*dto.AccountResponse, error) {
	if !isValidEmail(payload.Email) {
		return nil, errors.New("email tidak valid")
	}
	if !isValidPassword(payload.Password) {
		return nil, errors.New("password tidak valid (harus ada minimal 1 huruf besar, 1 angka, 1 simbol)")
	}

	emailExists, err := s.AccountRepository.CheckEmail(payload.Email)
	if err != nil {
		return nil, err
	}
	if emailExists {
		return nil, errors.New("email sudah terdaftar")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	account := &entity.Account{
		Email:    payload.Email,
		Username: payload.Username,
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
			Username: account.Username,
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
    fmt.Println("Start Login Function")

    if payload.Username == "" || payload.Password == "" {
        fmt.Println("Missing username or password")
        return nil, errors.New("username dan password harus diisi")
    }

    fmt.Println("Fetching account by username:", payload.Username)
    account, err := s.AccountRepository.GetAccountByUsername(payload.Username)
    if err != nil {
        fmt.Println("Error fetching account:", err)
        return nil, errors.New("not found: user not found")
    }

    fmt.Println("Validating password")
    err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(payload.Password))
    if err != nil {
        fmt.Println("Password mismatch")
        return nil, errors.New("password salah")
    }

    fmt.Println("Generating access token")
    accessToken, err := token.GenerateToken(&token.UserAuthToken{
        ID:       uint64(account.ID),
        Email:    account.Email,
        Username: account.Username,
    })
    if err != nil {
        fmt.Println("Error generating access token:", err)
        return nil, fmt.Errorf("failed to generate access token: %v", err)
    }

	fmt.Println("token=", accessToken) 

    fmt.Println("Generating refresh token")
    refreshToken, err := token.GenerateRefreshToken(uint64(account.ID))
    if err != nil {
        fmt.Println("Error generating refresh token:", err)
        return nil, fmt.Errorf("failed to generate refresh token: %v", err)
    }

    fmt.Println("Creating response")
    response := &dto.AccountResponse{
        StatusCode: http.StatusOK,
        Massage:    "Berhasil Login",
        Data: dto.AccountData{
            ID:        account.ID,
            Username:  account.Username,
            Email:     account.Email,
            CreatedAt: account.CreatedAt,
            UpdatedAt: account.UpdatedAt,
        },
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }

    fmt.Println("Login successful")
    return response, nil
}