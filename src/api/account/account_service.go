package account

import (
	"ReservApp/src/api/account/dtos"
	"ReservApp/src/cmd"
	"ReservApp/src/db"
	"ReservApp/src/db/models"
	"errors"
	"fmt"
)

type AccountService struct {
	cfg            *cmd.AppConfig
	psqlRepository *db.PsqlRepository
}

func NewAccountService(cfg *cmd.AppConfig) *AccountService {
	psqlRepository := db.NewPsqlRepository()
	return &AccountService{
		cfg:            cfg,
		psqlRepository: psqlRepository,
	}
}

func (s *AccountService) Create(dto *dtos.CreateAccountDto, user *models.User) (*models.Account, error) {
	fmt.Println("*********************")
	fmt.Println(dto.Currency)
	account := models.Account{Balance: 0, Currency: dto.Currency, Owner: *user.Username}
	createdAccount := s.psqlRepository.DB.Model(&models.Account{}).Create(&account)
	user.Account = account
	s.psqlRepository.DB.Save(user)
	if createdAccount.RowsAffected == 0 {
		return nil, errors.New("register unsuccessfull!")
	}

	return &account, nil
}
func (s *AccountService) FindAll() ([]models.User, error) {
	var users []models.User
	err := s.psqlRepository.DB.Model(&models.User{}).Preload("Account").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *AccountService) Update() {

}
func (s *AccountService) Delete() {

}
