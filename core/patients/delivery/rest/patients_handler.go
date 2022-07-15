package deliveryRest

import (
	"jubo-go-api/domain"

	"github.com/gin-gonic/gin"
)

type patientsHandler struct {
	patientsDelivery
	patientsUseCase domain.PatientUseCase
	e               *gin.Engine
}

type patientsDelivery interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	AddOrderById(c *gin.Context)
	AddOne(c *gin.Context)
}

func NewPatientsHandler(e *gin.Engine, patientsUseCase domain.PatientUseCase) {
	handler := &patientsHandler{
		e: e,
	}
	root := e.Group("/patients")
	root.GET("/", handler.FindAll)
	root.GET("/:id", handler.FindById)
	root.POST("/:id", handler.AddOrderById)
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
// @Success 200 {object} domain.RespPatientData "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /patients/{id}/order [post]
func (h *patientsHandler) AddOrderById(c *gin.Context) {}

// Create Patient
// @Summary Create Patient
// @Description Create Patient
// @Tags patients
// @Produce json
// @Success 200 {object} domain.RespPatientData "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /patients [post]
func (h *patientsHandler) AddOne(c *gin.Context) {}
