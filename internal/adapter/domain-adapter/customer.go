package domainadapter

import (
	"github.com/centraldigital/cfw-core-customer-api/pkg/enum"
	customer "github.com/centraldigital/cfw-core-customer-api/pkg/model/v1"
	"github.com/centraldigital/cfw-core-lib/pkg/util/parseutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

type GetCustomerByMemberIdsRequest struct {
	MemberIds         []string     `json:"member_ids"`
	RegisteredStoreId []string     `json:"registered_store_id"`
	OrderBy           enum.OrderBy `json:"order_by,omitempty"`
	ShowInactive      bool         `json:"show_inactive,omitempty"`
	PageSize          int          `json:"page_size,omitempty"`
}

func (GetCustomerByMemberIdsRequest) FromDomain(d domain.GetCustomersByStaffIdRequest, customerIds []string) GetCustomerByMemberIdsRequest {
	return GetCustomerByMemberIdsRequest{
		MemberIds:         customerIds,
		RegisteredStoreId: []string{d.StoreCode},
	}
}

type GetCustomersResponse customer.GetCustomersResponse

func (dma GetCustomersResponse) ToDomain() *domain.GetCustomersByStaffIdResponse {
	customers := make([]domain.Customer, len(dma.Customers))
	for i, customer := range dma.Customers {
		customers[i] = domain.Customer{
			CustomerId:  customer.GeneralInfo.MemberID,
			CompanyName: customer.GeneralInfo.CompanyName,
			Type:        customer.GeneralInfo.Type,
			Firstname:   customer.GeneralInfo.FirstName,
			Lastname:    customer.GeneralInfo.LastName,
			MobileNo:    customer.GeneralInfo.MobileNo,
		}
	}
	return &domain.GetCustomersByStaffIdResponse{
		Customers: customers,
	}
}
func (dma GetCustomersResponse) ToDomainCustomer() *domain.GetCustomersByCustomerIdResponse {
	customers := make([]domain.GetCustomerResponse, 0, len(dma.Customers))
	for i, customer := range dma.Customers {
		deliveryAddress := make([]domain.Addresses, 0)
		for _, da := range dma.Customers[i].DeliveryAddresses {
			deliveryAddress = append(deliveryAddress, domain.Addresses{
				AddressId:      da.AddressID,
				Name:           *parseutil.Str(da.Name).WithDefault("").ToVal(),
				Building:       *parseutil.Str(da.Building).WithDefault("").ToVal(),
				HouseNo:        *parseutil.Str(da.HouseNo).WithDefault("").ToVal(),
				Alley:          *parseutil.Str(da.Alley).WithDefault("").ToVal(),
				Street:         *parseutil.Str(da.Street).WithDefault("").ToVal(),
				Road:           *parseutil.Str(da.Road).WithDefault("").ToVal(),
				SubDistrict:    da.SubDistrict,
				District:       da.District,
				Province:       da.Province,
				PostalCode:     da.PostalCode,
				MobileNo:       *parseutil.Str(da.MobileNo).WithDefault("").ToVal(),
				IsDefault:      *parseutil.Bool(da.IsDefault).WithDefault(false).ToVal(),
				AdditionalInfo: nil,
			})
		}
		var billingAddress *domain.Addresses
		if customer.BillingAddress != nil {
			billingAddress = &domain.Addresses{
				AddressId:      customer.BillingAddress.AddressID,
				Name:           *parseutil.Str(customer.BillingAddress.Name).WithDefault("").ToVal(),
				Building:       *parseutil.Str(customer.BillingAddress.Building).WithDefault("").ToVal(),
				HouseNo:        *parseutil.Str(customer.BillingAddress.HouseNo).WithDefault("").ToVal(),
				Alley:          *parseutil.Str(customer.BillingAddress.Alley).WithDefault("").ToVal(),
				Street:         *parseutil.Str(customer.BillingAddress.Street).WithDefault("").ToVal(),
				Road:           *parseutil.Str(customer.BillingAddress.Road).WithDefault("").ToVal(),
				SubDistrict:    customer.BillingAddress.SubDistrict,
				District:       customer.BillingAddress.District,
				Province:       customer.BillingAddress.Province,
				PostalCode:     customer.BillingAddress.PostalCode,
				MobileNo:       *parseutil.Str(customer.BillingAddress.MobileNo).WithDefault("").ToVal(),
				IsDefault:      *parseutil.Bool(customer.BillingAddress.IsDefault).WithDefault(false).ToVal(),
				AdditionalInfo: nil,
			}
		}

		customers = append(customers, domain.GetCustomerResponse{
			GeneralInformation: domain.GeneralInformation{
				Customer_id:           customer.GeneralInfo.MemberID,
				CompanyName:           *parseutil.Str(customer.GeneralInfo.CompanyName).WithDefault("").ToVal(),
				RegisteredCompanyName: *parseutil.Str(customer.GeneralInfo.RegisteredCompanyName).WithDefault("").ToVal(),
				Type:                  customer.GeneralInfo.Type,
				TaxId:                 *parseutil.Str(customer.GeneralInfo.TaxId).WithDefault("").ToVal(),
				TaxBranchCode:         *parseutil.Str(customer.GeneralInfo.TaxBranchCode).WithDefault("").ToVal(),
				Firstname:             customer.GeneralInfo.FirstName,
				Lastname:              customer.GeneralInfo.LastName,
				CitizenId:             *parseutil.Str(customer.GeneralInfo.CitizenID).WithDefault("").ToVal(),
				PassportNo:            *parseutil.Str(customer.GeneralInfo.PassportNo).WithDefault("").ToVal(),
				MobileNo:              customer.GeneralInfo.MobileNo,
				Email:                 *parseutil.Str(customer.GeneralInfo.Email).WithDefault("").ToVal(),
				IsRequiredInvoice:     *parseutil.Bool(customer.GeneralInfo.IsRequiredInvoice).WithDefault(false).ToVal(),
			},
			DeliveryAddresses: deliveryAddress,
			BillingAddress:    billingAddress,
			Segment: domain.Segment{
				SegmentName:      customer.Segment.SegmentName,
				SegmentNameTh:    customer.Segment.SegmentNameTH,
				SegmentCode:      customer.Segment.SegmentCode,
				SubSegmentName:   customer.Segment.SubSegmentName,
				SubSegmentNameTh: customer.Segment.SubSegmentNameTH,
				SubSegmentCode:   customer.Segment.SubSegmentCode,
				TypeName:         customer.Segment.TypeName,
				TypeNameTh:       customer.Segment.TypeNameTH,
				TypeCode:         customer.Segment.TypeCode,
			},
			The1: domain.The1{
				The1MemberId:    customer.The_1.The1MemberID,
				The1MobileNo:    customer.The_1.The1MobileNo,
				The1FirstNameTh: customer.The_1.The1FirstNameTh,
				The1FirstNameEn: customer.The_1.The1FirstNameEn,
				The1LastNameTh:  customer.The_1.The1LastNameTh,
				The1LastNameEn:  customer.The_1.The1LastNameEn,
				The1Point:       int32(customer.The_1.The1Point),
			},
			Legal: domain.Legal{
				HasAlcoholLicense:      customer.Legal.HasAlcoholLicense,
				AlcoholExpirationDate:  customer.Legal.AlcoholExpirationDate,
				HasCompanyRegistration: customer.Legal.HasCompanyRegistration,
				HasVatCertificate:      customer.Legal.HasVatCertificate,
			},
			IsActive:              customer.IsActive,
			LastUpdatedAt:         "",
			UpdatedBy:             customer.UpdatedBy,
			AlcoholExpirationDate: customer.Legal.AlcoholExpirationDate,
		})
	}

	return &domain.GetCustomersByCustomerIdResponse{
		Customers: customers,
	}

}
