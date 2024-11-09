package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"shortlink/internal/user/app/user"
	"shortlink/internal/user/app/user/command"
	"shortlink/internal/user/app/user/query"
	"shortlink/internal/user/trigger/rest/dto/req"
	"shortlink/internal/user/trigger/rest/dto/resp"
)

type UserApi struct {
	app user.Application
}

func NewUserApi(app user.Application, router fiber.Router) {
	api := &UserApi{app: app}

	userRouter := router.Group("/users")
	userRouter.Get("/username/:username", api.GetUserByUsername)
	userRouter.Get("/actual/:username", api.GetUserByUsernameWithoutMask)
	userRouter.Get("/exist", api.CheckUserExist)
	userRouter.Get("/check-login", api.CheckLogin)
	userRouter.Post("/register", api.Register)
	userRouter.Post("/login", api.Login)
	userRouter.Post("/logout", api.Logout)
	userRouter.Put("", api.Update)
	userRouter.Delete("/username/:username", api.Delete)
}

// GetUserByUsername 根据用户名查询用户信息
func (h UserApi) GetUserByUsername(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	response := resp.UserResp{}
	res, err := h.app.Queries.GetUser.Handle(c.Context(), username)
	if err != nil {
		return err
	}
	if err = copier.Copy(&response, res); err != nil {
		return err
	}
	return c.JSON(response)
}

// GetUserByUsernameWithoutMask 根据用户名查询无脱敏用户信息
func (h UserApi) GetUserByUsernameWithoutMask(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	response := resp.UserActualResp{}
	res, err := h.app.Queries.GetUser.Handle(c.Context(), username)
	if err != nil {
		return err
	}
	if err = copier.Copy(&response, res); err != nil {
		return err
	}
	return c.JSON(response)
}

// CheckUserExist 查询用户是否存在
func (h UserApi) CheckUserExist(c *fiber.Ctx) error {
	username := c.Query("username")
	exist, err := h.app.Queries.CheckUserExist.Handle(c.Context(), username)
	if err != nil {
		return err
	}
	return c.JSON(exist)
}

// Register 用户注册
func (h UserApi) Register(c *fiber.Ctx) error {
	reqParam := req.UserRegisterReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		return err
	}
	cmd := command.UserRegisterCommand{}
	if err := copier.Copy(&cmd, &reqParam); err != nil {
		return err
	}
	if err := h.app.Commands.UserRegister.Handle(c.Context(), cmd); err != nil {
		return err
	}
	return nil
}

// Update 修改用户信息
func (h UserApi) Update(c *fiber.Ctx) error {
	reqParam := req.UserUpdateReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		return err
	}
	cmd := command.UpdateUserCommand{}
	if err := copier.Copy(&cmd, &reqParam); err != nil {
		return err
	}
	if err := h.app.Commands.UpdateUser.Handle(c.Context(), cmd); err != nil {
		return err
	}
	return nil
}

// Login 用户登录
func (h UserApi) Login(c *fiber.Ctx) error {
	reqParam := req.UserLoginReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		return err
	}
	response := resp.UserLoginResp{}
	cmd := &command.UserLoginCommand{}
	if err := copier.Copy(cmd, &reqParam); err != nil {
		return err
	}
	if err := h.app.Commands.UserLogin.Handle(c.Context(), cmd); err != nil {
		return err
	}
	response.Token = cmd.ExecutionResult()
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    response.Token,
		HTTPOnly: true,
		MaxAge:   3600,
	})
	return c.JSON(response)
}

// CheckLogin 检查用户是否登录
func (h UserApi) CheckLogin(c *fiber.Ctx) error {
	username := c.Query("username")
	token := c.Query("token")
	q := query.CheckLogin{
		Username: username,
		Token:    token,
	}
	login, err := h.app.Queries.CheckLogin.Handle(c.Context(), q)
	if err != nil {
		return err
	}
	return c.JSON(login)
}

// Logout 用户登出
func (h UserApi) Logout(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	token := c.Query("token")
	cmd := command.UserLogoutCommand{
		Username: username,
		Token:    token,
	}
	if err := h.app.Commands.UserLogout.Handle(c.Context(), cmd); err != nil {
		return err
	}
	return nil
}

// Delete 删除用户
func (h UserApi) Delete(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	if err := h.app.Commands.DeleteUser.Handle(c.Context(), username); err != nil {
		return err
	}
	return nil
}
