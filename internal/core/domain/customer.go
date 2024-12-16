package domain

import (
	"github.com/centraldigital/cfw-core-customer-api/pkg/enum"
	"github.com/centraldigital/cfw-core-lib/pkg/model/jsonmodel"
)

type Customer struct {
	CustomerId  string
	CompanyName *string
	Type        enum.CustomerType
	Firstname   string
	Lastname    string
	MobileNo    string
}

type GetCustomersByStaffIdRequest struct {
	StaffId   string
	StoreCode string
}

type GetCustomersByStaffIdResponse struct {
	Customers []Customer
}

type GeneralInformation struct {
	Customer_id           string
	CompanyName           string
	RegisteredCompanyName string
	Type                  enum.CustomerType
	TaxId                 string
	TaxBranchCode         string
	Firstname             string
	Lastname              string
	CitizenId             string
	PassportNo            string
	MobileNo              string
	Email                 string
	IsRequiredInvoice     bool
}

type Addresses struct {
	AddressId      string
	Name           string
	Building       string
	HouseNo        string
	Alley          string
	Street         string
	Road           string
	SubDistrict    string
	District       string
	Province       string
	PostalCode     string
	MobileNo       string
	IsDefault      bool
	AdditionalInfo *string
}

type Segment struct {
	SegmentName      string
	SegmentNameTh    string
	SegmentCode      string
	SubSegmentName   string
	SubSegmentNameTh string
	SubSegmentCode   string
	TypeName         string
	TypeNameTh       string
	TypeCode         string
}
type The1 struct {
	The1MemberId    string
	The1MobileNo    string
	The1FirstNameTh string
	The1FirstNameEn string
	The1LastNameTh  string
	The1LastNameEn  string
	The1Point       int32
}
type Legal struct {
	HasAlcoholLicense      bool
	AlcoholExpirationDate  *jsonmodel.DateMinTime
	HasCompanyRegistration bool
	HasVatCertificate      bool
}
type GetCustomerResponse struct {
	GeneralInformation    GeneralInformation
	DeliveryAddresses     []Addresses
	BillingAddress        *Addresses
	Segment               Segment
	The1                  The1
	Legal                 Legal
	IsActive              bool
	LastUpdatedAt         string
	UpdatedBy             string
	AlcoholExpirationDate *jsonmodel.DateMinTime
}

type GetCustomerRequest struct {
	MemberId string
}

type GetCustomersQuery struct {
	Pagesize     int32
	ShowInactive bool
	IsInactive   bool
	OrderBy      string
	MemberId     string
}
type GetCustomersByCustomerIdResponse struct {
	Customers []GetCustomerResponse
}
