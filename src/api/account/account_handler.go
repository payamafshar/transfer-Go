package account

import (
	"ReservApp/src/api/account/dtos"
	"ReservApp/src/cmd"
	"ReservApp/src/db/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	cfg            *cmd.AppConfig
	AccountService *AccountService
}

func NewAccountHandler() *AccountHandler {
	cfg := cmd.GetAppConfig()
	accountService := NewAccountService(cfg)
	return &AccountHandler{
		cfg:            cfg,
		AccountService: accountService,
	}
}

func (h *AccountHandler) Create(ctx *gin.Context) {
	user := ctx.MustGet("user")
	castedUser, ok := user.(*models.User)
	if !ok {
		panic("invalid user object in context")
	}
	fmt.Println(user)
	dto := new(dtos.CreateAccountDto)
	if err := ctx.ShouldBindJSON(dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdAccount, err := h.AccountService.Create(dto, castedUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, createdAccount)
}

func (h *AccountHandler) GetById(ctx *gin.Context) {

	var req GetAccountByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account, err := h.AccountService.GetById(req.Id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, account)
	return
}

func (h *AccountHandler) FindAll(ctx *gin.Context) {
	var req GetListAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accounts, err := h.AccountService.FindAll(int(req.PageSize), int(req.PageId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, accounts)
}
func (h *AccountHandler) Update(ctx *gin.Context) {

}
func (h *AccountHandler) Delete(ctx *gin.Context) {

}
