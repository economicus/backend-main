package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	e "main/internal/core/error"
	"main/internal/core/model"
	"main/internal/core/model/response"
	"main/internal/core/model/table"
	"main/internal/pkg/logger"
	"time"
)

type QuantRepo struct {
	mysqlDB *gorm.DB
	mongoDB *mongo.Database
}

func NewQuantRepo(mysqlDB *gorm.DB, mongoDB *mongo.Database) *QuantRepo {
	return &QuantRepo{
		mysqlDB: mysqlDB,
		mongoDB: mongoDB,
	}
}

func (repo *QuantRepo) GetQuantData(dataID string) (*response.QuantResponse, error) {
	var resp *response.QuantResponse

	hexId, err := primitive.ObjectIDFromHex(dataID)
	if err != nil {
		logger.Logger.Errorf("error in GetQuantData while getting object id from hex")
		return nil, err
	}
	if err = repo.mongoDB.Collection("chart").FindOne(context.TODO(), hexId).Decode(resp); err != nil {
		logger.Logger.Errorf("error in GetQuantData while getting data from db")
		return nil, err
	}
	return resp, nil
}

// GetAllQuants returns all uploaded quants
func (repo *QuantRepo) GetAllQuants(userID uint, option *model.Query) (model.Quants, error) {
	var quants model.Quants

	sql := fmt.Sprintf("select * from quants as q "+
		"join profiles up on q.user_id = up.user_id "+
		"where q.name like '%%%s%%' or up.nickname like '%%%s%%' "+
		"order by q.user_id in (select following_id from followings where followings.user_id = %d) desc, %s "+
		"limit %d offset %d;",
		option.Keyword, option.Keyword, userID, option.Order, option.PerPage, option.Page*option.PerPage)

	if err := repo.mysqlDB.Raw(sql).Find(&quants).Error; err != nil {
		logger.Logger.Errorf("error in GetAllQuants: %v\n", err)
		return nil, err
	}
	return quants, nil
}

func (repo *QuantRepo) GetLabList(userID uint) (model.Quants, error) {
	var quants model.Quants

	if err := repo.mysqlDB.Model(&model.Quant{}).Where("user_id = ? AND active = 0", userID).Find(&quants).Error; err != nil {
		logger.Logger.Errorf("error in GetLabList: %v\n", err)
		return nil, err
	}
	return quants, nil
}

func (repo *QuantRepo) GetChart(chartID string) (*response.ChartData, error) {
	var res response.ChartData

	objID, err := primitive.ObjectIDFromHex(chartID)
	if err != nil {
		logger.Logger.Errorf("error in GetChart while getting object id from chart id: %v\n", err)
		return nil, err
	}
	if err = repo.mongoDB.Collection("chart").FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&res); err != nil {
		logger.Logger.Errorf("error in GetChart while getting chart data from db: %v\n", err)
		return nil, err
	}
	return &res, nil
}

func (repo *QuantRepo) GetChartByIds(chartIDs []string) ([]response.ProfileQuantResponse, error) {
	var res []response.ProfileQuantResponse
	var ids []primitive.ObjectID

	for _, v := range chartIDs {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			logger.Logger.Errorf("error in GetChartByIds while getting object id")
			return nil, err
		}
		ids = append(ids, id)
	}
	c, err := repo.mongoDB.Collection("chart").Find(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		logger.Logger.Errorf("error in GetChartByIds while getting object from mongo db")
		return nil, err
	}
	if err = c.All(context.TODO(), &res); err != nil {
		logger.Logger.Errorf("error in GetChartByIds while getting object")
		return nil, err
	}
	return res, nil
}

func (repo *QuantRepo) GetOption(quantID uint) (*model.QuantOption, error) {
	var res model.QuantOption

	if err := repo.mysqlDB.Where("quant_id = ?", quantID).First(&res).Error; err != nil {
		logger.Logger.Errorf("error in GetOption while getting option from db: %v\n", err)
		return nil, err
	}
	return &res, nil
}

func (repo *QuantRepo) GetUserQuantIds(userID uint) ([]string, error) {
	var ids []string

	if err := repo.mysqlDB.Raw("select chart_id from quants where user_id = ?", userID).Find(&ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}

// GetMyQuants returns quants of the user
func (repo *QuantRepo) GetMyQuants(userID uint) (*model.Quants, error) {
	var quants model.Quants

	if err := repo.mysqlDB.Model(&model.Quant{}).Where("user_id = ? AND active = 1", userID).Find(&quants).Error; err != nil {
		logger.Logger.Errorf("error in GetMyQuants: %v\n", err)
		return nil, err
	}
	return &quants, nil
}

// GetQuant returns a quant of quant id
func (repo *QuantRepo) GetQuant(quantID uint) (*model.Quant, error) {
	var quant model.Quant

	if err := repo.mysqlDB.Preload("QuantOption").First(&quant, quantID).Error; err != nil {
		logger.Logger.Errorf("error in GetQuant: %v\n", err)
		return nil, err
	}
	return &quant, nil
}

func (repo *QuantRepo) CheckModelName(name string) error {
	err := repo.mysqlDB.Where("name = ?", name).First(&model.Quant{}).Error
	if err == nil {
		return e.ErrDuplicateModelName
	} else if err == gorm.ErrRecordNotFound {
		return nil
	} else {
		logger.Logger.Errorf("error in CheckModelName: %v\n", err)
		return err
	}
}

// CreateQuant creates a quant
func (repo *QuantRepo) CreateQuant(quant *model.Quant) (uint, error) {
	q := quant.Quant
	if err := repo.mysqlDB.Create(&q).Error; err != nil {
		logger.Logger.Errorf("error in CreateQuant: %v\n", err)
		return 0, err
	}
	quant.ID = q.ID
	return q.ID, nil
}

func (repo *QuantRepo) CreateQuantOption(quantOption *table.QuantOption) error {
	if err := repo.mysqlDB.Create(quantOption).Error; err != nil {
		logger.Logger.Errorf("error in CreateQuantOption: %v\n", err)
		return err
	}
	return nil
}

func (repo *QuantRepo) CreateQuantResult(quantID uint, chart []float32) (interface{}, error) {
	res, err := repo.mongoDB.Collection("chart").InsertOne(context.TODO(), bson.D{{"quant_id", quantID}, {"chart", chart}})
	if err != nil {
		logger.Logger.Errorf("error in CreateQuantResult: %v", err)
		return nil, err
	}
	return res.InsertedID, nil
}

func (repo *QuantRepo) UpdateQuant(quantID uint, req map[string]interface{}) error {
	req["updated_at"] = time.Now()
	if err := repo.mysqlDB.First(&model.Quant{}, quantID).Updates(req).Error; err != nil {
		logger.Logger.Errorf("error in UpdateQuant: %v\n", err)
		return err
	}
	return nil
}

func (repo *QuantRepo) UpdateQuantOption(quantID uint, req map[string]interface{}) error {
	req["updated_at"] = time.Now()
	err := repo.mysqlDB.Where("quant_id = ?", quantID).First(&model.QuantOption{}).Updates(req).Error
	if err != nil {
		logger.Logger.Errorf("error in UpdateQuantOption while finding quant: %v\n", err)
		return err
	}
	return nil
}

func (repo *QuantRepo) DeleteQuant(quantID uint) error {
	if err := repo.mysqlDB.Delete(&model.Quant{}, quantID).Error; err != nil {
		logger.Logger.Errorf("error in DeleteQuant: %v\n", err)
		return err
	}
	return nil
}
