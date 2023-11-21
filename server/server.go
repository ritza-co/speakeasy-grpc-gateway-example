package server

import (
	"context"
	"fmt"

	speakeasyv1 "github.com/speakeasy-api/speakeasy-grpc-gateway-example/proto/speakeasy/v1"
)

type Backend struct {
	Drinks []*speakeasyv1.Drink
}

func New() *Backend {
	return &Backend{
		Drinks: drinks,
	}
}

var (
	drinks = []*speakeasyv1.Drink{
		{
			Name:  "Negroni",
			Type:  speakeasyv1.Drink_DRINK_TYPE_COCKTAIL,
			Price: 12.99,
		},
		{
			Name:  "Old Fashioned",
			Type:  speakeasyv1.Drink_DRINK_TYPE_COCKTAIL,
			Price: 14.99,
		},
		{
			Name:  "La Chouffe",
			Type:  speakeasyv1.Drink_DRINK_TYPE_BEER,
			Price: 4.99,
		},
	}
)

func (b *Backend) ListDrinks(ctx context.Context, in *speakeasyv1.ListDrinksRequest) (*speakeasyv1.ListDrinksResponse, error) {
	return &speakeasyv1.ListDrinksResponse{Drinks: b.Drinks}, nil
}

func (b *Backend) GetDrink(ctx context.Context, in *speakeasyv1.GetDrinkRequest) (*speakeasyv1.GetDrinkResponse, error) {
	for _, drink := range drinks {
		if drink.ProductCode == in.ProductCode {
			return &speakeasyv1.GetDrinkResponse{Drink: drink}, nil
		}
	}
	return nil, fmt.Errorf("drink %s not found", in.ProductCode)
}
