package v1

import (
	"net/http"

	"example.com/log"
	"github.com/gin-gonic/gin"
)

type PostCallRequest struct {
	Data string `form:"data_id" json:"data_id" binding:"required"`
}

func PostCall(c *gin.Context) {
	var req PostCallRequest
	log.Infof("PostCallRequest: %v", req)
	if err := c.ShouldBind(&req); err != nil {
		log.Error("Bind PostCallRequest failed, err: ", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, struct{}{})
		return
	}
	log.Info(req)
	c.PureJSON(http.StatusOK, req)
}
