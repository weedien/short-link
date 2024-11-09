package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func New(redisClient *redis.Client, excludes []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//// 排除不需要鉴权的接口
		//baseUrl := base.GetConfig().Server.BasePath
		//for _, exclude := range excludes {
		//	if baseUrl+exclude == c.Path() {
		//		return c.Next()
		//	}
		//}
		//// 获取请求头中的 Authorization 字段
		//token := c.Get("Authorization")
		//if token == "" {
		//	return errno.ErrUnauthorized
		//}
		//// 查询 token 对应的用户信息 基于redis
		//user, err := redisClient.Get(c.Context(), token).Result()
		//if err != nil {
		//	return errno.ErrUnauthorized
		//}
		//userinfo := make(map[string]interface{})
		//if err = sonic.Unmarshal([]byte(user), &userinfo); err != nil {
		//	return err
		//}
		//// 将用户信息存入上下文
		//c.Locals("user", userinfo)
		//c.Locals("username", userinfo["username"])
		c.Locals("username", "default")
		return c.Next()
	}
}
