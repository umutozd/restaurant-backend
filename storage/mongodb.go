package storage

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/umutozd/restaurant-backend/types"
	"github.com/umutozd/restaurant-backend/types/requests"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	menuItemsCollection  = "MenuItems"
	categoriesCollection = "Categories"
	cartsCollection      = "Carts"
)

type Storage interface {
	ListMenu(ctx context.Context) (*types.Menu, error)
	CreateMenuItem(ctx context.Context, item *types.MenuItem) (*types.MenuItem, error)
	ListMenuItems(ctx context.Context) ([]*types.MenuItem, error)
	UpdateMenuItem(ctx context.Context, req *requests.UpdateMenuItemReq) (*types.MenuItem, error)
	DeleteMenuItem(ctx context.Context, req *requests.DeleteMenuItemReq) error

	CreateCategory(ctx context.Context, category *types.Category) (*types.Category, error)
	ListCategories(ctx context.Context) ([]*types.Category, error)
	UpdateCategory(ctx context.Context, req *requests.UpdateCategoryReq) (*types.Category, error)
	DeleteCategory(ctx context.Context, req *requests.DeleteCategoryReq) error

	UpdateCart(ctx context.Context, req *requests.UpdateCartReq) (*types.Cart, error)
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

func (s *storage) carts() *mongo.Collection {
	return s.client.Database(s.dbName).Collection(cartsCollection)
}

func NewStorage(dbURL, dbName string) (Storage, error) {
	logrus.Debug("Connecting to mongodb")
	client, err := mongo.NewClient(options.Client().ApplyURI(dbURL))
	if err != nil {
		logrus.WithError(err).Errorf("error creating mongo client")
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	err = client.Connect(ctx)
	cancel()
	if err != nil {
		logrus.WithError(err).Errorf("error connecting to mongo server")
		return nil, err
	}
	logrus.Debug("Connected to mongodb")
	return &storage{client: client, dbName: dbName}, nil
}
