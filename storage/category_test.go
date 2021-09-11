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

func TestCreateCategory(t *testing.T) {
	var err error
	s := newTestStorage(t)

	cases := []*struct {
		category *types.Category
		created  *types.Category
	}{
		{
			category: &types.Category{
				Name: "Name1",
			},
		},
		{
			category: &types.Category{
				Name: "Name2",
			},
		},
		{
			category: &types.Category{
				Name: "Name3",
			},
		},
		{
			category: &types.Category{
				Name: "Name4",
			},
		},
		{
			category: &types.Category{
				Name: "Name5",
			},
		},
	}

	for i, c := range cases {
		c.created, err = s.CreateCategory(context.TODO(), c.category)
		if err != nil {
			t.Fatalf("case %d; error creating category: %v", i, err)
		}
	}

	// now, check the created categories are actually inserted
	for i, c := range cases {
		var category types.Category
		err = s.categories().
			FindOne(context.TODO(), bson.M{"_id": c.created.ID}).
			Decode(&category)
		if err != nil {
			t.Fatalf("case %d; error getting category with id=%s: %v", i, c.created.ID, err)
		}
		if diff := compareCategories(c.created, &category); diff != "" {
			t.Fatalf("case %d; retrieved category is not the same as the created one, diff: %s", i, diff)
		}
	}
}

func TestListCategories(t *testing.T) {
	var err error
	s := newTestStorage(t)

	cases := []*struct {
		category *types.Category
	}{
		{
			category: &types.Category{
				ID:   primitive.NewObjectID().Hex(),
				Name: "Name1",
			},
		},
		{
			category: &types.Category{
				ID:   primitive.NewObjectID().Hex(),
				Name: "Name2",
			},
		},
		{
			category: &types.Category{
				ID:   primitive.NewObjectID().Hex(),
				Name: "Name3",
			},
		},
		{
			category: &types.Category{
				ID:   primitive.NewObjectID().Hex(),
				Name: "Name4",
			},
		},
		{
			category: &types.Category{
				ID:   primitive.NewObjectID().Hex(),
				Name: "Name5",
			},
		},
	}

	for i, c := range cases {
		_, err = s.categories().InsertOne(context.TODO(), c.category)
		if err != nil {
			t.Fatalf("case %d; error inserting test category: %v", i, err)
		}
	}

	// now, list all categories
	result, err := s.ListCategories(context.TODO())
	if err != nil {
		t.Fatalf("error listing categories: %v", err)
	}
	if len(result) != len(cases) {
		t.Fatalf("document count in result does not match number of cases: result=%d, cases=%d", len(result), len(cases))
	}

	// for each case, check if result containts it
	for i, c := range cases {
		diff := "not found yet"
		for _, resItem := range result {
			diff = compareCategories(c.category, resItem)
			if diff == "" {
				break
			}
		}
		if diff != "" {
			t.Fatalf("case %d; unable to find case in result", i)
		}
	}
}

func TestUpdateCategory(t *testing.T) {
	s := newTestStorage(t)

	ids := []string{
		primitive.NewObjectID().Hex(),
	}

	cases := []*struct {
		category *types.Category
		req      *requests.UpdateCategoryReq
		expected *types.Category
	}{
		{
			category: &types.Category{
				ID:   ids[0],
				Name: "Name1",
			},
			req: &requests.UpdateCategoryReq{
				Fields: []requests.UpdateCategoryField{
					requests.CAGETORY_NAME,
				},
				Category: &types.Category{
					ID:   ids[0],
					Name: "New Name1",
				},
			},
			expected: &types.Category{
				ID:   ids[0],
				Name: "New Name1",
			},
		},
	}

	for i, c := range cases {
		_, err := s.categories().InsertOne(context.TODO(), c.category)
		if err != nil {
			t.Fatalf("case %d; error inserting test category: %v", i, err)
		}
	}

	for i, c := range cases {
		updated, err := s.UpdateCategory(context.TODO(), c.req)
		if err != nil {
			t.Fatalf("case %d; error updating category: %v", i, err)
		}

		if diff := compareCategories(c.expected, updated); diff != "" {
			t.Fatalf("case %d; wrong result after update, diff: %s", i, diff)
		}
	}
}

func TestDeleteCategory(t *testing.T) {
	s := newTestStorage(t)

	cases := []*types.Category{
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
		_, err := s.categories().InsertOne(context.TODO(), c)
		if err != nil {
			t.Fatalf("case %d; error inserting test category: %v", i, err)
		}
	}

	// delete test cases
	for i, c := range cases {
		err := s.DeleteCategory(context.TODO(), &requests.DeleteCategoryReq{ID: c.ID})
		if err != nil {
			t.Fatalf("case %d; error deleting category: %v", i, err)
		}
	}

	// check if deleted categories are actually deleted by attempting to fetch them
	for i, c := range cases {
		err := s.categories().FindOne(context.TODO(), bson.M{"_id": c.ID}).Err()
		if err != nil {
			if err == mongo.ErrNoDocuments {
				continue
			} else {
				t.Fatalf("case %d; an unexpected error occured when fetching deleted document: %v", i, err)
			}
		} else {
			t.Fatalf("case %d; category is not deleted!", i)
		}
	}
}
