package comm

type ReturnError struct {
    code    RetCode
    error   error
}

func NewReturnError(code RetCode, err error) *ReturnError {
    return &ReturnError{
        code: code,
        error: err,
    }
}

func (re *ReturnError) Code() RetCode {
    return re.code
}

func (re *ReturnError) Message() string {
    if re.error == nil {
        return ""
    }
    return re.error.Error()
}

func (re *ReturnError) ResetCode(code RetCode) {
    re.code = code
}

func (re *ReturnError) ResetError(err error) {
    re.error = err
}
