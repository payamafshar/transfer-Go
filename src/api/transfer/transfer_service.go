package transfer

import (
	"ReservApp/src/cmd"
	"ReservApp/src/db"
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

func (t *TransferService) Create()
