package controller

import (
	"net/http"

	"vector-tile/src/service"

	"github.com/gin-gonic/gin"
)

//* Change o controller content with ALT+C(Case sensitive) then CTRL+D o:
//! and don't forget to register the controller in main.go
//? mvt and Mvt

type MvtController struct {
	router  *gin.RouterGroup
	service *service.MvtService
}

func NewMvtController(router *gin.RouterGroup) *MvtController {
	o := &MvtController{router: router, service: service.NewMvtService()}

	mvt := o.router.Group("/mvt")
	mvt.GET("/:x/:y/:z", o.GetOne)

	return o
}

type MvtParams struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

// @Tags mvt
// @Accept json
// @Param x query int true "x"
// @Param y query int true "y"
// @Param z query int true "z"
// @Produce application/x-protobuf
// // @Success 200 {object} object{data=model.Mvt_View,meta_data=model.MetadataResponse} "OK"
// @Success 200 {string} string "OK"
// @Router /mvt/ [get]
// @Security JWT
func (o *MvtController) GetOne(ctx *gin.Context) {
	var params MvtParams
	if err := ctx.ShouldBindQuery(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid parameters")
		return
	}

	dataByte, errMessage := o.service.Get(params.X, params.Y, params.Z)
	if errMessage != "" {
		ctx.JSON(http.StatusBadRequest, errMessage)
		return
	}

	ctx.ProtoBuf(http.StatusOK, dataByte)
}
