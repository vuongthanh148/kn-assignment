package domainadapter

import (
	"github.com/centraldigital/cfw-core-lib/pkg/model/basemodel"
	"github.com/centraldigital/cfw-core-lib/pkg/model/jsonmodel"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

type Order struct {
	OrderID                  string                  `json:"order_id"`
	OrderStatus              string                  `json:"order_status"`
	OrderDatetime            jsonmodel.DateTimeLocal `json:"order_datetime"`
	OrderCreationDatetime    jsonmodel.DateTimeLocal `json:"order_creation_datetime"`
	OrderUpdatedDatetime     jsonmodel.DateTimeLocal `json:"order_updated_datetime"`
	StoreCode                string                  `json:"store_code"`
	CustomerSegment          string                  `json:"customer_segment"`
	SaleChannel              string                  `json:"sale_channel"`
	NetAmount                float64                 `json:"net_amount"`
	TotalDiscount            float64                 `json:"total_discount"`
	TotalMarketingDiscount   float64                 `json:"total_marketing_discount"`
	TotalMerchandizeDiscount float64                 `json:"total_merchandize_discount"`
	MobileNo                 *string                 `json:"mobile_no"`
	Email                    *string                 `json:"email"`
	TaxDestination           *string                 `json:"tax_destination"`
	Items                    []Item                  `json:"items"`
	DeliveryMethod           string                  `json:"delivery_method"`
	DeliveryAddress          DeliveryAddress         `json:"delivery_address"`
	DeliveryCost             float64                 `json:"delivery_cost"`
	DeliveryDatetime         jsonmodel.DateTimeLocal `json:"delivery_datetime"`
	TotalAmount              float64                 `json:"total_amount"`
	DeliverySlot             string                  `json:"delivery_slot"`
	Promotions               []struct{}              `json:"promotions"`
	Coupons                  []struct{}              `json:"coupons"`
	Payments                 []Payment               `json:"payments"`
}

type Item struct {
	LineItemNo          int     `json:"line_item_no"`
	SkuCode             string  `json:"sku_code"`
	PrCode              string  `json:"pr_code"`
	Quantity            int     `json:"quantity"`
	IsWeightItem        bool    `json:"is_weight_item"`
	AvgWeight           float64 `json:"avg_weight"`
	QuantityWeightItem  float64 `json:"quantity_weight_item"`
	SaleUnit            string  `json:"sale_unit"`
	LineItemPrice       float64 `json:"line_item_price"`
	UnitPrice           float64 `json:"unit_price"`
	PromotionCode       string  `json:"promotion_code"`
	ReferenceLineItemNo string  `json:"reference_line_item_no"`
}

type DeliveryAddress struct {
	AddressID   string `json:"address_id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	IsCompany   bool   `json:"is_company"`
	CompanyName string `json:"company_name"`
	TaxID       string `json:"tax_id"`
	BranchID    string `json:"branch_id"`
	PhoneNo     string `json:"phone_no"`
	AddressNo   string `json:"address_no"`
	Building    string `json:"building"`
	Floor       string `json:"floor"`
	Room        string `json:"room"`
	Moo         string `json:"moo"`
	Soi         string `json:"soi"`
	Road        string `json:"road"`
	Subdistrict string `json:"subdistrict"`
	District    string `json:"district"`
	Province    string `json:"province"`
	Zipcode     string `json:"zipcode"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}

type Payment struct {
	PaymentID       string                   `json:"payment_id"`
	PaymentType     string                   `json:"payment_type"`
	PaymentMethod   string                   `json:"payment_method"`
	PaymentAmount   float64                  `json:"payment_amount"`
	Tendor          string                   `json:"tendor"`
	Creditcard      string                   `json:"creditcard"`
	ApproveCode     string                   `json:"approve_code"`
	BatchID         string                   `json:"batch_id"`
	TraceNo         string                   `json:"trace_no"`
	PaymentDatetime *jsonmodel.DateTimeLocal `json:"payment_datetime"`
	PaymentStatus   string                   `json:"payment_status"`
	PaymentReason   *string                  `json:"payment_reason"`
	CreatedAt       jsonmodel.DateTimeLocal  `json:"created_at"`
	CreatedBy       string                   `json:"created_by"`
	UpdatedAt       jsonmodel.DateTimeLocal  `json:"updated_at"`
	UpdatedBy       string                   `json:"updated_by"`
	PaidAt          *jsonmodel.DateTimeLocal `json:"paid_at"`
	AdditionalInfo  *struct {
		KbankOrderID string `json:"kbank_order_id"`
	} `json:"additional_info"`
}

type GetOrderHistoryResponse struct {
	Orders     []Order                            `json:"orders"`
	Pagination basemodel.PaginationOffsetResponse `json:"pagination"`
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
	orders := make([]Order, len(d.Orders))
	for i, order := range d.Orders {
		orders[i] = fromConvertOrderHistory(order)
	}

	return &GetOrderHistoryResponse{
		Orders:     orders,
		Pagination: d.Pagination,
	}
}

// utility to domain
func toConvertOrderHistory(order Order) domain.Order {
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

func toConvertDeliveryAddress(address DeliveryAddress) domain.DeliveryAddress {
	return domain.DeliveryAddress(address)
}

func toConvertItems(items []Item) []domain.Item {
	result := make([]domain.Item, len(items))
	for i, item := range items {
		result[i] = domain.Item(item)
	}
	return result
}

func toConvertPayments(payments []Payment) []domain.Payment {
	result := make([]domain.Payment, len(payments))
	for i, payment := range payments {
		result[i] = domain.Payment(payment)
	}
	return result
}

// utility form domain
func fromConvertOrderHistory(order domain.Order) Order {
	return Order{
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

func fromConvertDeliveryAddress(address domain.DeliveryAddress) DeliveryAddress {
	return DeliveryAddress(address)
}

func fromConvertItems(items []domain.Item) []Item {
	result := make([]Item, len(items))
	for i, item := range items {
		result[i] = Item(item)
	}
	return result
}

func fromConvertPayments(payments []domain.Payment) []Payment {
	result := make([]Payment, len(payments))
	for i, payment := range payments {
		result[i] = Payment(payment)
	}
	return result
}
