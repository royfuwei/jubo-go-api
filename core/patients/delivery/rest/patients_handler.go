package deliveryRest

import (
	"jubo-go-api/domain"
	"jubo-go-api/domain/errcode"
	"jubo-go-api/infrastructures/tools"
	"net/http"

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
	root.POST("/", handler.AddOne)
	root.POST("/:id/order", handler.AddOrderById)
}

// Find All patients
// @Summary Find All patients
// @Description Find All patients
// @Tags patients
// @Produce json
// @Success 200 {object} domain.RespFindAll "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /patients [get]
func (h *patientsHandler) FindAll(c *gin.Context) {
	result, uCaseErr := h.patientsUseCase.FindAll()
	if uCaseErr != nil {
		status := tools.ParseErrStatus(uCaseErr.Err)
		c.JSON(status, tools.NewResponseErr(uCaseErr.Err, uCaseErr.ErrorCode, c.Request.URL.Path, nil))
		return
	}
	c.JSON(http.StatusOK, result)
}

// Find patient by Id
// @Summary Find patient by Id
// @Description Find patient by Id
// @Tags patients
// @Produce json
// @Param id path string true "Patient ID"
// @Success 200 {object} domain.RespPatientData "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /patients/{id} [get]
func (h *patientsHandler) FindById(c *gin.Context) {
	id := c.Param("id")
	result, uCaseErr := h.patientsUseCase.FindById(id)
	if uCaseErr != nil {
		status := tools.ParseErrStatus(uCaseErr.Err)
		c.JSON(status, tools.NewResponseErr(uCaseErr.Err, uCaseErr.ErrorCode, c.Request.URL.Path, nil))
		return
	}
	c.JSON(http.StatusOK, result)
}

// Create Order By patientId
// @Summary Create Order By patientId
// @Description Create Order By patientId
// @Tags patients
// @Produce json
// @Param id path string true "Patient ID"
// @Param default body domain.ReqOrderData true "Create Patient's Order request body"
// @Success 201 {object} domain.RespPatientData "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /patients/{id}/order [post]
func (h *patientsHandler) AddOrderById(c *gin.Context) {
	id := c.Param("id")
	var body domain.ReqOrderData
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, tools.NewResponseErr(err, errcode.ValidationFailed, c.Request.URL.Path, tools.Int(http.StatusBadRequest)))
		return
	}
	result, uCaseErr := h.patientsUseCase.AddOrderById(id, &body)
	if uCaseErr != nil {
		status := tools.ParseErrStatus(uCaseErr.Err)
		c.JSON(status, tools.NewResponseErr(uCaseErr.Err, uCaseErr.ErrorCode, c.Request.URL.Path, nil))
		return
	}
	c.JSON(http.StatusCreated, result)
}

// Create Patient
// @Summary Create Patient
// @Description Create Patient
// @Tags patients
// @Produce json
// @Param default body domain.ReqAddOne true "Create Patient request body"
// @Success 201 {object} domain.RespPatientData "success response"
// @Failure 400 {object} domain.ResponseError "請求的body、header驗證失敗"
// @Router /patients [post]
func (h *patientsHandler) AddOne(c *gin.Context) {
	var body domain.ReqAddOne
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, tools.NewResponseErr(err, errcode.ValidationFailed, c.Request.URL.Path, tools.Int(http.StatusBadRequest)))
		return
	}
	result, uCaseErr := h.patientsUseCase.AddOne(&body)
	if uCaseErr != nil {
		status := tools.ParseErrStatus(uCaseErr.Err)
		c.JSON(status, tools.NewResponseErr(uCaseErr.Err, uCaseErr.ErrorCode, c.Request.URL.Path, nil))
		return
	}
	c.JSON(http.StatusCreated, result)
}
