package service

import (
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/port"
)

type service struct {
	miscRepository port.MiscRepo
	repository     port.Repository

	coreCustomerAdapter port.CoreCustomerAdapter
	coreProductAdapter  port.CoreProductAdapter
	coreOrderAdapter    port.CoreOrderAdapter
	coreStoreAdapter    port.CoreStoreAdapter
}

func New(
	miscRepository port.MiscRepo,
	repository port.Repository,
	coreCustomerAdapter port.CoreCustomerAdapter,
	coreProductAdapter port.CoreProductAdapter,
	coreOrderAdapter port.CoreOrderAdapter,
	coreStoreAdapter port.CoreStoreAdapter,
) port.Service {
	return &service{
		miscRepository: miscRepository,
		repository:     repository,

		coreCustomerAdapter: coreCustomerAdapter,
		coreProductAdapter:  coreProductAdapter,
		coreOrderAdapter:    coreOrderAdapter,
		coreStoreAdapter:    coreStoreAdapter,
	}
}
