package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/centraldigital/cfw-core-customer-api/pkg/enum"
	"github.com/centraldigital/cfw-core-lib/pkg/util/typeconvertutil"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetCustomersByStaffId(t *testing.T) {
	t.Run("01_success", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()
		tm.repository.On("GetCustomersByStaffId", ctx, "47cd52d1-3e09-4143-9371-0178029b9037").Return([]string{"2312000105", "2312000106"}, nil)
		tm.coreCustomerAdapter.On("GetCustomersProfilesByIds", ctx, domain.GetCustomersByStaffIdRequest{
			StaffId:   "47cd52d1-3e09-4143-9371-0178029b9037",
			StoreCode: "001",
		}, []string{"2312000105", "2312000106"}).Return(&domain.GetCustomersByStaffIdResponse{
			Customers: []domain.Customer{
				{
					CustomerId:  "2312000105",
					CompanyName: nil,
					Type:        enum.CustomerTypeEndUser,
					Firstname:   "John",
					Lastname:    "Doe",
					MobileNo:    "0812345678",
				},
				{
					CustomerId:  "2312000106",
					CompanyName: typeconvertutil.ToPtr("Central"),
					Type:        enum.CustomerTypeJuristic,
					Firstname:   "Jane",
					Lastname:    "Doe",
					MobileNo:    "0812345679",
				},
			},
		}, nil)
		_, err := tm.service.GetCustomersByStaffId(ctx, domain.GetCustomersByStaffIdRequest{
			StaffId:   "47cd52d1-3e09-4143-9371-0178029b9037",
			StoreCode: "001"})
		assert.NoError(t, err)
	})

	t.Run("02_success_with_empty_customer_list", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()
		tm.repository.On("GetCustomersByStaffId", ctx, "47cd52d1-3e09-4143-9371-0178029b9037").Return([]string{}, nil)
		_, err := tm.service.GetCustomersByStaffId(ctx, domain.GetCustomersByStaffIdRequest{
			StaffId:   "47cd52d1-3e09-4143-9371-0178029b9037",
			StoreCode: "001"})
		assert.NoError(t, err)
	})

	t.Run("03_fail_error_getting_customers_by_staff_id", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()
		tm.repository.On("GetCustomersByStaffId", ctx, "47cd52d1-3e09-4143-9371-0178029b9037").Return(nil, errors.New(("Error getting customers by staff id")))
		_, err := tm.service.GetCustomersByStaffId(ctx, domain.GetCustomersByStaffIdRequest{
			StaffId:   "47cd52d1-3e09-4143-9371-0178029b9037",
			StoreCode: "001"})
		assert.Error(t, err)
	})

	t.Run("04_fail_error_getting_customer_profiles_by_ids", func(t *testing.T) {
		tm := newTestingModule(t)
		ctx := context.Background()
		tm.repository.On("GetCustomersByStaffId", ctx, "47cd52d1-3e09-4143-9371-0178029b9037").Return([]string{"2312000105", "2312000106"}, nil)
		tm.coreCustomerAdapter.On("GetCustomersProfilesByIds", ctx, domain.GetCustomersByStaffIdRequest{
			StaffId:   "47cd52d1-3e09-4143-9371-0178029b9037",
			StoreCode: "001",
		}, []string{"2312000105", "2312000106"}).Return(nil, errors.New("Error getting customer profiles by ids"))
		_, err := tm.service.GetCustomersByStaffId(ctx, domain.GetCustomersByStaffIdRequest{
			StaffId:   "47cd52d1-3e09-4143-9371-0178029b9037",
			StoreCode: "001"})
		assert.Error(t, err)
	})

}
