package product

type Domain struct {
	resource ResourceItf
}

type Resource struct {
}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
	Rating      string `json:"rating"`
	Price       string `json:"price"`
	StoreName   string `json:"store_name"`
}

type ProductBuilder interface {
	SetName(name string) ProductBuilder
	SetDescription(description string) ProductBuilder
	SetImageLink(imageLink string) ProductBuilder
	SetRating(rating string) ProductBuilder
	SetPrice(price string) ProductBuilder
	SetStoreName(storeName string) ProductBuilder
	Build() Product
}

type ConcreteProductBuilder struct {
	product Product
}

func (b *ConcreteProductBuilder) SetName(name string) ProductBuilder {
	b.product.Name = name
	return b
}

func (b *ConcreteProductBuilder) SetDescription(description string) ProductBuilder {
	b.product.Description = description
	return b
}

func (b *ConcreteProductBuilder) SetImageLink(imageLink string) ProductBuilder {
	b.product.ImageLink = imageLink
	return b
}

func (b *ConcreteProductBuilder) SetRating(rating string) ProductBuilder {
	b.product.Rating = rating
	return b
}

func (b *ConcreteProductBuilder) SetPrice(price string) ProductBuilder {
	b.product.Price = price
	return b
}

func (b *ConcreteProductBuilder) SetStoreName(storeName string) ProductBuilder {
	b.product.StoreName = storeName
	return b
}

func (b *ConcreteProductBuilder) Build() Product {
	return b.product
}
