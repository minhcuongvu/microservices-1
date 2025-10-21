package configurations

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/minhcuongvu/microservices-1/internal/pkg/logger"

	getting_product_by_id "github.com/minhcuongvu/microservices-1/internal/services/product_service/product/features/getting_product_by_id/v1/endpoints"
)

func ConfigEndpoints(log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	// creating_product.MapRoute(validator, log, echo, ctx)
	// deleting_product.MapRoute(validator, log, echo, ctx)
	getting_product_by_id.MapRoute(log, echo, ctx)
	// getting_products.MapRoute(validator, log, echo, ctx)
	// searching_product.MapRoute(validator, log, echo, ctx)
	// updating_product.MapRoute(validator, log, echo, ctx)
}
