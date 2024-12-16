package service_test

import (
	"testing"

	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/port"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/port/mocks"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/service"
)

type testingModule struct {
	service port.Service

	miscRepository *mocks.MiscRepo
	repository     *mocks.Repository

	coreCustomerAdapter *mocks.CoreCustomerAdapter
	coreProductAdapter  *mocks.CoreProductAdapter
	coreOrderAdapter    *mocks.CoreOrderAdapter
	coreStoreAdapter    *mocks.CoreStoreAdapter
}

func newTestingModule(t *testing.T) *testingModule {
	var (
		miscRepository = mocks.NewMiscRepo(t)
		repository     = mocks.NewRepository(t)

		coreCustomerAdapter = mocks.NewCoreCustomerAdapter(t)
		coreProductAdapter  = mocks.NewCoreProductAdapter(t)
		coreOrderAdapter    = mocks.NewCoreOrderAdapter(t)
		coreStoreAdapter    = mocks.NewCoreStoreAdapter(t)

		service = service.New(
			miscRepository,
			repository,
			coreCustomerAdapter,
			coreProductAdapter,
			coreOrderAdapter,
			coreStoreAdapter,
		)
	)

	return &testingModule{
		service:        service,
		miscRepository: miscRepository,
		repository:     repository,

		coreCustomerAdapter: coreCustomerAdapter,
		coreProductAdapter:  coreProductAdapter,
		coreOrderAdapter:    coreOrderAdapter,
		coreStoreAdapter:    coreStoreAdapter,
	}
}
