package dto

import (
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	modelv1 "github.com/centraldigital/cfw-sales-x-ordering-api/pkg/model/v1"
)

type Order modelv1.Order

type Item modelv1.Item

type DeliveryAddress modelv1.DeliveryAddress

type Payment modelv1.Payment

type GetOrderHistoryResponse modelv1.GetOrderHistoryResponse

type GetOrderHistoryRequest modelv1.GetOrderHistoryRequest

func (d GetOrderHistoryRequest) ToDomain() domain.GetOrderHistoryRequest {
	return domain.GetOrderHistoryRequest(d)
}

func (d GetOrderHistoryResponse) ToDomain() *domain.GetOrderHistoryResponse {
	orders := make([]domain.Order, len(d.Orders))
	for i, order := range d.Orders {
		orders[i] = toConvertOrderHistory(order)
	}

	return &domain.GetOrderHistoryResponse{
		Orders:     orders,
		Pagination: d.Pagination,
	}
}

func (GetOrderHistoryResponse) FromDomain(d *domain.GetOrderHistoryResponse) *GetOrderHistoryResponse {
	orders := make([]modelv1.Order, len(d.Orders))
	for i, order := range d.Orders {
		orders[i] = fromConvertOrderHistory(order)
	}

	return &GetOrderHistoryResponse{
		Orders:     orders,
		Pagination: d.Pagination,
	}
}

// utility to domain
func toConvertOrderHistory(order modelv1.Order) domain.Order {
	return domain.Order{
		OrderID:                  order.OrderID,
		OrderStatus:              order.OrderStatus,
		OrderDatetime:            order.OrderDatetime,
		OrderCreationDatetime:    order.OrderCreationDatetime,
		OrderUpdatedDatetime:     order.OrderUpdatedDatetime,
		StoreCode:                order.StoreCode,
		CustomerSegment:          order.CustomerSegment,
		SaleChannel:              order.SaleChannel,
		NetAmount:                order.NetAmount,
		TotalDiscount:            order.TotalDiscount,
		TotalMarketingDiscount:   order.TotalMarketingDiscount,
		TotalMerchandizeDiscount: order.TotalMerchandizeDiscount,
		MobileNo:                 order.MobileNo,
		Email:                    order.Email,
		TaxDestination:           order.TaxDestination,
		Items:                    toConvertItems(order.Items),
		DeliveryMethod:           order.DeliveryMethod,
		DeliveryAddress:          toConvertDeliveryAddress(order.DeliveryAddress),
		DeliveryCost:             order.DeliveryCost,
		DeliveryDatetime:         order.DeliveryDatetime,
		TotalAmount:              order.TotalAmount,
		DeliverySlot:             order.DeliverySlot,
		Promotions:               order.Promotions,
		Coupons:                  order.Coupons,
		Payments:                 toConvertPayments(order.Payments),
	}
}

func toConvertDeliveryAddress(address modelv1.DeliveryAddress) domain.DeliveryAddress {
	return domain.DeliveryAddress(address)
}

func toConvertItems(items []modelv1.Item) []domain.Item {
	result := make([]domain.Item, len(items))
	for i, item := range items {
		result[i] = domain.Item(item)
	}
	return result
}

func toConvertPayments(payments []modelv1.Payment) []domain.Payment {
	result := make([]domain.Payment, len(payments))
	for i, payment := range payments {
		result[i] = domain.Payment(payment)
	}
	return result
}

// utility form domain
func fromConvertOrderHistory(order domain.Order) modelv1.Order {
	return modelv1.Order{
		OrderID:                  order.OrderID,
		OrderStatus:              order.OrderStatus,
		OrderDatetime:            order.OrderDatetime,
		OrderCreationDatetime:    order.OrderCreationDatetime,
		OrderUpdatedDatetime:     order.OrderUpdatedDatetime,
		StoreCode:                order.StoreCode,
		CustomerSegment:          order.CustomerSegment,
		SaleChannel:              order.SaleChannel,
		NetAmount:                order.NetAmount,
		TotalDiscount:            order.TotalDiscount,
		TotalMarketingDiscount:   order.TotalMarketingDiscount,
		TotalMerchandizeDiscount: order.TotalMerchandizeDiscount,
		MobileNo:                 order.MobileNo,
		Email:                    order.Email,
		TaxDestination:           order.TaxDestination,
		Items:                    fromConvertItems(order.Items),
		DeliveryMethod:           order.DeliveryMethod,
		DeliveryAddress:          fromConvertDeliveryAddress(order.DeliveryAddress),
		DeliveryCost:             order.DeliveryCost,
		DeliveryDatetime:         order.DeliveryDatetime,
		TotalAmount:              order.TotalAmount,
		DeliverySlot:             order.DeliverySlot,
		Promotions:               order.Promotions,
		Coupons:                  order.Coupons,
		Payments:                 fromConvertPayments(order.Payments),
	}
}

func fromConvertDeliveryAddress(address domain.DeliveryAddress) modelv1.DeliveryAddress {
	return modelv1.DeliveryAddress(address)
}

func fromConvertItems(items []domain.Item) []modelv1.Item {
	result := make([]modelv1.Item, len(items))
	for i, item := range items {
		result[i] = modelv1.Item(item)
	}
	return result
}

func fromConvertPayments(payments []domain.Payment) []modelv1.Payment {
	result := make([]modelv1.Payment, len(payments))
	for i, payment := range payments {
		result[i] = modelv1.Payment(payment)
	}
	return result
}
