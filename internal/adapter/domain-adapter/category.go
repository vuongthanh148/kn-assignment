package domainadapter

import (
	pbmockv1 "github.com/centraldigital/cfw-sales-x-ordering-api/gen/proto"
	"github.com/centraldigital/cfw-sales-x-ordering-api/internal/core/domain"
)

type Category pbmockv1.Category

func (da *Category) ToDomain() *domain.CategoryData {
	if da == nil {
		return nil
	}

	subCategoryData := make([]domain.SubCategoryData, 0, len(da.SubCategories))
	for i := range da.SubCategories {
		subCat := (*SubCategory)(da.SubCategories[i]).ToDomain()
		if subCat != nil {
			subCategoryData = append(subCategoryData, *subCat)
		}
	}

	return &domain.CategoryData{
		Code:            da.Code,
		NameTh:          da.NameTh,
		NameEn:          da.NameEn,
		ImageUrl:        da.ImageUrl,
		SubCategoryData: subCategoryData,
	}
}

type SubCategory pbmockv1.SubCategory

func (da *SubCategory) ToDomain() *domain.SubCategoryData {
	if da == nil {
		return nil
	}

	return &domain.SubCategoryData{
		Code:     da.Code,
		NameTh:   da.NameTh,
		NameEn:   da.NameEn,
		ImageUrl: da.ImageUrl,
	}
}

type CategoryGetResp pbmockv1.CategoryResponse

func (da *CategoryGetResp) ToDomain() *domain.CategoryGetResp {
	if da == nil {
		return nil
	}

	categories := make([]domain.CategoryData, 0, len(da.Categories))
	for i := range da.Categories {
		cat := (*Category)(da.Categories[i]).ToDomain()
		if cat != nil {
			categories = append(categories, *cat)
		}
	}

	return &domain.CategoryGetResp{
		Category: categories,
	}
}
