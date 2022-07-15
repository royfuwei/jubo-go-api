package deliveryRest

import (
	"jubo-go-api/domain"

	"github.com/gin-gonic/gin"
)

type ordersHandler struct {
	ordersDelivery
	orderUcase domain.OrdersUseCase
	e          *gin.Engine
}

type ordersDelivery interface {
	UpdateById(c *gin.Context)
}

func NewOrdersHandler(e *gin.Engine, orderUcase domain.OrdersUseCase) {
	handler := &ordersHandler{
		e:          e,
		orderUcase: orderUcase,
	}
	root := e.Group("/orders")
	root.PATCH("/:id", handler.UpdateById)
}

// Update Order By Id
// @Summary Update Order By Id
// @Description Update Order By Id
// @Tags order
// @Produce json
// @Param default body domain.ReqOrderData true "update message"
// @Success 200 {object} domain.OrderDTO "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /orders/{id} [patch]
func (h *ordersHandler) UpdateById(c *gin.Context) {}
