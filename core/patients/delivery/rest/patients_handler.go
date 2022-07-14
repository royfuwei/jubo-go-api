package deliveryRest

import "github.com/gin-gonic/gin"

type patientsHandler struct {
	patientsDelivery
	e *gin.Engine
}

type patientsDelivery interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	CreateOrderById(c *gin.Context)
}

func NewPatientsHandler(e *gin.Engine) {
	handler := &patientsHandler{
		e: e,
	}
	root := e.Group("/patients")
	root.GET("/", handler.FindAll)
	root.GET("/:id", handler.FindById)
	root.POST("/:id", handler.CreateOrderById)
}

// Find All patients
// @Summary Find All patients
// @Description Find All patients
// @Tags patients
// @Produce json
// @Success 200 {object} domain.RespFindAll "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /patients [get]
func (h *patientsHandler) FindAll(c *gin.Context) {}

// Find patient by Id
// @Summary Find patient by Id
// @Description Find patient by Id
// @Tags patients
// @Produce json
// @Success 200 {object} domain.RespPatientData "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /patients/{id} [get]
func (h *patientsHandler) FindById(c *gin.Context) {}

// Create Order By patientId
// @Summary Create Order By patientId
// @Description Create Order By patientId
// @Tags patients
// @Produce json
// @Success 200 {object} domain.OrderDTO "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /patients/{id} [post]
func (h *patientsHandler) CreateOrderById(c *gin.Context) {}
