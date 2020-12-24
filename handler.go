package Handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler struct {
	Gin *gin.Engine
}

func NewHandler() *Handler {
	handler := Handler{Gin: gin.Default()}
	return &handler
}

var Module = fx.Options(fx.Provide(NewHandler()))
