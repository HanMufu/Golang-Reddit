package models

// use to define signup parameters
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// use to define login parameters
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostId    int64 `json:"post_id,string" binding:"required"`
	Direction int8  `json:"direction,string" binding:"required,oneof=-1 0 1"` // agree or disagree or neither disagree or agree
}
