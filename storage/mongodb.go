package storage

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/umutozd/restaurant-backend/types"
	"github.com/umutozd/restaurant-backend/types/requests"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	menuItemsCollection  = "MenuItems"
	categoriesCollection = "Categories"
)

type Storage interface {
	ListMenu(ctx context.Context,) (*types.Menu, error)
	CreateMenuItem(ctx context.Context,item *types.MenuItem) (*types.MenuItem, error)
	ListMenuItems(ctx context.Context,) ([]*types.MenuItem, error)
	UpdateMenuItem(ctx context.Context,req *requests.UpdateMenuItemReq) (*types.MenuItem, error)
	DeleteMenuItem(ctx context.Context,req *requests.DeleteMenuItemReq) error

	CreateCategory(ctx context.Context,category *types.Category) (*types.Category, error)
	ListCategories(ctx context.Context,) ([]*types.Category, error)
	UpdateCategory(ctx context.Context,req *requests.UpdateCategoryReq) (*types.Category, error)
	DeleteCategory(ctx context.Context,req *requests.DeleteCategoryReq) error
}

type storage struct {
	client *mongo.Client
	dbName string
}

func (s *storage) menuItems() *mongo.Collection {
	return s.client.Database(s.dbName).Collection(menuItemsCollection)
}

func (s *storage) categories() *mongo.Collection {
	return s.client.Database(s.dbName).Collection(categoriesCollection)
}

func NewStorage(dbURL, dbName string) (Storage, error) {
	logrus.Debug("Connecting to mongodb")
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURL))
	if err != nil {
		logrus.WithError(err).Errorf("could not create mongo client")
		return nil, err
	}
	err = client.Connect(context.Background())
	if err != nil {
		logrus.WithError(err).Errorf("could not connect to mongo client")
		return nil, err
	}
	return &storage{client: client, dbName: dbName}, nil
}
