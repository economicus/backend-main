package handler

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/service"
	e "main/internal/core/error"
	"main/internal/core/model"
	"main/internal/core/model/request"
	"main/internal/core/model/response"
	"net/http"
)

type UserHandler struct {
	userService  *service.UserService
	quantService *service.QuantService
}

func NewUserHandler(s *service.UserService, q *service.QuantService) *UserHandler {
	return &UserHandler{
		userService:  s,
		quantService: q,
	}
}

// Register godoc
// @Summary      Register a user
// @Description  이메일, 비밀번호로 유저 회원가입
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body  request.RegisterReq  true  "A user information"
// @Success      201
// @Failure      400            {object}  httpError                 "Bad request error"
// @Failure      401            {object}  httpError                 "Unauthorized error"
// @Failure      404            {object}  httpError                 "Not found error"
// @Failure      500            {object}  httpError                 "Internal server error"
// @Router       /register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	var req request.RegisterReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	if err := h.userService.Register(&req); err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

// GetAllUsers godoc
// @Summary      Return all users
// @Description  이메일, 비밀번호로 유저 회원가입
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string                    true  "Bearer {access_token}"
// @Param        body           body    request.RegisterReq  true  "A user information"
// @Success      201
// @Failure      400  {object}  httpError  "Bad request error"
// @Failure      401  {object}  httpError  "Unauthorized error"
// @Failure      404  {object}  httpError  "Not found error"
// @Failure      500  {object}  httpError  "Internal server error"
// @Router       /users [get]
func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	option := model.NewQuery()

	if err := ctx.BindQuery(option); err != nil {
		sendErr(ctx, err)
		return
	}

	users, err := h.userService.GetUsers(option)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetUser returns a user with id
func (h *UserHandler) GetUser(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	userID, err := extractUserId(ctx)
	if err != nil {
		userID = user.ID
	}

	fields := getFieldsFromContext(ctx)

	resp, err := h.userService.GetUser(userID, fields)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// GetUserProfile godoc
// @Summary      Return user's profile and quants
// @Description  프로필 화면에서 유저의 정보 및 보유한 퀀트 모델을 반환
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string               true  "Bearer {access_token}"
// @Param        user_id        path      uint                      true  "User id to get profile"
// @Success      200            {object}  response.ProfileResponse  "Profile response"
// @Failure      400  {object}  httpError  "Bad request error"
// @Failure      401  {object}  httpError  "Unauthorized error"
// @Failure      404  {object}  httpError  "Not found error"
// @Failure      500  {object}  httpError  "Internal server error"
// @Router       /users/profile/{user_id}  [get]
func (h *UserHandler) GetUserProfile(ctx *gin.Context) {
	var uri struct {
		UserID uint `uri:"user_id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		sendInvalidPathErr(ctx, err)
		return
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	if uri.UserID != user.ID {
		sendErr(ctx, e.ErrPermissionDenied)
		return
	}

	quant, err := h.quantService.GetUsersQuant(user.ID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	res := response.ProfileResponse{
		User:  *user,
		Quant: quant,
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteUser deletes a user
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = h.userService.DeleteUser(user.ID)
	if err != nil {

		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// EditUserProfile edit a user profile
func (h *UserHandler) EditUserProfile(ctx *gin.Context) {
	var req model.Profile

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.userService.UpdateProfile(user.ID, &req)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// UploadUserProfileImage edit user's profile image
func (h *UserHandler) UploadUserProfileImage(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	file, header, err := ctx.Request.FormFile("profile_image")
	if err != nil {
		sendErr(ctx, e.ErrInvalidFile)
		return
	}
	defer file.Close()

	err = h.userService.UploadProfileImage(user.ID, file, header)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// GetFollowings returns list of followings
func (h *UserHandler) GetFollowings(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	followings, err := h.userService.GetFollowings(user.ID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, followings)
}

// GetFollowers returns list of followers
func (h *UserHandler) GetFollowers(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	followings, err := h.userService.GetFollowers(user.ID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, followings)
}

// FollowUser refreshes access token
func (h *UserHandler) FollowUser(ctx *gin.Context) {
	var data struct {
		FollowerID uint `json:"follower_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.userService.Follow(user.ID, data.FollowerID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// UnfollowUser refreshes access token
func (h *UserHandler) UnfollowUser(ctx *gin.Context) {
	var data struct {
		FollowingID uint `json:"following_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.userService.UnFollow(user.ID, data.FollowingID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

// GetFavoriteQuants returns a favorite quant list of user
func (h *UserHandler) GetFavoriteQuants(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	quants, err := h.userService.GetFavoriteQuants(user.ID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count":  len(quants),
		"quants": quants,
	})
}

// AddToFavoriteQuants add a quant to favorite list
func (h *UserHandler) AddToFavoriteQuants(ctx *gin.Context) {
	var data struct {
		QuantID uint `json:"quant_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.userService.AddToFavoriteQuants(user.ID, data.QuantID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

// DeleteFromFavoriteQuants add a quant to favorite list
func (h *UserHandler) DeleteFromFavoriteQuants(ctx *gin.Context) {
	var data struct {
		QuantID uint `json:"quant_id"`
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	err = h.userService.DeleteFromFavoriteQuants(user.ID, data.QuantID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
