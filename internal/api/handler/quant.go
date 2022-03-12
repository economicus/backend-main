package handler

import (
	"github.com/gin-gonic/gin"
	"main/internal/api/service"
	"main/internal/core/data"
	"main/internal/core/model"
	"main/internal/core/model/request"
	"net/http"
)

type QuantHandler struct {
	service *service.QuantService
}

func NewQuantHandler(s *service.QuantService) *QuantHandler {
	return &QuantHandler{
		service: s,
	}
}

// GetAllQuants godoc
// @Summary      Return a list of quant models
// @Description  메인 화면에서 사용될 퀀트 모델들 반환, 개선 필요
// @Tags         quant
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string        true   "Bearer {access_token}"
// @Param        page           query     int           false  "number of page"      default(10)
// @Param        per_page       query     int           false  "number of elements"  default(10)
// @Param        order          query     string        false  "fields for order"    default("")
// @Param        keyword        query     string        false  "keyword for query"   default("")
// @Success      200            {object}  model.Quants  "List of quants"
// @Failure      400            {object}  httpError     "Bad request error"
// @Failure      401            {object}  httpError     "Unauthorized error"
// @Failure      404            {object}  httpError     "Not found error"
// @Failure      500            {object}  httpError     "Internal server error"
// @Router       /quants [get]
func (h *QuantHandler) GetAllQuants(ctx *gin.Context) {
	option := model.NewQuery()

	if err := ctx.BindQuery(option); err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	quants, err := h.service.GetAllQuants(user.ID, option)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, quants)
}

// GetQuant godoc
// @Summary      Return a quant model
// @Description  모델 상세페이지에서 사용될 퀀트 모델 반환, 개선 필요
// @Tags         quant
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string       true  "Bearer {access_token}"
// @Param        quant_id       path      uint         true  "ID of a quant"
// @Success      200            {object}  model.Quant  "A quant"
// @Failure      400            {object}  httpError    "Bad request error"
// @Failure      401            {object}  httpError    "Unauthorized error"
// @Failure      404            {object}  httpError    "Not found error"
// @Failure      500            {object}  httpError    "Internal server error"
// @Router       /quants/quant/{quant_id} [get]
func (h *QuantHandler) GetQuant(ctx *gin.Context) {
	var uri struct {
		QuantID uint `uri:"quant_id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&uri); err != nil {
		sendInvalidPathErr(ctx, err)
		return
	}

	quant, err := h.service.GetQuant(uri.QuantID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, quant)
}

// GetLabList godoc
// @Summary      Return a list of user's quants
// @Description  실험실에 들어갈 때, 유저가 실험실에서 테스트하던 모델들 반환
// @Tags         quant
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string               true  "Bearer {access_token}"
// @Success      200            {object}  []response.LabQuant  "Quant list"
// @Failure      400            {object}  httpError            "Bad request error"
// @Failure      401            {object}  httpError            "Unauthorized error"
// @Failure      404            {object}  httpError            "Not found error"
// @Failure      500            {object}  httpError            "Internal server error"
// @Router       /lab/list [get]
func (h *QuantHandler) GetLabList(ctx *gin.Context) {
	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	resp, err := h.service.GetLabList(user.ID)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"quants": resp,
		"kospi":  data.FixedKD,
	})
}

// GetLabData godoc
// @Summary      Return chart and option data
// @Description  실험실에서 모델을 체크할 때, 모델의 차트 데이터와 옵션 정보를 반환
// @Tags         quant
// @Accept       json
// @Produce      json
// @Param        Authorization  header    string            true  "Bearer {access_token}"
// @Param        quant_id       path      uint              true  "QuantID to get data"
// @Success      200            {object}  response.LabData  "Data of a quant"
// @Failure      400            {object}  httpError         "Bad request error"
// @Failure      401            {object}  httpError         "Unauthorized error"
// @Failure      404            {object}  httpError         "Not found error"
// @Failure      500            {object}  httpError         "Internal server error"
// @Router       /lab/data/{quant_id} [get]
func (h *QuantHandler) GetLabData(ctx *gin.Context) {
	var uri struct {
		QuantID uint `uri:"quant_id" binding:"required"`
	}

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		sendInvalidPathErr(ctx, err)
		return
	}

	resp, err := h.service.GetLabData(uri.QuantID)
	if err != nil {
		sendInternalErr(ctx)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

// CreateQuant godoc
// @Summary      Create a quant model
// @Description  실험실에서 모델 만들기를 눌렀을 때, 모델 생성
// @Tags         quant
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string             true  "Bearer {access_token}"
// @Param        quant          body    model.QuantOption  true  "Quant option data"  "desc"
// @Success      201
// @Failure      400  {object}  httpError  "Bad request error"
// @Failure      401  {object}  httpError  "Unauthorized error"
// @Failure      404  {object}  httpError  "Not found error"
// @Failure      500  {object}  httpError  "Internal server error"
// @Router       /quants/quant [post]
func (h *QuantHandler) CreateQuant(ctx *gin.Context) {
	var req model.QuantOption

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

	res, err := h.service.CreateQuant(user.ID, &req)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// UpdateQuant godoc
// @Summary      Update a quant model
// @Description  모델 저장버튼(activate)을 누르거나, 모델 설명을 변경하고자 할 경우 사용
// @Tags         quant
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string                true  "Bearer {access_token}"
// @Param        quant_id       path      uint         true  "ID of a quant"
// @Param        body           body    request.EditQuantReq  true  "Quant data"
// @Success      204
// @Failure      400  {object}  httpError  "Bad request error"
// @Failure      401  {object}  httpError  "Unauthorized error"
// @Failure      403  {object}  httpError  "Forbidden error"
// @Failure      404  {object}  httpError  "Not found error"
// @Failure      500  {object}  httpError  "Internal server error"
// @Router       /quants/quant/{quant_id} [patch]
func (h *QuantHandler) UpdateQuant(ctx *gin.Context) {
	var uri struct {
		QuantID uint `uri:"quant_id" binding:"required"`
	}
	var req request.EditQuantReq

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		sendInvalidPathErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = h.service.UpdateQuant(user.ID, uri.QuantID, &req)
	if err != nil {
		sendErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// UpdateQuantOption godoc
// @Summary      Update a quant option
// @Description  퀀트 옵션을 변경하고자 할 경우 사용
// @Tags         quant
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string                      true  "Bearer {access_token}"
// @Param        body           body    request.EditQuantOptionReq  true  "Quant option data"
// @Param        quant_id       path    uint                        true  "QuantID to update"
// @Success      204
// @Failure      400  {object}  httpError  "Bad request error"
// @Failure      401  {object}  httpError  "Unauthorized error"
// @Failure      403  {object}  httpError  "Forbidden error"
// @Failure      404  {object}  httpError  "Not found error"
// @Failure      500  {object}  httpError  "Internal server error"
// @Router       /quants/quant-option/{quant_id} [patch]
func (h *QuantHandler) UpdateQuantOption(ctx *gin.Context) {
	var uri struct {
		QuantID uint `uri:"quant_id" binding:"required" example:"1"`
	}
	var req request.EditQuantOptionReq

	err := ctx.ShouldBindUri(&uri)
	if err != nil {
		sendInvalidPathErr(ctx, err)
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		sendJsonParsingErr(ctx, err)
		return
	}

	user, err := getUserFromContext(ctx)
	if err != nil {
		sendErr(ctx, err)
		return
	}

	err = h.service.UpdateQuantOption(user.ID, uri.QuantID, &req)
	if err != nil {
		sendErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// DeleteQuant godoc
// @Summary      Delete a quant model
// @Description  퀀트 모델을 제거할 경우 사용
// @Tags         quant
// @Accept       json
// @Produce      json
// @Param        Authorization  header  string  true  "Bearer {access_token}"
// @Param        quant_id       path    uint    true  "Quant ID to delete"
// @Success      204
// @Failure      400  {object}  httpError  "Bad request error"
// @Failure      401  {object}  httpError  "Unauthorized error"
// @Failure      403  {object}  httpError  "Forbidden error"
// @Failure      404  {object}  httpError  "Not found error"
// @Failure      500  {object}  httpError  "Internal server error"
// @Router       /quants/quant/{quant_id} [delete]
func (h *QuantHandler) DeleteQuant(ctx *gin.Context) {
	var uri struct {
		QuantID uint `uri:"quant_id" binding:"required" example:"3"`
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

	err = h.service.DeleteQuant(user.ID, uri.QuantID)
	if err != nil {
		sendErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
