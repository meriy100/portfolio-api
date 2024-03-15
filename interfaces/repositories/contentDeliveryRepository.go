package repositories

import (
	"context"
	"github.com/meriy100/portfolio-api/adapters"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type ContentDeliveryRepository struct {
}

func NewContentDeliveryRepository() ports.ContentDeliveryRepository {
	return &ContentDeliveryRepository{}
}

func (r *ContentDeliveryRepository) Deploy(ctx context.Context) error {
	return adapters.InitialVercelHookRequest(ctx)
}
