package product

const (
	QueryInsertProduct = `
		INSERT INTO products (name, description, image_link, rating, price, store_name)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
)
