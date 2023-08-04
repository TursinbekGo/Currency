package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// Create sale godoc
// @ID create_sale
// @Router /sale [POST]
// @Summary Create Sale
// @Description Create Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param Sale body models.CreateSale true "CreateSaleRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateSale(c *gin.Context) {

	var createSale models.CreateSale
	err := c.ShouldBindJSON(&createSale)
	if err != nil {
		h.handlerResponse(c, "error Sale should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Sale().Create(c.Request.Context(), &createSale)
	if err != nil {
		h.handlerResponse(c, "storage.Sale.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Sale().GetByID(c.Request.Context(), &models.SalePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Sale.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Sale resposne", http.StatusCreated, resp)
}

// GetByID sale godoc
// @ID get_by_id_sale
// @Router /sale/{id} [GET]
// @Summary Get By ID Sale
// @Description Get By ID Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdSale(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Sale().GetByID(c.Request.Context(), &models.SalePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Sale.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Sale resposne", http.StatusOK, resp)
}

// GetList sale godoc
// @ID get_list_sale
// @Router /sale [GET]
// @Summary Get List Sale
// @Description Get List Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListSale(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Sale offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Sale limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Sale().GetList(c.Request.Context(), &models.SaleGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Sale.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Sale resposne", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update sale godoc
// @ID update_sale
// @Router /sale/{id} [PUT]
// @Summary Update Sale
// @Description Update Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Sale body models.UpdateSale true "UpdateSaleRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateSale(c *gin.Context) {

	var (
		id         string = c.Param("id")
		updateSale models.UpdateSale
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateSale)
	if err != nil {
		h.handlerResponse(c, "error Sale should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateSale.Id = id
	rowsAffected, err := h.strg.Sale().Update(c.Request.Context(), &updateSale)
	if err != nil {
		h.handlerResponse(c, "storage.Sale.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Sale.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Sale().GetByID(c.Request.Context(), &models.SalePrimaryKey{Id: updateSale.Id})
	if err != nil {
		h.handlerResponse(c, "storage.Sale.getById", http.StatusInternalServerError, err.Error())
		return
	}

	sale_product, err := h.strg.SaleProduct().GetByID(c.Request.Context(), &models.SaleProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.SaleProduct.getById", http.StatusInternalServerError, err.Error())
		return
	}

	if sale_product.DiscountType == "fixed" {

		resp.TotalPrice += sale_product.PriceWithDiscount * float64(sale_product.Count)
		resp.TotalCount += sale_product.Count
	}
	if sale_product.DiscountType == "percent" {

		resp.TotalPrice += sale_product.PriceWithDiscount * float64(sale_product.Count)
		resp.TotalCount += sale_product.Count
	}

	h.handlerResponse(c, "create Sale resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete sale godoc
// @ID delete_sale
// @Router /sale/{id} [DELETE]
// @Summary Delete Sale
// @Description Delete Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteSale(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Sale().Delete(c.Request.Context(), &models.SalePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Sale.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Sale resposne", http.StatusNoContent, nil)
}
