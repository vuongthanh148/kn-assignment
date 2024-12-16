package dto

import (
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	modelv1 "github.com/centraldigital/cfw-sales-x-ordering-api/pkg/model/v1"
)

type Customer modelv1.Customer

type GetCustomersByStaffIdRequest modelv1.GetCustomersByStaffIdRequest

func (d GetCustomersByStaffIdRequest) ToDomain() domain.GetCustomersByStaffIdRequest {
	return domain.GetCustomersByStaffIdRequest(d)
}

type GetCustomersByStaffIdResponse modelv1.GetCustomersByStaffIdResponse

func (GetCustomersByStaffIdResponse) FromDomain(d *domain.GetCustomersByStaffIdResponse) *GetCustomersByStaffIdResponse {
	customers := make([]modelv1.Customer, len(d.Customers))
	for i, customerDomain := range d.Customers {
		customers[i] = modelv1.Customer(customerDomain)
	}

	return &GetCustomersByStaffIdResponse{
		Customers: customers,
	}
}

type GetCustomerRequest modelv1.GetCustomerRequest

func (d GetCustomerRequest) ToDomain() domain.GetCustomerRequest {
	return domain.GetCustomerRequest(d)
}

type GetCustomerResponse modelv1.GetCustomerResponse

func (GetCustomerResponse) FromDomain(d *domain.GetCustomerResponse) *GetCustomerResponse {
	deliveryAddresses := make([]modelv1.Addresses, len(d.DeliveryAddresses))
	for i, deliveryAddress := range d.DeliveryAddresses {
		deliveryAddresses[i] = modelv1.Addresses(deliveryAddress)
	}

	return &GetCustomerResponse{
		GeneralInformation:    modelv1.GeneralInformation(d.GeneralInformation),
		DeliveryAddresses:     deliveryAddresses,
		BillingAddress:        (*modelv1.Addresses)(d.BillingAddress),
		Segment:               modelv1.Segment(d.Segment),
		The1:                  modelv1.The1(d.The1),
		Legal:                 modelv1.Legal(d.Legal),
		IsActive:              d.IsActive,
		LastUpdatedAt:         d.LastUpdatedAt,
		UpdatedBy:             d.UpdatedBy,
		AlcoholExpirationDate: d.AlcoholExpirationDate,
	}
}
