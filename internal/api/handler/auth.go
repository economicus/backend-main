package handler

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/service"
	e "main/internal/core/error"
	"net/http"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

// LoginInLocal godoc
// @Summary      Local login
// @Description  이메일, 비밀번호로 로그인하기
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        email     body      string  true  "User login email"
// @Param        password  body      string  true  "User login password"
// @Success      200       {object}  jwt.Token
// @Failure      400       {object}  httpError  "Bad request error"
// @Failure      401       {object}  httpError  "Unauthorized error"
// @Failure      404       {object}  httpError  "Not found error"
// @Failure      500       {object}  httpError  "Internal server error"
// @Router       /login [post]
func (h *AuthHandler) LoginInLocal(ctx *gin.Context) {
	req := struct {
		Email    string `json:"email" example:"example@economicus.kr"`
		Password string `json:"password" example:"some password"`
	}{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}
	if req.Email == "" || req.Password == "" {
		sendErr(ctx, e.ErrMissingRequest)
		return
	}

	token, err := h.service.LoginInLocal(req.Email, req.Password)
	if err != nil {
		sendErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, token)
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"Message": "not set up yet",
	})
}

// RefreshToken godoc
// @Summary      Refresh jwt token
// @Description  Access token 기간 만료시, Refresh token을 사용하여 jwt 토큰 재발급
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        refresh_token  body      string     true  "Refresh token"
// @Success      200            {object}  string     "refreshed access token"
// @Failure      400            {object}  httpError  "Bad request error"
// @Failure      401            {object}  httpError  "Unauthorized error"
// @Failure      404            {object}  httpError  "Not found error"
// @Failure      500            {object}  httpError  "Internal server error"
// @Router       /refresh-token [post]
func (h *AuthHandler) RefreshToken(ctx *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" example:"some refresh token"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	token, err := h.service.RefreshToken(req.RefreshToken)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}
