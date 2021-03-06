package router

import (
	"supersign/internal/api"
	v1 "supersign/internal/api/v1"
	"supersign/internal/model"
	"supersign/internal/server/middleware"
	"supersign/pkg/log"

	"github.com/gin-gonic/gin"
)

func registerAppleDeveloperRouter(store *model.Store, r *gin.RouterGroup) {
	appleDeveloper := v1.AppleDeveloper{
		Base: api.New(store, log.New("AppleDeveloper").L()),
	}
	appleDeveloperR := r.Group("/appleDeveloper", middleware.JWT())

	{
		appleDeveloperR.POST("", appleDeveloper.Upload)
		appleDeveloperR.DELETE("", appleDeveloper.Del)
		appleDeveloperR.GET("", appleDeveloper.List)
		appleDeveloperR.PATCH("", appleDeveloper.Update)
	}
}
