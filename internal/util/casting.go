package util

type domainifiable[domain any] interface {
	ToDomain() domain
}

func CastListToDomain[domain any, entity domainifiable[domain]](items []entity) []domain {
	domainItems := make([]domain, len(items))
	for i, item := range items {
		domainItems[i] = item.ToDomain()
	}

	return domainItems
}

type undomainifiable[domain any, entity any] interface {
	FromDomain(domain) entity
}

func CastListFromDomain[v undomainifiable[domain, v], domain any](domainItems []domain) []v {
	items := make([]v, len(domainItems))
	for i, domainItem := range domainItems {
		items[i] = items[i].FromDomain(domainItem)
	}

	return items
}
