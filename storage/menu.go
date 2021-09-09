package storage

import (
	"context"

	"github.com/umutozd/restaurant-backend/types"
	"github.com/umutozd/restaurant-backend/types/requests"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *storage) ListMenu(ctx context.Context) (*types.Menu, error) {
	items, err := s.ListMenuItems(ctx)
	if err != nil {
		return nil, err
	}
	categories, err := s.ListCategories(ctx)
	if err != nil {
		return nil, err
	}

	menu := &types.Menu{
		All: make([]*types.CategoryAndItems, 0, len(categories)),
	}
	for _, c := range categories {
		categoryAndItems := &types.CategoryAndItems{
			Category: c,
		}
		for _, item := range items {
			if item.CategoryID == c.ID {
				categoryAndItems.Items = append(categoryAndItems.Items, item)
			}
		}
		menu.All = append(menu.All, categoryAndItems)
	}

	return menu, nil
}

func (s *storage) CreateMenuItem(ctx context.Context, item *types.MenuItem) (*types.MenuItem, error) {
	item.ID = primitive.NewObjectID().Hex()
	_, err := s.menuItems().InsertOne(ctx, item)
	if err != nil {
		return nil, types.Err(types.ERR_DB_INSERT, err)
	}
	return item, nil
}

func (s *storage) ListMenuItems(ctx context.Context) ([]*types.MenuItem, error) {
	cursor, err := s.menuItems().Find(ctx, bson.M{})
	if err != nil {
		return nil, types.Err(types.ERR_DB_LIST, err)
	}
	var items []*types.MenuItem
	if err = cursor.All(ctx, &items); err != nil {
		return nil, types.Err(types.ERR_DB_DECODE, err)
	}

	return items, nil
}

func (s *storage) UpdateMenuItem(ctx context.Context, req *requests.UpdateMenuItemReq) (*types.MenuItem, error) {
	filter := bson.M{"_id": req.Item.ID}
	set := bson.M{}
	for _, f := range req.Fields {
		switch f {
		case requests.CATEGORY_ID:
			set["category_id"] = req.Item.CategoryID
		case requests.ITEM_NAME:
			set["name"] = req.Item.Name
		case requests.PRICE:
			set["price"] = req.Item.Price
		case requests.IMAGE:
			set["image"] = req.Item.Image
		default:
			continue
		}
	}
	update := bson.M{"$set": set}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	res := s.menuItems().FindOneAndUpdate(ctx, filter, update, opts)
	item := &types.MenuItem{}
	if err := res.Decode(item); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, types.Errf(types.ERR_NOT_FOUND, "menu item with id=%s not found", req.Item.ID)
		}
		return nil, types.Err(types.ERR_DB_UPDATE, err)
	}

	return item, nil
}

func (s *storage) DeleteMenuItem(ctx context.Context, req *requests.DeleteMenuItemReq) error {
	filter := bson.M{"_id": req.ID}
	res, err := s.menuItems().DeleteOne(ctx, filter)
	if err != nil {
		return types.Err(types.ERR_DB_DELETE, err)
	}
	if res.DeletedCount != 1 {
		return types.Errf(types.ERR_NOT_FOUND, "deleted count is %d, not 1", res.DeletedCount)
	}

	return nil
}
