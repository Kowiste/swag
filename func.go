package main

import (
	"net/http"

	"github.com/Kowiste/swag/error"

	"github.com/gin-gonic/gin"
)

type ErrorTest struct {
	Err    error
	Status int
}

// Test godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} errorTest.ErrorTx
// @Router /test1/ [get]
func Test1(c *gin.Context) {
	res := new(gin.Error)
	c.JSON(http.StatusOK, res)
}

// Test2 godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} errorTest.ErrorTx
// @Router /test2/ [get]
func Test2(c *gin.Context) {
	res := new(errorTest.ErrorTx)
	c.JSON(http.StatusOK, res)
}

// Test3 godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} ErrorTest
// @Router /test3/ [get]
func Test3(c *gin.Context) {
	res := new(ErrorTest)
	c.JSON(http.StatusOK, res)
}
