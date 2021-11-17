package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// 存放请求参数的模型

// ParamSignUp 用户注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

// ParamLogin 用户登录参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票参数
type ParamVoteData struct {
	//UserId从请求中获取
	UserId    int64
	PostId    int64 `json:"post_id,string" binding:"required"`        // 帖子id
	Direction int8  `json:"direction,string" binding:"oneof= -1 0 1"` // 赞成还是反对（+1、 -1），0为取消投票
}

// ParamPostList 获取帖子列表的QueryString参数
type ParamPostList struct {
	CommunityId int64  `json:"community_id" form:"community_id"`
	Offset      int64  `json:"offset" form:"offset"`
	Limit       int64  `json:"limit" form:"limit"`
	Order       string `json:"order" form:"order"`
}

//type ParamCommunityPostList struct {
//	CommunityId int64 `json:"community_id" form:"community_id"`
//	*ParamPostList
//}
