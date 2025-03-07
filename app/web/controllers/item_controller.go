package controllers

import (
	"github.com/berthojoris/cart-backend/app/models"
	"github.com/berthojoris/cart-backend/app/services/interface"
	"github.com/berthojoris/cart-backend/app/web/response"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
)

type ItemController struct {
	Db *gorm.DB
	ItemService _interface.IItemService
}

func NewItemController(db *gorm.DB, itemService _interface.IItemService) *ItemController {
	return &ItemController{
		Db: db,
		ItemService: itemService,
	}
}

func (c *ItemController) GetIndexHandler(ctx iris.Context) {
	var items []models.Item

	c.ItemService.GetAll(c.Db, &items)

	if len(items) == 0 {
		response.SuccessResponse(ctx, response.OK, response.OK_MESSAGE, make([]interface{}, 0))
		return
	}

	itemResponse := response.NewItemResponse(c.Db)
	result := itemResponse.Collection(items)

	response.SuccessResponse(ctx, response.OK, response.OK_MESSAGE, result)
}

func (c *ItemController) GetDetailHandler(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	var detailItem models.Item
	
	c.ItemService.GetById(c.Db, &detailItem, int(id))

	if detailItem == (models.Item{}) {
		response.ErrorResponse(ctx, response.UNPROCESSABLE_ENTITY, "Item doesn't exists.")
		return
	}

	detailItemResponse := response.NewItemResponse(c.Db)
	result := detailItemResponse.New(detailItem)

	response.SuccessResponse(ctx, response.OK, response.OK_MESSAGE, result)
}
