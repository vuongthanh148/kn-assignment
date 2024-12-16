package modelv1

type Store struct {
	Code      string `json:"code"`
	Channel   string `json:"channel"`
	NameTh    string `json:"name_th"`
	NameEn    string `json:"name_en"`
	AddressTh string `json:"address_th"`
	AddressEn string `json:"address_en"`
}

type GetStoresByStaffIdRequest struct {
	StaffId string `json:"-" uri:"staff-id"`
}

type GetStoresByStaffIdResponse struct {
	Stores []Store `json:"stores"`
}
