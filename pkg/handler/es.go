package handler

import (
	"es-tranform/pkg/model"
	"es-tranform/pkg/repo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type EsHandlers struct {
	es *repo.Repo
}

func NewEsHandlers(es *repo.Repo) *EsHandlers {
	return &EsHandlers{es: es}
}

func (h *EsHandlers) CreateQuery(c *gin.Context) {

	var req model.RequestQuery
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}
	result, err := h.es.CreateQuery(req)
	if err != nil {
		log.Println("err sending request", err)
	}

	c.JSON(http.StatusOK, result)

}

func (h *EsHandlers) CreateQueryFLex(c *gin.Context) {
	var req model.Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}
	result, err := h.es.CreateQueryFlex(req)
	if err != nil {
		log.Println("err sending request", err)
	}

	c.JSON(http.StatusOK, result)

}
