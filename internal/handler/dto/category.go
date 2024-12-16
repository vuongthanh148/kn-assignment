package dto

import (
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
	modelv1 "github.com/centraldigital/cfw-sales-x-ordering-api/pkg/model/v1"
)

type CategoryData modelv1.CategoryData

type SubCategoryData modelv1.SubCategoryData

type CategoryGetResp modelv1.CategoryGetResp

func (CategoryGetResp) FromDomain(d *domain.CategoryGetResp) *CategoryGetResp {
	categories := make([]modelv1.CategoryData, len(d.Category))
	for i, cat := range d.Category {
		subCategories := make([]modelv1.SubCategoryData, len(cat.SubCategoryData))
		for j, subCat := range cat.SubCategoryData {
			subCategories[j] = modelv1.SubCategoryData(subCat)
		}
		categories[i] = modelv1.CategoryData{
			Code:            cat.Code,
			NameTh:          cat.NameTh,
			NameEn:          cat.NameEn,
			ImageUrl:        cat.ImageUrl,
			SubCategoryData: subCategories,
		}
	}
	return &CategoryGetResp{
		Category: categories,
	}
}
