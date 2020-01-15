package comm

type RetCode int

const (
    Success RetCode = 200
    Error   RetCode = 500

    AuthCheckFailed     RetCode = 1000
    LackUserOrTokenInfo RetCode = 10001

    DBReadError         RetCode = 2000
    DBWriteError        RetCode = 2001

    RequestParameterError   RetCode = 3000
)

var codeMessageMapping = map[RetCode]string{
    Success: "执行成功",
    Error: "未知错误",

    AuthCheckFailed: "权限校验失败，没有操作权限",
    LackUserOrTokenInfo: "权限校验失败，缺少用户名或Token",

    DBReadError: "查询数据错误",
    DBWriteError: "写入数据库错误",

    RequestParameterError: "请求参数错误",
}

func (code RetCode) Message() string {
    if msg, ok := codeMessageMapping[code]; ok {
        return msg
    }
    return "Unknown"
}
