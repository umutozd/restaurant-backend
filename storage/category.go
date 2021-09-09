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

func (s *storage) CreateCategory(ctx context.Context, category *types.Category) (*types.Category, error) {
	category.ID = primitive.NewObjectID().Hex()
	_, err := s.categories().InsertOne(ctx, category)
	if err != nil {
		return nil, types.Err(types.ERR_DB_INSERT, err)
	}
	return category, nil
}

func (s *storage) ListCategories(ctx context.Context) ([]*types.Category, error) {
	cursor, err := s.categories().Find(ctx, bson.M{})
	if err != nil {
		return nil, types.Err(types.ERR_DB_LIST, err)
	}
	var categories []*types.Category
	if err = cursor.All(ctx, &categories); err != nil {
		return nil, types.Err(types.ERR_DB_DECODE, err)
	}

	return categories, nil
}

func (s *storage) UpdateCategory(ctx context.Context, req *requests.UpdateCategoryReq) (*types.Category, error) {
	filter := bson.M{"_id": req.Category.ID}
	set := bson.M{}
	for _, f := range req.Fields {
		switch f {
		case requests.CAGETORY_NAME:
			set["name"] = req.Category.Name
		default:
			continue
		}
	}
	update := bson.M{"$set": set}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	res := s.categories().FindOneAndUpdate(ctx, filter, update, opts)
	category := &types.Category{}
	if err := res.Decode(category); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, types.Errf(types.ERR_NOT_FOUND, "category with id=%s not found", req.Category.ID)
		}
		return nil, types.Err(types.ERR_DB_UPDATE, err)
	}

	return category, nil
}

func (s *storage) DeleteCategory(ctx context.Context, req *requests.DeleteCategoryReq) error {
	filter := bson.M{"_id": req.ID}
	res, err := s.categories().DeleteOne(ctx, filter)
	if err != nil {
		return types.Err(types.ERR_DB_DELETE, err)
	}
	if res.DeletedCount != 1 {
		return types.Errf(types.ERR_NOT_FOUND, "deleted count is %d, not 1", res.DeletedCount)
	}

	return nil
}
