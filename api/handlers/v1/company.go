package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"bitbucket.org/udevs/example_api_gateway/api/models"
	"bitbucket.org/udevs/example_api_gateway/genproto/company_service"
	"bitbucket.org/udevs/example_api_gateway/pkg/util"
)

// Create Company godoc
// @ID create-company
// @Router /v1/company [POST]
// @Summary create company
// @Description Create company
// @Tags Company
// @Accept json
// @Produce json
// @Param company body models.CreateCompany true "company"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateCompany(c *gin.Context) {
	var company models.CreateCompany

	if err := c.BindJSON(&company); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	resp, err := h.services.CompanyService().Create(
		context.Background(),
		&company_service.CreateCompany{
			Name: company.Name,
		},
	)

	if !handleError(h.log, c, err, "error while creating company") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// Get company godoc
// @ID get-company
// @Router /v1/company/{company_id} [GET]
// @Summary get company
// @Description Get company
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "company_id"
// @Success 200 {object} models.ResponseModel{data=models.Company} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetCompany(c *gin.Context) {
	var company models.Company
	company_id := c.Param("company_id")

	if !util.IsValidUUID(company_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "company id is not valid", errors.New("company id is not valid"))
		return
	}

	resp, err := h.services.CompanyService().Get(
		context.Background(),
		&company_service.CompanyId{
			Id: company_id,
		},
	)

	if !handleError(h.log, c, err, "error while getting company") {
		return
	}

	err = ParseToStruct(&company, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", company)
}

// Get All company godoc
// @ID get-all-company
// @Router /v1/company [GET]
// @Summary get all company
// @Description Get All company
// @Tags Company
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllCompanyResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllCompanies(c *gin.Context) {
	var companies models.GetAllCompanyResponse

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.CompanyService().GetAll(
		context.Background(),
		&company_service.GetAllCompanyRequest{
			Limit:  uint32(limit),
			Offset: uint32(offset),
			Name:   c.Query("name"),
		},
	)

	if !handleError(h.log, c, err, "error while getting all companies") {
		return
	}

	err = ParseToStruct(&companies, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing to struct", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", companies)
}

// Update company godoc
// @ID update-company
// @Router /v1/company [PUT]
// @Summary update company
// @Description Update company
// @Tags Company
// @Accept json
// @Produce json
// @Param company body models.Company true "company"
// @Success 200 {object} models.ResponseModel{data=models.MsgResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateCompany(c *gin.Context) {
	var company models.Company

	if err := c.BindJSON(&company); err != nil {
		h.handleErrorResponse(c, 400, "error while binging json", err)
		return
	}

	_, err := h.services.CompanyService().Update(
		context.Background(),
		&company_service.Company{
			Id:   company.ID,
			Name: company.Name,
		},
	)

	if !handleError(h.log, c, err, "error while creating company") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", &models.MsgResponse{Msg: "Updated"})
}

// Delete company godoc
// @ID delete-company
// @Router /v1/company/{company_id} [DELETE]
// @Summary delete company
// @Description Delete company
// @Tags Company
// @Accept json
// @Produce json
// @Param company_id path string true "company_id"
// @Success 200 {object} models.ResponseModel{data=models.MsgResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteCompany(c *gin.Context) {
	company_id := c.Param("company_id")

	if !util.IsValidUUID(company_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "company id is not valid", errors.New("company id is not valid"))
		return
	}

	_, err := h.services.CompanyService().Delete(
		context.Background(),
		&company_service.CompanyId{
			Id: company_id,
		},
	)

	if !handleError(h.log, c, err, "error while deleting company") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", &models.MsgResponse{Msg: "Deleted"})
}
