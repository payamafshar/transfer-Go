package transfer

import (
	"ReservApp/src/api/transfer/dtos"
	"ReservApp/src/cmd"
	"ReservApp/src/db"
	"ReservApp/src/db/models"
	"errors"
)

type TransferService struct {
	cfg            *cmd.AppConfig
	psqlRepository *db.PsqlRepository
}

func NewTransferService(cfg *cmd.AppConfig) *TransferService {
	psqlRepository := db.NewPsqlRepository()
	return &TransferService{
		cfg:            cfg,
		psqlRepository: psqlRepository,
	}
}

func (t *TransferService) Create(dto *dtos.CreateTransferDto) (*models.Transfer, error) {

	var fromAccount models.Account
	var toAccount models.Account
	if findedFromAccount := t.psqlRepository.DB.Where("id = ?", dto.FromAccountId).First(&fromAccount); findedFromAccount.Error != nil {
		return nil, errors.New("FromAccount Not Found")
	}
	if findedToAccount := t.psqlRepository.DB.Where("id = ?", dto.ToAccountId).First(&toAccount); findedToAccount.Error != nil {
		return nil, errors.New("FromAccount Not Found")
	}
	if fromAccount.Currency != toAccount.Currency {
		return nil, errors.New("account currencys do not match")
	}
	transfer := models.Transfer{FromAccountID: fromAccount.Id, FromAccount: fromAccount, ToAccountID: toAccount.Id, ToAccount: toAccount, Amount: dto.Amount}
	createdDbTransfer := t.psqlRepository.DB.Model(&models.Transfer{}).Create(&transfer)

	if createdDbTransfer.RowsAffected < 1 {
		return nil, errors.New("Transfer not created")
	}

	entryForSenderAccount := models.Entrie{AccountId: fromAccount.Id, Amount: -transfer.Amount, TransferID: transfer.Id}
	createdEntryForSender := t.psqlRepository.DB.Model(&models.Entrie{}).Create(&entryForSenderAccount)
	if createdEntryForSender.Error != nil {
		return nil, createdEntryForSender.Error
	}
	entryForReciverAccount := models.Entrie{AccountId: toAccount.Id, Amount: transfer.Amount, TransferID: transfer.Id}
	createdEntryForReciver := t.psqlRepository.DB.Model(&models.Entrie{}).Create(&entryForReciverAccount)
	if createdEntryForReciver.Error != nil {
		return nil, createdEntryForReciver.Error
	}

	transfer.Entrie = append(transfer.Entrie, entryForSenderAccount, entryForReciverAccount)

	t.psqlRepository.DB.Save(&transfer)
	return &transfer, nil
}
