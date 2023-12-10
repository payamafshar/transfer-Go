package transfer

import (
	"ReservApp/src/api/transfer/dtos"
	"ReservApp/src/cmd"
	"net/http"

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
	dto := new(dtos.CreateTransferDto)
	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transfer, err := t.TransferService.Create(dto)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, transfer)
}
