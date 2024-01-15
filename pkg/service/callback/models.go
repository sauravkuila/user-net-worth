package callback

type GetIdirectSessionKeyRequest struct {
	ApiSession int64 `form:"apisession" binding:"required"`
}
