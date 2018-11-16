package products

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Product struct {
	ID       string `db:"product_id" json:"id"`
	Name     string `db:"name" json:"name"`
	Cost     int    `db:"cost" json:"cost"`
	Quantity int    `db:"quantity" json:"quantity"`
}

func List(ctx context.Context, db *sqlx.DB) ([]Product, error) {
	// TODO: Catch "null" in tests with empty list.
	var products []Product

	// TODO: Talk about issues of using '*' after talking about migrations.
	if err := db.SelectContext(ctx, &products, "SELECT * FROM products"); err != nil {
		return nil, errors.Wrap(err, "selecting products")
	}

	return products, nil
}

func Create(ctx context.Context, db *sqlx.DB, p *Product) error {
	p.ID = uuid.New().String()

	_, err := db.ExecContext(ctx, `
		INSERT INTO products
		(product_id, name, cost, quantity)
		VALUES ($1, $2, $3, $4)`,
		p.ID, p.Name, p.Cost, p.Quantity,
	)
	if err != nil {
		return errors.Wrap(err, "inserting product")
	}

	return nil
}

func Get(ctx context.Context, db *sqlx.DB, id string) (*Product, error) {
	var p Product

	err := db.GetContext(ctx, &p, `
		SELECT * FROM products
		WHERE product_id = $1
		LIMIT 1`,
		id,
	)
	if err != nil {
		return nil, errors.Wrap(err, "selecting single product")
	}

	return &p, nil
}
