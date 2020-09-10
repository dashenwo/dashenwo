package schema

// 用户信息字段（对外展示）
type Account struct {
	// 昵称
	Nickname string `json:"nickname"`
	// 密码
	Password string `json:"password"`
	// 手机号
	Phone string `json:"phone"`
	// 邮箱
	Email string `json:"email"`
	// 性别
	Sex int `json:"sex"`
	// 生日
	Birthday int32 `json:"birthday"`
	// 签名
	UserSign string `json:"user_sign"`
	// 粉丝数量
	FansNum int32 `json:"fans_num"`
	// 头像地址
	AvatarHd string `json:"avatar_hd"`
	// 关注数量
	FollowNum string `json:"follow_num"`
	// 获得的点赞数量
	LoveNum string `json:"love_num"`
	// 注册时间
	RegisterTime int32 `json:"register_time"`
}
