package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	externalService ExternalService
}

func NewController(externalService ExternalService) *Controller {
	return &Controller{
		externalService: externalService,
	}
}

func (ctrl *Controller) GetNegaras(c *gin.Context) {
	negaras, err := ctrl.externalService.GetNegaras()
	fmt.Println("negara:", negaras)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch countries data",
			Error:   err.Error(),
		})
		return
	}

	if len(negaras) == 0 {
		c.JSON(http.StatusOK, APIResponse{
			Status:  "success",
			Message: "No countries found",
			Data:    []Negara{},
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: "Countries data retrieved successfully",
		Data:    negaras,
	})
}

func (ctrl *Controller) GetPelabuhans(c *gin.Context) {
	idNegaraStr := c.Query("id_negara")
	if idNegaraStr == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Parameter id_negara is required",
		})
		return
	}

	idNegaraStr = strings.TrimSpace(idNegaraStr)
	if idNegaraStr == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Parameter id_negara cannot be empty",
		})
		return
	}

	pelabuhans, err := ctrl.externalService.GetPelabuhans(idNegaraStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch ports data",
			Error:   err.Error(),
		})
		return
	}

	if len(pelabuhans) == 0 {
		c.JSON(http.StatusOK, APIResponse{
			Status:  "success",
			Message: "No ports found for the specified country",
			Data:    []Pelabuhan{},
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: "Ports data retrieved successfully",
		Data:    pelabuhans,
	})
}

func (ctrl *Controller) GetBarangs(c *gin.Context) {
	idPelabuhanStr := c.Query("id_pelabuhan")
	if idPelabuhanStr == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Parameter id_pelabuhan is required",
		})
		return
	}

	idPelabuhanStr = strings.TrimSpace(idPelabuhanStr)
	if idPelabuhanStr == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Parameter id_pelabuhan cannot be empty",
		})
		return
	}

	idPelabuhanInt, err := strconv.Atoi(idPelabuhanStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  "error",
			Message: "Invalid id_pelabuhan parameter, must be a number",
		})
		return
	}

	barangs, err := ctrl.externalService.GetBarangs(idPelabuhanStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch goods data",
			Error:   err.Error(),
		})
		return
	}

	var filteredBarangs []Barang
	for _, barang := range barangs {
		if barang.IDPelabuhan == idPelabuhanInt {
			filteredBarangs = append(filteredBarangs, barang)
		}
	}

	if len(filteredBarangs) == 0 {
		c.JSON(http.StatusOK, APIResponse{
			Status:  "success",
			Message: "No goods found for the specified port",
			Data:    []Barang{},
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: "Goods data retrieved successfully",
		Data:    filteredBarangs,
	})
}
