package logic

import (
	"go-web-app/dao/redis"
	"go-web-app/models"
	"strconv"

	"go.uber.org/zap"
)

// vote
/*
when direction = 1:
	if haven't voted
	if voted downVote
when direction = 0:
	if voted upVote
	if voted DownVote
when direction = -1:
	if haven't voted
	if voted upVote

Limitations:
1. After one week the post posted, make user unable to vote anymore
2. After one week the post posted, delete KeyPostVotedZSetPF
*/
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostId),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostId, float64(p.Direction))
}
