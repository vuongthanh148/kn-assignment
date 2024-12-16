package modelv1

import (
	"github.com/centraldigital/cfw-core-customer-api/pkg/enum"
	"github.com/centraldigital/cfw-core-lib/pkg/model/jsonmodel"
)

type Customer struct {
	CustomerId  string            `json:"customer_id"`
	CompanyName *string           `json:"company_name"`
	Type        enum.CustomerType `json:"type"`
	Firstname   string            `json:"firstname"`
	Lastname    string            `json:"lastname"`
	MobileNo    string            `json:"mobile_no"`
}

type GetCustomersByStaffIdRequest struct {
	StaffId   string `json:"-" uri:"staff-id"`
	StoreCode string `json:"-" uri:"store-code"`
}

type GetCustomersByStaffIdResponse struct {
	Customers []Customer `json:"customers"`
}

type GeneralInformation struct {
	Customer_id           string            `json:"customer_id"`
	CompanyName           string            `json:"company_name"`
	RegisteredCompanyName string            `json:"registered_company_name"`
	Type                  enum.CustomerType `json:"type"`
	TaxId                 string            `json:"tax_id"`
	TaxBranchCode         string            `json:"tax_branch_code"`
	Firstname             string            `json:"firstname"`
	Lastname              string            `json:"lastname"`
	CitizenId             string            `json:"citizen_id"`
	PassportNo            string            `json:"passport_no"`
	MobileNo              string            `json:"mobile_no"`
	Email                 string            `json:"email"`
	IsRequiredInvoice     bool              `json:"is_required_invoice"`
}

type Addresses struct {
	AddressId      string  `json:"address_id"`
	Name           string  `json:"name"`
	Building       string  `json:"building"`
	HouseNo        string  `json:"house_no"`
	Alley          string  `json:"alley"`
	Street         string  `json:"street"`
	Road           string  `json:"road"`
	SubDistrict    string  `json:"sub_district"`
	District       string  `json:"district"`
	Province       string  `json:"province"`
	PostalCode     string  `json:"postal_code"`
	MobileNo       string  `json:"mobile_no"`
	IsDefault      bool    `json:"is_default"`
	AdditionalInfo *string `json:"additional_info"`
}

type Segment struct {
	SegmentName      string `json:"segment_name"`
	SegmentNameTh    string `json:"segment_name_th"`
	SegmentCode      string `json:"segment_code"`
	SubSegmentName   string `json:"sub_segment_name"`
	SubSegmentNameTh string `json:"sub_segment_name_th"`
	SubSegmentCode   string `json:"sub_segment_code"`
	TypeName         string `json:"type_name"`
	TypeNameTh       string `json:"type_name_th"`
	TypeCode         string `json:"type_code"`
}
type The1 struct {
	The1MemberId    string `json:"the_1_member_id"`
	The1MobileNo    string `json:"the_1_mobile_no"`
	The1FirstNameTh string `json:"the_1_first_name_th"`
	The1FirstNameEn string `json:"the_1_first_name_en"`
	The1LastNameTh  string `json:"the_1_last_name_th"`
	The1LastNameEn  string `json:"the_1_last_name_en"`
	The1Point       int32  `json:"the_1_point"`
}
type Legal struct {
	HasAlcoholLicense      bool                   `json:"has_alcohol_license"`
	AlcoholExpirationDate  *jsonmodel.DateMinTime `json:"alcohol_expiration_date"`
	HasCompanyRegistration bool                   `json:"has_company_registration"`
	HasVatCertificate      bool                   `json:"has_vat_certificate"`
}
type GetCustomerResponse struct {
	GeneralInformation    GeneralInformation     `json:"general_information"`
	DeliveryAddresses     []Addresses            `json:"delivery_addresses"`
	BillingAddress        *Addresses             `json:"billing_address"`
	Segment               Segment                `json:"segment"`
	The1                  The1                   `json:"the_1"`
	Legal                 Legal                  `json:"legal"`
	IsActive              bool                   `json:"is_active"`
	LastUpdatedAt         string                 `json:"last_updated_at"`
	UpdatedBy             string                 `json:"updated_by"`
	AlcoholExpirationDate *jsonmodel.DateMinTime `json:"alcohol_expiration_date"`
}

type GetCustomerRequest struct {
	MemberId string `json:"-" uri:"customer-id"`
}
