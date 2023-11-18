package ucproduct

import "toped-scrapper/domain/product"

type Usecase struct {
	productDomain product.DomainItf
}

func InitUsecase(productDomain product.DomainItf) *Usecase {
	return &Usecase{
		productDomain: productDomain,
	}
}
