package transfer

import (
	"ReservApp/src/cmd"

	"github.com/gin-gonic/gin"
)

type TransferHandler struct {
	cfg             *cmd.AppConfig
	TransferService *TransferService
}

func NewTransferHandler() *TransferHandler {
	cfg := cmd.GetAppConfig()
	transferService := NewTransferService(cfg)
	return &TransferHandler{
		cfg:             cfg,
		TransferService: transferService,
	}
}

func (t *TransferHandler) Create(ctx *gin.Context) {

}
