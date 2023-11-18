package product

func InitDomain(rsc ResourceItf) *Domain {
	return &Domain{
		resource: rsc,
	}
}

func InitResource() Resource {
	return Resource{}
}
