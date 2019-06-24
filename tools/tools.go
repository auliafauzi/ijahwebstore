package tools

import (
    "github.com/kataras/iris"
    "encoding/base64"
                )

func ResponseJSON(ctx iris.Context, code int, json map[string]interface{}) {
    if code == 0 {
        code = 200
    }
    ctx.StatusCode(code)
    ctx.JSON(json)
}

func Base64Encode(str string) string {
    return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) (string, bool) {
    data, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        return "", true
    }
    return string(data), false
}