package request

type RegisterReq struct {
	Email    string `json:"email" example:"example@economicus.kr"`
	Password string `json:"password" example:"some password"`
	Nickname string `json:"nickname" example:"user nickname"`
}
