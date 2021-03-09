package app

import (
	"context"
	"dataCollectionTree/models"
)

type ToDoRepoInterface interface {
	Insert(ctx context.Context, t *models.Data) (*models.ApiResponse, error)
	Query(ctx context.Context, t *models.Data) (*models.ApiResponse, error)
}
