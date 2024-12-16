package domain

type CategoryData struct {
	Code            string
	NameTh          string
	NameEn          string
	ImageUrl        string
	SubCategoryData []SubCategoryData
}

type SubCategoryData struct {
	Code     string
	NameTh   string
	NameEn   string
	ImageUrl string
}

type CategoryGetResp struct {
	Category []CategoryData
}
