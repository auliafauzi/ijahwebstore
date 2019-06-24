package middleware

import (
    "github.com/kataras/iris"
    "ijahwebstore/service/userservice"
    "github.com/kataras/iris/context"
    "ijahwebstore/tools"
)

func TokenValidation(ctx iris.Context){
    accessToken := ctx.GetHeader("x-access-token")
    if userservice.TokenChecker(accessToken) {
        ctx.Next()
        return
    }
    tools.ResponseJSON(ctx, 401, context.Map{
        "message": "Token required or not valid",
        "result": false,
    })
}