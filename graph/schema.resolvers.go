package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"
	"strconv"

	"github.com/lpc0503/Grocery-Tracker/graph/model"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, userID string) (*model.User, error) {
	for _, user := range r.users {
		if user.UserID == userID {
			return nil, errors.New("User exist")
		}
	}

	user := &model.User{
		UserID: userID,
	}
	r.users = append(r.users, user)
	return user, nil
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, userID string) (*model.User, error) {
	for _, user := range r.users {
		if user.UserID == userID {
			return user, nil
		}
	}
	return nil, errors.New("invalid credentials")
}

// AddUserGroceryItem is the resolver for the addUserGroceryItem field.
func (r *mutationResolver) AddUserGroceryItem(ctx context.Context, userID string, name string, quantity *int, purchaseDate *string, expirationDate *string, price *float64, materials []*string, category *string) (*model.GroceryItem, error) {
	item := &model.GroceryItem{
		ID:             strconv.Itoa(len(r.groceryItems[userID]) + 1),
		UserID:         userID,
		Name:           name,
		Quantity:       quantity,
		PurchaseDate:   purchaseDate,
		ExpirationDate: expirationDate,
		Price:          price,
		Materials:      materials,
		Category:       category,
	}

	if _, ok := r.groceryItems[userID]; !ok {
		r.groceryItems[userID] = []*model.GroceryItem{}
	}

	r.groceryItems[userID] = append(r.groceryItems[userID], item)
	return item, nil
}

// UpdateUserGroceryItem is the resolver for the updateUserGroceryItem field.
func (r *mutationResolver) UpdateUserGroceryItem(ctx context.Context, userID string, id string, name *string, quantity *int, purchaseDate *string, expirationDate *string, price *float64, materials []*string, category *string) (*model.GroceryItem, error) {
	for _, item := range r.groceryItems[userID] {
		if item.ID == id {
			if name != nil {
				item.Name = *name
			}
			if quantity != nil {
				item.Quantity = quantity
			}
			if purchaseDate != nil {
				item.PurchaseDate = purchaseDate
			}
			if expirationDate != nil {
				item.ExpirationDate = expirationDate
			}
			if price != nil {
				item.Price = price
			}
			if materials != nil {
				item.Materials = materials
			}
			if category != nil {
				item.Category = category
			}
			return item, nil
		}
	}
	return nil, errors.New("item not found")
}

// DeleteUserGroceryItem is the resolver for the deleteUserGroceryItem field.
func (r *mutationResolver) DeleteUserGroceryItem(ctx context.Context, userID string, id string) (bool, error) {

	for i, item := range r.groceryItems[userID] {
		println(item.ID)
		if item.ID == id {
			println("in")
			r.groceryItems[userID] = append(r.groceryItems[userID][:i], r.groceryItems[userID][i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("item not found")
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	for _, user := range r.users {
		if user.UserID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// GetUserGroceryItem is the resolver for the getUserGroceryItem field.
func (r *queryResolver) GetUserGroceryItem(ctx context.Context, userID string, id string) (*model.GroceryItem, error) {
	for _, item := range r.groceryItems[userID] {

		if item.ID == id {

			return item, nil
		}

	}
	return nil, errors.New("Item not found")
}

// GetUserGroceryItems is the resolver for the getUserGroceryItems field.
func (r *queryResolver) GetUserGroceryItems(ctx context.Context, userID string) ([]*model.GroceryItem, error) {
	//TODO check user status
	return r.groceryItems[userID], nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
