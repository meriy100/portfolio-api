package ports

import "context"

type ContentDeliveryRepository interface {
	Deploy(ctx context.Context) error
}
