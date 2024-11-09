package httperr

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"shortlink/internal/base/errno"
)

type Response struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func RespondWithError(c *fiber.Ctx, err error) error {
	// 在这里去处理一些很特殊的error，比如能和http状态码直接对应
	//if errors.Is(err, errno.TooManyRequests) {
	//	// 请求频率过高
	//	r := Response{
	//		Status: "Too many requests",
	//		Msg:    err.Error(),
	//	}
	//	return c.Status(fiber.StatusTooManyRequests).JSON(r)
	//}
	//if errors.Is(err, errno.RouteNotFound) {
	//	// 请求频率过高
	//	r := Response{
	//		Status: "Route not found",
	//		Msg:    err.Error(),
	//	}
	//	return c.Status(fiber.StatusNotFound).JSON(r)
	//}

	var fiberError *fiber.Error
	if ok := errors.As(err, &fiberError); ok {
		// fiber框架的错误
		r := Response{
			Status: "error",
			Msg:    fiberError.Message,
		}
		return c.Status(fiberError.Code).JSON(r)
	}

	var slugError errno.SlugError
	if ok := errors.As(err, &slugError); !ok {
		// 未定义的内部异常
		r := Response{
			Status: "Internal server error",
			Msg:    "系统异常",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}

	switch slugError.Type() {
	case errno.ErrorTypeAuthorization:
		// 未授权
		r := Response{
			Status: "Unauthorized",
			Msg:    slugError.Error(),
		}
		return c.Status(fiber.StatusUnauthorized).JSON(r)
	case errno.ErrorTypeRequestParam:
		// 请求参数错误
		r := Response{
			Status: "Bad request",
			Msg:    slugError.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(r)
	case errno.ErrorTypeResourceNotFound:
		// 资源未找到，返回 notfound 页面
		return c.Status(fiber.StatusNotFound).SendFile("../resources/notfound.html")
	case errno.ErrorTypeServiceError:
		// 业务异常
		r := Response{
			Status: "Business error",
			Msg:    slugError.Error(),
		}
		return c.Status(fiber.StatusOK).JSON(r)
	default:
		// 未定义的内部异常
		r := Response{
			Status: "Internal server error",
			Msg:    "系统异常",
		}
		return c.Status(fiber.StatusInternalServerError).JSON(r)
	}
}

// ErrorHandler 全局错误处理（fiber专用）
func ErrorHandler(c *fiber.Ctx, err error) error {
	return RespondWithError(c, err)
}
