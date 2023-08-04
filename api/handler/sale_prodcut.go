package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// Create sale_product godoc
// @ID create_sale_product
// @Router /sale_product [POST]
// @Summary Create SaleProduct
// @Description Create SaleProduct
// @Tags SaleProduct
// @Accept json
// @Procedure json
// @Param SaleProduct body models.CreateSaleProduct true "CreateSaleProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateSaleProduct(c *gin.Context) {

	var createSaleProduct models.CreateSaleProduct
	err := c.ShouldBindJSON(&createSaleProduct)
	if err != nil {
		h.handlerResponse(c, "error SaleProduct should bind json", http.StatusBadRequest, err.Error())
		return
	}
	if createSaleProduct.DiscountType == "fixed" {
		if createSaleProduct.Discount < int64(createSaleProduct.ProductPrice) {
			createSaleProduct.PriceWithDiscount += createSaleProduct.ProductPrice - float64(createSaleProduct.Discount)
			createSaleProduct.DiscountPrice += createSaleProduct.ProductPrice - createSaleProduct.PriceWithDiscount
		}
		fmt.Println("fixed discount cannot be greater than product price")
		// salee.TotalPrice += resp.PriceWithDiscount * float64(resp.Count)
		// salee.TotalCount += resp.Count
	}
	if createSaleProduct.DiscountType == "percent" {
		createSaleProduct.PriceWithDiscount += (createSaleProduct.ProductPrice * float64(createSaleProduct.Discount)) / 100
		createSaleProduct.DiscountPrice += createSaleProduct.ProductPrice - createSaleProduct.PriceWithDiscount
		// salee.TotalPrice = +resp.PriceWithDiscount * float64(resp.Count)
		// salee.TotalCount += resp.Count
	}
	id, err := h.strg.SaleProduct().Create(c.Request.Context(), &createSaleProduct)
	if err != nil {
		h.handlerResponse(c, "storage.SaleProduct.create", http.StatusInternalServerError, err.Error())
		return
	}
	resp, err := h.strg.SaleProduct().GetByID(c.Request.Context(), &models.SaleProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.SaleProduct.getById", http.StatusInternalServerError, err.Error())
		return
	}

	salee, err := h.strg.Sale().GetByID(c.Request.Context(), &models.SalePrimaryKey{Id: resp.SaleID})
	if err != nil {
		h.handlerResponse(c, "storage.Sale.getById", http.StatusInternalServerError, err.Error())
		return
	}

	salee.TotalCount = salee.TotalCount + resp.Count
	salee.TotalPrice = salee.TotalPrice + resp.ProductPrice*float64(resp.Count)

	_, err = h.strg.Sale().Update(c.Request.Context(), &models.UpdateSale{
		Id:         salee.Id,
		TotalPrice: salee.TotalPrice,
		TotalCount: salee.TotalCount,
	})
	if err != nil {
		h.handlerResponse(c, "storage.Sale.Update", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create SaleProduct resposne", http.StatusCreated, resp)
}

// GetByID sale_product godoc
// @ID get_by_id_sale_product
// @Router /sale_product/{id} [GET]
// @Summary Get By ID SaleProduct
// @Description Get By ID SaleProduct
// @Tags SaleProduct
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdSaleProduct(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.SaleProduct().GetByID(c.Request.Context(), &models.SaleProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.SaleProduct.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id SaleProduct resposne", http.StatusOK, resp)
}

// GetList sale_product godoc
// @ID get_list_sale_product
// @Router /sale_product [GET]
// @Summary Get List SaleProduct
// @Description Get List SaleProduct
// @Tags SaleProduct
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListSaleProduct(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list SaleProduct offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list SaleProduct limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.SaleProduct().GetList(c.Request.Context(), &models.SaleProductGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.SaleProduct.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list SaleProduct resposne", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update sale_product godoc
// @ID update_sale_product
// @Router /sale_product/{id} [PUT]
// @Summary Update SaleProduct
// @Description Update SaleProduct
// @Tags SaleProduct
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param SaleProduct body models.UpdateSaleProduct true "UpdateSaleProductRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateSaleProduct(c *gin.Context) {

	var (
		id                string = c.Param("id")
		updateSaleProduct models.UpdateSaleProduct
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateSaleProduct)
	if err != nil {
		h.handlerResponse(c, "error SaleProduct should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateSaleProduct.Id = id
	rowsAffected, err := h.strg.SaleProduct().Update(c.Request.Context(), &updateSaleProduct)
	if err != nil {
		h.handlerResponse(c, "storage.SaleProduct.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.SaleProduct.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.SaleProduct().GetByID(c.Request.Context(), &models.SaleProductPrimaryKey{Id: updateSaleProduct.Id})
	if err != nil {
		h.handlerResponse(c, "storage.SaleProduct.getById", http.StatusInternalServerError, err.Error())
		return
	}
	h.handlerResponse(c, "create SaleProduct resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete sale_product godoc
// @ID delete_sale_product
// @Router /sale_product/{id} [DELETE]
// @Summary Delete SaleProduct
// @Description Delete SaleProduct
// @Tags SaleProduct
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteSaleProduct(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.SaleProduct().Delete(c.Request.Context(), &models.SaleProductPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.SaleProduct.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create SaleProduct resposne", http.StatusNoContent, nil)
}
