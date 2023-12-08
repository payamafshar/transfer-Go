package account

import (
	"ReservApp/src/api/account/dtos"
	"ReservApp/src/cmd"
	"ReservApp/src/db"
	"ReservApp/src/db/models"
	"errors"
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
	var findedAccount models.Account
	if existAccount := s.psqlRepository.DB.Where("owner = ?", *user.Username).First(&findedAccount); existAccount.Error == nil {
		return nil, errors.New("Cannot create more than 1 account")
	}
	account := models.Account{Balance: 0, Currency: dto.Currency, Owner: *user.Username}
	createdAccount := s.psqlRepository.DB.Model(&models.Account{}).Create(&account)
	user.Account = account
	s.psqlRepository.DB.Save(user)
	if createdAccount.RowsAffected == 0 {
		return nil, errors.New("create account unsuccessfull!")
	}

	return &account, nil
}
func (s *AccountService) GetById(Id int32) (*models.Account, error) {

	var account models.Account
	if findedAccount := s.psqlRepository.DB.Where("id = ?", Id).First(&account); findedAccount.Error != nil {
		return nil, errors.New("Account Not Found")
	}
	return &account, nil
}
func (s *AccountService) FindAll(pageSize int, page int) (*[]models.Account, error) {
	offset := (page - 1) * pageSize
	var accounts []models.Account
	err := s.psqlRepository.DB.Model(&models.Account{}).Offset(offset).Limit(pageSize).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return &accounts, nil
}
func (s *AccountService) Update() {

}
func (s *AccountService) Delete() {

}
