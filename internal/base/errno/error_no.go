package errno

var (
	ErrorTypeRequestParam     ErrorType = "request-param"      // 400
	ErrorTypeAuthorization    ErrorType = "authorization"      // 401
	ErrorTypeResourceNotFound ErrorType = "resource-not-found" // 404
	ErrorTypeExternalError    ErrorType = "external-error"     // 500
	ErrorTypeServiceError     ErrorType = "service-error"      // 500/200
)

var (
	RequestParamError  SlugError = SlugError{errorType: ErrorTypeRequestParam, msg: "请求参数错误"}
	AuthorizationError SlugError = SlugError{errorType: ErrorTypeAuthorization, msg: "鉴权失败"}
	ExternalError      SlugError = SlugError{errorType: ErrorTypeExternalError, msg: "系统异常"}

	ErrUnauthorized = SlugError{errorType: ErrorTypeAuthorization, msg: "未授权"}

	// 短链接异常

	LinkInvalidOriginalUrl     = SlugError{errorType: ErrorTypeRequestParam, msg: "不合法的原始链接"}
	LinkInvalidValidType       = SlugError{errorType: ErrorTypeRequestParam, msg: "不合法的有效期类型"}
	LinkEndTimeBeforeStartTime = SlugError{errorType: ErrorTypeRequestParam, msg: "结束时间早于开始时间"}

	LinkGroupEmpty           = SlugError{errorType: ErrorTypeResourceNotFound, msg: "分组下没有短链接"}
	LinkAlreadyExists        = SlugError{errorType: ErrorTypeServiceError, msg: "短链接已存在"}
	LinkNotExists            = SlugError{errorType: ErrorTypeServiceError, msg: "短链接不存在"}
	LinkDisabled             = SlugError{errorType: ErrorTypeServiceError, msg: "短链接已停用"}
	LinkExpired              = SlugError{errorType: ErrorTypeServiceError, msg: "短链接已过期"}
	LinkForbidden            = SlugError{errorType: ErrorTypeServiceError, msg: "短链接已禁用"}
	LinkReserved             = SlugError{errorType: ErrorTypeServiceError, msg: "短链接已保留"}
	LinkInvalidStatus        = SlugError{errorType: ErrorTypeServiceError, msg: "不合法的短链接状态"}
	LinkTooManyAttempts      = SlugError{errorType: ErrorTypeServiceError, msg: "多次尝试生成唯一短链接失败"}
	LinkDisallowedDomain     = SlugError{errorType: ErrorTypeServiceError, msg: "不支持跳转的域名"}
	LinkGroupLinkCountExceed = SlugError{errorType: ErrorTypeServiceError, msg: "超过组内短链接数量限制"}

	// 自定义系统异常

	LockAcquireFailed = SlugError{errorType: ErrorTypeExternalError, msg: "锁获取失败"}
	LockReleaseFailed = SlugError{errorType: ErrorTypeExternalError, msg: "锁释放失败"}

	RedisError       = SlugError{errorType: ErrorTypeExternalError, msg: "Redis异常"}
	RedisKeyNotExist = SlugError{errorType: ErrorTypeExternalError, msg: "Redis key不存在"}
)

type ErrorType string

// 我对全局异常处理的理解
// 1. 划分所有异常为3类：请求异常、服务异常、系统异常
// 2. 如果需要抛出请求异常和服务异常，需要进行自定义
// 3. 所有未定义的异常，比如直接使用 errors.New 创建的，都作为系统异常
// 4. 那么，只有领域内部的异常需要自定义，因为只有在领域内才会去验证参数的合法性，或者进行业务逻辑的处理

// SlugError 自定义异常
type SlugError struct {
	errorType ErrorType
	msg       string
}

func (s SlugError) Error() string {
	return s.msg
}

func (s SlugError) Type() ErrorType {
	return s.errorType
}

func NewRequestError(msg string) SlugError {
	return SlugError{errorType: ErrorTypeRequestParam, msg: msg}
}
