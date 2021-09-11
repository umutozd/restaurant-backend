package storage

import (
	"context"
	"testing"

	"github.com/umutozd/restaurant-backend/types"
	"github.com/umutozd/restaurant-backend/types/requests"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestListMenu(t *testing.T) {
	var err error
	s := newTestStorage(t)

	categories := []*types.Category{
		{
			ID:   primitive.NewObjectID().Hex(),
			Name: "Category 1",
		},
		{
			ID:   primitive.NewObjectID().Hex(),
			Name: "Category 2",
		},
	}
	for _, c := range categories {
		_, err = s.categories().InsertOne(context.TODO(), c)
		if err != nil {
			t.Fatalf("error inserting category %s: %v", c.Name, err)
		}
	}

	items := []*types.MenuItem{
		{
			CategoryID: categories[0].ID,
			ID:         primitive.NewObjectID().Hex(),
			Name:       "Menu Item 1",
			Price:      21,
			Image:      "Image 1",
		},
		{
			CategoryID: categories[0].ID,
			ID:         primitive.NewObjectID().Hex(),
			Name:       "Menu Item 2",
			Price:      23,
			Image:      "Image 2",
		},
		{
			CategoryID: categories[1].ID,
			ID:         primitive.NewObjectID().Hex(),
			Name:       "Menu Item 3",
			Price:      8,
			Image:      "Image 3",
		},
		{
			CategoryID: categories[1].ID,
			ID:         primitive.NewObjectID().Hex(),
			Name:       "Menu Item 4",
			Price:      9.50,
			Image:      "Image 4",
		},
		{
			CategoryID: categories[1].ID,
			ID:         primitive.NewObjectID().Hex(),
			Name:       "Menu Item 5",
			Price:      6.75,
			Image:      "Image 5",
		},
	}
	for _, i := range items {
		_, err = s.menuItems().InsertOne(context.TODO(), i)
		if err != nil {
			t.Fatalf("error inserting menu item %s: %v", i.Name, err)
		}
	}

	expectedMenu := &types.Menu{All: []*types.CategoryAndItems{
		{
			Category: categories[0],
			Items: []*types.MenuItem{
				items[0],
				items[1],
			},
		},
		{
			Category: categories[1],
			Items: []*types.MenuItem{
				items[2],
				items[3],
				items[4],
			},
		},
	}}
	result, err := s.ListMenu(context.TODO())
	if err != nil {
		t.Fatalf("error fetching menu: %v", err)
	}

	same := compareMenus(expectedMenu, result)
	if !same {
		t.Fatalf("expected menu and the result are not the same: \nexpected: %v \nresult: %v", expectedMenu, result)
	}
}

func TestCreateMenuItem(t *testing.T) {
	var err error
	s := newTestStorage(t)

	cases := []*struct {
		item    *types.MenuItem
		created *types.MenuItem
	}{
		{
			item: &types.MenuItem{
				Name:       "Name1",
				CategoryID: "123456",
				Price:      8.99,
				Image:      "image 1",
			},
		},
		{
			item: &types.MenuItem{
				Name:       "Name2",
				CategoryID: "123456",
				Price:      3.99,
				Image:      "image 2",
			},
		},
		{
			item: &types.MenuItem{
				Name:       "Name3",
				CategoryID: "654321",
				Price:      3.50,
				Image:      "image 3",
			},
		},
		{
			item: &types.MenuItem{
				Name:       "Name4",
				CategoryID: "654321",
				Price:      4.50,
				Image:      "image 4",
			},
		},
		{
			item: &types.MenuItem{
				Name:       "Name5",
				CategoryID: "101010",
				Price:      10,
				Image:      "image 5",
			},
		},
	}

	for i, c := range cases {
		c.created, err = s.CreateMenuItem(context.TODO(), c.item)
		if err != nil {
			t.Fatalf("case %d; error creating item: %v", i, err)
		}
	}

	// now, check the created items are actually inserted
	for i, c := range cases {
		var item types.MenuItem
		err = s.menuItems().
			FindOne(context.TODO(), bson.M{"_id": c.created.ID}).
			Decode(&item)
		if err != nil {
			t.Fatalf("case %d; error getting item with id=%s: %v", i, c.created.ID, err)
		}
		if diff := compareMenuItems(c.created, &item); diff != "" {
			t.Fatalf("case %d; retrieved item is not the same as the created one, diff: %s", i, diff)
		}
	}
}

func TestListMenuItems(t *testing.T) {
	var err error
	s := newTestStorage(t)

	cases := []*struct {
		item *types.MenuItem
	}{
		{
			item: &types.MenuItem{
				ID:         primitive.NewObjectID().Hex(),
				Name:       "Name1",
				CategoryID: "123456",
				Price:      8.99,
				Image:      "image 1",
			},
		},
		{
			item: &types.MenuItem{
				ID:         primitive.NewObjectID().Hex(),
				Name:       "Name2",
				CategoryID: "123456",
				Price:      3.99,
				Image:      "image 2",
			},
		},
		{
			item: &types.MenuItem{
				ID:         primitive.NewObjectID().Hex(),
				Name:       "Name3",
				CategoryID: "654321",
				Price:      3.50,
				Image:      "image 3",
			},
		},
		{
			item: &types.MenuItem{
				ID:         primitive.NewObjectID().Hex(),
				Name:       "Name4",
				CategoryID: "654321",
				Price:      4.50,
				Image:      "image 4",
			},
		},
		{
			item: &types.MenuItem{
				ID:         primitive.NewObjectID().Hex(),
				Name:       "Name5",
				CategoryID: "101010",
				Price:      10,
				Image:      "image 5",
			},
		},
	}

	for i, c := range cases {
		_, err = s.menuItems().InsertOne(context.TODO(), c.item)
		if err != nil {
			t.Fatalf("case %d; error inserting test item: %v", i, err)
		}
	}

	// now, list all items
	result, err := s.ListMenuItems(context.TODO())
	if err != nil {
		t.Fatalf("error listing items: %v", err)
	}
	if len(result) != len(cases) {
		t.Fatalf("document count in result does not match number of cases: result=%d, cases=%d", len(result), len(cases))
	}

	// for each case, check if result containts it
	for i, c := range cases {
		diff := "not found yet"
		for _, resItem := range result {
			diff = compareMenuItems(c.item, resItem)
			if diff == "" {
				break
			}
		}
		if diff != "" {
			t.Fatalf("case %d; unable to find case in result", i)
		}
	}
}

func TestUpdateMenuItem(t *testing.T) {
	s := newTestStorage(t)

	ids := []string{
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
		primitive.NewObjectID().Hex(),
	}

	cases := []*struct {
		item     *types.MenuItem
		req      *requests.UpdateMenuItemReq
		expected *types.MenuItem
	}{
		{
			item: &types.MenuItem{
				ID:         ids[0],
				Name:       "Name1",
				CategoryID: "123456",
				Price:      8.99,
				Image:      "image 1",
			},
			req: &requests.UpdateMenuItemReq{
				Fields: []requests.UpdateMenuItemField{
					requests.CATEGORY_ID,
				},
				Item: &types.MenuItem{
					ID:         ids[0],
					CategoryID: "111222",
				},
			},
			expected: &types.MenuItem{
				ID:         ids[0],
				Name:       "Name1",
				CategoryID: "111222",
				Price:      8.99,
				Image:      "image 1",
			},
		},
		{
			item: &types.MenuItem{
				ID:         ids[1],
				Name:       "Name2",
				CategoryID: "123456",
				Price:      3.99,
				Image:      "image 2",
			},
			req: &requests.UpdateMenuItemReq{
				Fields: []requests.UpdateMenuItemField{
					requests.IMAGE,
				},
				Item: &types.MenuItem{
					ID:    ids[1],
					Image: "image 2 updated",
				},
			},
			expected: &types.MenuItem{
				ID:         ids[1],
				Name:       "Name2",
				CategoryID: "123456",
				Price:      3.99,
				Image:      "image 2 updated",
			},
		},
		{
			item: &types.MenuItem{
				ID:         ids[2],
				Name:       "Name3",
				CategoryID: "654321",
				Price:      3.50,
				Image:      "image 3",
			},
			req: &requests.UpdateMenuItemReq{
				Fields: []requests.UpdateMenuItemField{
					requests.ITEM_NAME,
				},
				Item: &types.MenuItem{
					ID:   ids[2],
					Name: "Name3 updated",
				},
			},
			expected: &types.MenuItem{
				ID:         ids[2],
				Name:       "Name3 updated",
				CategoryID: "654321",
				Price:      3.50,
				Image:      "image 3",
			},
		},
		{
			item: &types.MenuItem{
				ID:         ids[3],
				Name:       "Name4",
				CategoryID: "654321",
				Price:      4.50,
				Image:      "image 4",
			},
			req: &requests.UpdateMenuItemReq{
				Fields: []requests.UpdateMenuItemField{
					requests.PRICE,
				},
				Item: &types.MenuItem{
					ID:    ids[3],
					Price: 10.25,
				},
			},
			expected: &types.MenuItem{
				ID:         ids[3],
				Name:       "Name4",
				CategoryID: "654321",
				Price:      10.25,
				Image:      "image 4",
			},
		},
	}

	for i, c := range cases {
		_, err := s.menuItems().InsertOne(context.TODO(), c.item)
		if err != nil {
			t.Fatalf("case %d; error inserting test item: %v", i, err)
		}
	}

	for i, c := range cases {
		updated, err := s.UpdateMenuItem(context.TODO(), c.req)
		if err != nil {
			t.Fatalf("case %d; error updating item: %v", i, err)
		}

		if diff := compareMenuItems(c.expected, updated); diff != "" {
			t.Fatalf("case %d; wrong result after update, diff: %s", i, diff)
		}
	}
}

func TestDeleteMenuItem(t *testing.T) {
	s := newTestStorage(t)

	cases := []*types.MenuItem{
		{
			ID:   primitive.NewObjectID().Hex(),
			Name: "Name1",
		},
		{
			ID:   primitive.NewObjectID().Hex(),
			Name: "Name2",
		},
		{
			ID:   primitive.NewObjectID().Hex(),
			Name: "Name3",
		},
		{
			ID:   primitive.NewObjectID().Hex(),
			Name: "Name4",
		},
		{
			ID:   primitive.NewObjectID().Hex(),
			Name: "Name5",
		},
	}

	// insert test cases
	for i, c := range cases {
		_, err := s.menuItems().InsertOne(context.TODO(), c)
		if err != nil {
			t.Fatalf("case %d; error inserting test item: %v", i, err)
		}
	}

	// delete test cases
	for i, c := range cases {
		err := s.DeleteMenuItem(context.TODO(), &requests.DeleteMenuItemReq{ID: c.ID})
		if err != nil {
			t.Fatalf("case %d; error deleting item: %v", i, err)
		}
	}

	// check if deleted items are actually deleted by attempting to fetch them
	for i, c := range cases {
		err := s.menuItems().FindOne(context.TODO(), bson.M{"_id": c.ID}).Err()
		if err != nil {
			if err == mongo.ErrNoDocuments {
				continue
			} else {
				t.Fatalf("case %d; an unexpected error occured when fetching deleted document: %v", i, err)
			}
		} else {
			t.Fatalf("case %d; item is not deleted!", i)
		}
	}
}
