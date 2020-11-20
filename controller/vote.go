package controller

import (
	"go-web-app/logic"
	"go-web-app/models"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type ParamVoteData struct {
	PostId    int64 `json:"post_id,string"`
	Direction int   `json:"direction,string"` // agree or disagree
}

func PostVoteHandler(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBind(p); err != nil {
		validationError, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(validationError.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	logic.PostVote()
	ResponseSuccess(c, nil)
}
