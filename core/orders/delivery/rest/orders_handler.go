package deliveryRest

import (
	"jubo-go-api/domain"
	"jubo-go-api/domain/errcode"
	"jubo-go-api/infrastructures/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ordersHandler struct {
	ordersDelivery
	orderUcase domain.OrdersUseCase
	e          *gin.Engine
}

type ordersDelivery interface {
	UpdateById(c *gin.Context)
	FindManyByPatientId(c *gin.Context)
}

func NewOrdersHandler(e *gin.Engine, orderUcase domain.OrdersUseCase) {
	handler := &ordersHandler{
		e:          e,
		orderUcase: orderUcase,
	}
	root := e.Group("/orders")
	root.PATCH("/:id", handler.UpdateById)
	root.GET("/patients/:patientId", handler.FindManyByPatientId)
}

// Update Order By Id
// @Summary Update Order By Id
// @Description Update Order By Id
// @Tags order
// @Produce json
// @Param id path string true "Order ID"
// @Param default body domain.ReqOrderData true "update message"
// @Success 200 {object} domain.OrderDTO "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /orders/{id} [patch]
func (h *ordersHandler) UpdateById(c *gin.Context) {
	id := c.Param("id")
	var body domain.ReqOrderData
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, tools.NewResponseErr(err, errcode.ValidationFailed, c.Request.URL.Path, tools.Int(http.StatusBadRequest)))
	}
	result, uCaseErr := h.orderUcase.UpdateById(id, &body)
	if uCaseErr != nil {
		status := tools.ParseErrStatus(uCaseErr.Err)
		c.JSON(status, tools.NewResponseErr(uCaseErr.Err, uCaseErr.ErrorCode, c.Request.URL.Path, nil))
		return
	}
	c.JSON(http.StatusOK, result)
}

// Find Orders By patientId
// @Summary Find Orders By patientId
// @Description Find Orders By patientId
// @Tags order
// @Produce json
// @Param patientId path string true "Patient ID"
// @Success 200 {object} domain.RespOrders "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /orders/patients/{patientId} [get]
func (h *ordersHandler) FindManyByPatientId(c *gin.Context) {
	id := c.Param("patientId")
	result, uCaseErr := h.orderUcase.FindManyByPatientId(id)
	if uCaseErr != nil {
		status := tools.ParseErrStatus(uCaseErr.Err)
		c.JSON(status, tools.NewResponseErr(uCaseErr.Err, uCaseErr.ErrorCode, c.Request.URL.Path, nil))
		return
	}
	c.JSON(http.StatusOK, result)
}
