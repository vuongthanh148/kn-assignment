package domain

type Store struct {
	Code      string
	Channel   string
	NameTh    string
	NameEn    string
	AddressTh string
	AddressEn string
}

type GetStoresByStaffIdRequest struct {
	StaffId string
}

type GetStoresByStaffIdResponse struct {
	Stores []Store
}
