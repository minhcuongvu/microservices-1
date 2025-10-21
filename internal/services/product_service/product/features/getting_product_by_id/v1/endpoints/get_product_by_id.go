package endpoints

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minhcuongvu/microservices-1/internal/pkg/logger"
)

func MapRoute(log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/products")
	group.GET("/:id", getProductByID(log, ctx))
}

// GetProductByID
// @Tags        Products
// @Summary     Get product
// @Description Get product by id
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Product ID"
// @Success     200 {object} dtos.GetProductByIdResponseDto
// @Security ApiKeyAuth
// @Router      /api/v1/products/{id} [get]
func getProductByID(log logger.ILogger, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		resp := map[string]interface{}{
			"method":      req.Method,
			"url":         req.URL.String(),
			"headers":     req.Header,
			"remote_addr": req.RemoteAddr,
			"context":     req.Context(),
		}
		return c.JSON(http.StatusOK, resp)

		// request := &dtosv1.GetProductByIdRequestDto{}
		// if err := c.Bind(request); err != nil {
		// 	log.Warn("Bind", err)
		// 	return echo.NewHTTPError(http.StatusBadRequest, err)
		// }
		//
		// query := queriesv1.NewGetProductById(request.ProductId)
		//
		// if err := validator.StructCtx(ctx, query); err != nil {
		// 	log.Warn("validate", err)
		// 	return echo.NewHTTPError(http.StatusBadRequest, err)
		// }
		//
		// queryResult, err := mediatr.Send[*queriesv1.GetProductById, *dtosv1.GetProductByIdResponseDto](ctx, query)
		//
		// if err != nil {
		// 	log.Warn("GetProductById", err)
		// 	return echo.NewHTTPError(http.StatusBadRequest, err)
		// }
		//
		// return c.JSON(http.StatusOK, queryResult)
	}
}
