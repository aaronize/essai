package comm

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type ResponseUniform struct {
    Meta    Meta        `json:"meta"`
    Data    interface{} `json:"data"`
}

type Meta struct {
    Code    RetCode     `json:"code"`
    Message string      `json:"message"`
}

type WrappedGin struct {
    ctx     *gin.Context
}

func GinWrapper(ctx *gin.Context) *WrappedGin {
    return &WrappedGin{
        ctx: ctx,
    }
}

func (g *WrappedGin) Response(code RetCode, data interface{}, patchMessage ...string) {
    if data == nil {
        // 统一返回一个空数组
        data = []interface{}{}
    }

    patch := ""
    for _, pm := range patchMessage {
        patch += ", " + pm
    }
    resp := ResponseUniform{
        Meta: Meta{
            Code: code,
            Message: code.Message() + patch,
        },
        Data: data,
    }

    g.ctx.JSON(http.StatusOK, resp)
    g.ctx.Abort()
}
