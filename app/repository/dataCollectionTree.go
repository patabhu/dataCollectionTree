package repository

import (
	"context"
	"dataCollectionTree/app"
	"dataCollectionTree/app/node"
	"dataCollectionTree/models"
	"errors"

	"github.com/jinzhu/gorm"
)

type ToDoRepo struct {
	DbConn *gorm.DB
	tree   *node.HeadNode
}

func NewToDoRepo() app.ToDoRepoInterface {
	return &ToDoRepo{}
}

func (r *ToDoRepo) Insert(ctx context.Context, data *models.Data) (*models.ApiResponse, error) {
	var country, device string
	for _, dim := range data.Dimensions {
		switch dim.Key {
		case "country":
			country = dim.Value
		case "device":
			device = dim.Value
		}
	}
	if country == "" {
		return &models.ApiResponse{Msg: "Country parameter is required"}, errors.New("Country parameter not found")
	}
	if device == "" {
		return &models.ApiResponse{Msg: "Device parameter is required"}, errors.New("Device parameter not found")
	}
	r.tree, _ = r.tree.UpdateMetric(country, device, &data.Metrics)
	return &models.ApiResponse{Msg: "OK"}, nil
}
func (r *ToDoRepo) Query(ctx context.Context, data *models.Data) (*models.ApiResponse, error) {
	var country string
	for _, dim := range data.Dimensions {
		switch dim.Key {
		case "country":
			country = dim.Value
		}
	}
	if country == "" {
		return &models.ApiResponse{Msg: "Country parameter is required"}, errors.New("Country parameter not found")
	}
	err := r.tree.GetMetricByCountry(data, country)
	if err != nil {
		return &models.ApiResponse{Msg: "error"}, err
	}
	return &models.ApiResponse{Data: data}, nil
}
