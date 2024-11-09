package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"shortlink/internal/user/app/group"
	"shortlink/internal/user/app/group/command"
	"shortlink/internal/user/trigger/rest/dto/req"
	"shortlink/internal/user/trigger/rest/dto/resp"
)

type GroupApi struct {
	app group.Application
}

func NewGroupApi(app group.Application, router fiber.Router) {
	api := &GroupApi{app: app}

	groupRouter := router.Group("/group")
	groupRouter.Get("", api.List)
	groupRouter.Post("", api.Create)
	groupRouter.Post("/sort", api.Sort)
	groupRouter.Put("", api.Update)
	groupRouter.Delete("/:gid", api.Delete)
}

// Create 新增短链接分组
func (h GroupApi) Create(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	reqParam := req.LinkGroupSaveReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		return err
	}
	cmd := command.CreateGroupCommand{
		GroupName: reqParam.Name,
		Username:  username,
	}
	if err := h.app.Commands.CreateGroup.Handle(c.Context(), cmd); err != nil {
		return err
	}
	return nil
}

// List 查询短链接分组集合
func (h GroupApi) List(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	response := resp.LinkGroupResp{}
	res, err := h.app.Queries.ListGroup.Handle(c.Context(), username)
	if err != nil {
		return err
	}
	if err = copier.Copy(&response, &res); err != nil {
		return err
	}
	return c.JSON(response)
}

// Update 修改短链接分组名称
func (h GroupApi) Update(c *fiber.Ctx) error {
	reqParam := req.LinkGroupUpdateReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		return err
	}
	cmd := command.UpdateGroupCommand{
		Gid:  reqParam.Gid,
		Name: reqParam.Name,
	}
	if err := h.app.Commands.UpdateGroup.Handle(c.Context(), cmd); err != nil {
		return err
	}
	return nil
}

// Delete 删除短链接分组
func (h GroupApi) Delete(c *fiber.Ctx) error {
	gid := c.Params("gid")
	if err := h.app.Commands.DeleteGroup.Handle(c.Context(), gid); err != nil {
		return err
	}
	return nil
}

// Sort 排序短链接分组
func (h GroupApi) Sort(c *fiber.Ctx) error {
	reqParam := req.LinkGroupSortReq{}
	if err := c.BodyParser(&reqParam); err != nil {
		return err
	}
	cmd := command.SortGroupCommand{}
	for _, entry := range reqParam {
		cmd = append(cmd, command.GroupSortEntry{
			Gid:       entry.Gid,
			SortOrder: entry.SortOrder,
		})
	}
	if err := h.app.Commands.SortGroup.Handle(c.Context(), cmd); err != nil {
		return err
	}
	return nil
}
