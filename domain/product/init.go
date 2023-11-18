package product

import "toped-scrapper/pkg/database/postgres"

func InitDomain(rsc ResourceItf) *Domain {
	return &Domain{
		resource: rsc,
	}
}

func InitResource(db postgres.PostgresHandler) Resource {
	return Resource{
		db: db,
	}
}
