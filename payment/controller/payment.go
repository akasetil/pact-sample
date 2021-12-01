package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment/model"
	"payment/service"
)

func Payment(c *gin.Context) {
	payment := model.Payment{}
	err := c.Bind(&payment)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	paymentService := service.PaymentService{}
	err = paymentService.Payment(&payment)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
