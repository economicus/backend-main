package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/internal/api/repo"
	g "main/internal/conf/grpc"
	"main/internal/core/model"
	"main/internal/core/model/request"
	"main/internal/core/model/response"
	"main/internal/pkg/objconv"
)

type QuantService struct {
	repo *repo.QuantRepo
	grpc *g.Quant
}

func NewQuantService(repo *repo.QuantRepo) *QuantService {
	return &QuantService{
		repo: repo,
		grpc: g.New(),
	}
}

func (s *QuantService) GetAllQuants(userID uint, option *model.Query) (model.Quants, error) {
	return s.repo.GetAllQuants(userID, option)
}

func (s *QuantService) GetQuant(quantID uint) (*model.Quant, error) {
	return s.repo.GetQuant(quantID)
}

func (s *QuantService) GetUsersQuant(userID uint) ([]response.ProfileQuantResponse, error) {
	ids, err := s.repo.GetUserQuantIds(userID)
	if err != nil {
		return nil, err
	}
	return s.repo.GetChartByIds(ids)
}

func (s *QuantService) GetMyQuants(userID uint) (*model.Quants, error) {
	return s.repo.GetMyQuants(userID)
}

func (s *QuantService) GetLabList(userID uint) (model.Quants, error) {
	return s.repo.GetLabList(userID)
}

func (s *QuantService) GetLabData(quantID uint) (*response.LabData, error) {
	q, err := s.repo.GetQuant(quantID)
	if err != nil {
		return nil, err
	}
	chart, err := s.repo.GetChart(q.ChartID)
	if err != nil {
		return nil, err
	}
	option := model.NewQuantOption(&q.QuantOption)
	res := response.LabData{
		Option: *option,
		Chart:  *chart,
	}
	return &res, nil
}

func (s *QuantService) CreateQuant(userID uint, req *model.QuantOption) (*response.QuantResponse, error) {
	if err := s.repo.CheckModelName(req.Name); err != nil {
		return nil, err
	}

	quant := model.NewQuant(userID, req.Name)
	quantID, err := s.repo.CreateQuant(quant)
	if err != nil {
		return nil, err
	}

	req.QuantID = quantID

	if err = s.repo.CreateQuantOption(req.ToTable()); err != nil {
		return nil, err
	}

	resp, err := s.getQuantResponse(req)
	if err != nil {
		return nil, err
	}

	resp.QuantID = quantID
	resId, err := s.repo.CreateQuantResult(resp.QuantID, resp.ChartData.ProfitRateData)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	m["chart_id"] = resId.(primitive.ObjectID).Hex()
	m["cumulative_return"] = resp.CumulativeReturn
	m["annual_average_return"] = resp.AnnualAverageReturn
	m["winning_percentage"] = resp.WinningPercentage
	m["max_loss_rate"] = resp.MaxLossRate
	m["holdings_count"] = resp.HoldingsCount

	if err = s.repo.UpdateQuant(quantID, m); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *QuantService) getQuantResponse(req *model.QuantOption) (*response.QuantResponse, error) {
	gReq := req.ToRequest()
	result, err := s.grpc.Request(gReq)
	if err != nil {
		return nil, err
	}

	resp := response.NewQuantResultFromPB(result)
	return resp, nil
}

func (s *QuantService) UpdateQuant(userID, quantID uint, req *request.EditQuantReq) error {
	q, err := s.GetQuant(quantID)
	if err != nil {
		return err
	}

	if err = repo.CheckPermission(userID, q); err != nil {
		return err
	}

	reqBody := objconv.ToMap(req)

	return s.repo.UpdateQuant(q.ID, reqBody)
}

func (s *QuantService) UpdateQuantOption(userID, quantID uint, req *request.EditQuantOptionReq) error {
	q, err := s.GetQuant(quantID)
	if err != nil {
		return err
	}

	if err = repo.CheckPermission(userID, q); err != nil {
		return err
	}

	return s.repo.UpdateQuantOption(q.ID, req.ToMap())
}

func (s *QuantService) DeleteQuant(userID, quantID uint) error {
	q, err := s.GetQuant(quantID)
	if err != nil {
		return err
	}

	if err = repo.CheckPermission(userID, q); err != nil {
		return err
	}

	return s.repo.DeleteQuant(quantID)
}
