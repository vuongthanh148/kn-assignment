package infrastructure

import (
	"context"
	"log"

	"github.com/centraldigital/cfw-core-lib/pkg/adaptor"
	"github.com/centraldigital/cfw-sales-x-ordering-api/property"
	"google.golang.org/api/idtoken"
)

// sample
func NewApiHttpClient(ctx context.Context) adaptor.ClientAdaptor {
	property := property.Get()
	_ = property

	hc, err := idtoken.NewClient(ctx, "host from property")
	if err != nil {
		log.Fatalf("unable to initial http-client: %+v", err)
	}

	cb := adaptor.NewDefaultClientBreaker(hc, "client-name")
	return adaptor.NewClient(cb)
}
