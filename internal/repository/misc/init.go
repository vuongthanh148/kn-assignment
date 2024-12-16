package miscrepo

import "github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/port"

type miscRepo struct{}

func New() port.MiscRepo {
	return &miscRepo{}
}
