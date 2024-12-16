package domain

import (
	"github.com/centraldigital/cfw-core-lib/pkg/model/basemodel"
	"github.com/centraldigital/cfw-core-lib/pkg/model/jsonmodel"
)

type Order struct {
	OrderID                  string
	OrderStatus              string
	OrderDatetime            jsonmodel.DateTimeLocal
	OrderCreationDatetime    jsonmodel.DateTimeLocal
	OrderUpdatedDatetime     jsonmodel.DateTimeLocal
	StoreCode                string
	CustomerSegment          string
	SaleChannel              string
	NetAmount                float64
	TotalDiscount            float64
	TotalMarketingDiscount   float64
	TotalMerchandizeDiscount float64
	MobileNo                 *string
	Email                    *string
	TaxDestination           *string
	Items                    []Item
	DeliveryMethod           string
	DeliveryAddress          DeliveryAddress
	DeliveryCost             float64
	DeliveryDatetime         jsonmodel.DateTimeLocal
	TotalAmount              float64
	DeliverySlot             string
	Promotions               []struct{}
	Coupons                  []struct{}
	Payments                 []Payment
}

type Item struct {
	LineItemNo          int
	SkuCode             string
	PrCode              string
	Quantity            int
	IsWeightItem        bool
	AvgWeight           float64
	QuantityWeightItem  float64
	SaleUnit            string
	LineItemPrice       float64
	UnitPrice           float64
	PromotionCode       string
	ReferenceLineItemNo string
}

type DeliveryAddress struct {
	AddressID   string
	Firstname   string
	Lastname    string
	IsCompany   bool
	CompanyName string
	TaxID       string
	BranchID    string
	PhoneNo     string
	AddressNo   string
	Building    string
	Floor       string
	Room        string
	Moo         string
	Soi         string
	Road        string
	Subdistrict string
	District    string
	Province    string
	Zipcode     string
	Latitude    string
	Longitude   string
}

type Payment struct {
	PaymentID       string
	PaymentType     string
	PaymentMethod   string
	PaymentAmount   float64
	Tendor          string
	Creditcard      string
	ApproveCode     string
	BatchID         string
	TraceNo         string
	PaymentDatetime *jsonmodel.DateTimeLocal
	PaymentStatus   string
	PaymentReason   *string
	CreatedAt       jsonmodel.DateTimeLocal
	CreatedBy       string
	UpdatedAt       jsonmodel.DateTimeLocal
	UpdatedBy       string
	PaidAt          *jsonmodel.DateTimeLocal
	AdditionalInfo  *struct {
		KbankOrderID string
	}
}

type GetOrderHistoryResponse struct {
	Orders     []Order
	Pagination basemodel.PaginationOffsetResponse
}

type GetOrderHistoryRequest struct {
	CustomerId string
	PageID     int
	PageSize   int
}
