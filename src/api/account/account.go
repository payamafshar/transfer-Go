package account

type GetAccountByIdRequest struct {
	Id int32 `uri:"id" binding:"required,min=1"`
}
type GetListAccountRequest struct {
	PageSize int `form:"pageSize" binding:"min=5,max=10"`
	PageId   int `form:"pageId" binding:"min=1"`
}
