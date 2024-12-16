package modelv1

type CategoryData struct {
	Code            string            `json:"code"`
	NameTh          string            `json:"name_th"`
	NameEn          string            `json:"name_en"`
	ImageUrl        string            `json:"image_url"`
	SubCategoryData []SubCategoryData `json:"sub_categories"`
}

type SubCategoryData struct {
	Code     string `json:"code"`
	NameTh   string `json:"name_th"`
	NameEn   string `json:"name_en"`
	ImageUrl string `json:"image_url"`
}

type CategoryGetResp struct {
	Category []CategoryData `json:"categories"`
}
