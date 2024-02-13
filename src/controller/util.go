package controller

import (
	"net/http"
	"time"

	"vector-tile/src/model"

	"github.com/gin-gonic/gin"
)

func SetMetadataResponse(ctx *gin.Context, startTime time.Time, resp *model.Response) {
	code := http.StatusOK
	resp.Metadata.Status = true
	if resp.Metadata.Message == "" {
		resp.Metadata.Message = "OK"
	}

	if resp.Metadata.Message != "OK" {
		resp.Metadata.Status = false
	}

	resp.Metadata.TimeExecution = time.Since(startTime).String()
	ctx.JSON(code, resp)
}
