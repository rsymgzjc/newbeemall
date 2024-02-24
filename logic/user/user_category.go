package user

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

const (
	LevelOne   = 1
	LevelTwo   = 2
	LevelThree = 3
)

func GetGoodsCategory() (datas []*models.IndexCategory, err error) {
	datas1, err := mysql.SelectGoodsCategoryOne(LevelOne)
	if err != nil {
		zap.L().Error("查询一级分类失败了", zap.Error(err))
		return
	}
	if len(datas1) > 0 {
		var firstLevelCategoryIds []int64
		for _, data := range datas1 {
			firstLevelCategoryIds = append(firstLevelCategoryIds, data.CategoryID)
		}
		datas2, err := mysql.SelectGoodsCategoryTwo(LevelTwo, firstLevelCategoryIds)
		if err != nil {
			zap.L().Error("查询二级分类失败了", zap.Error(err))
			return
		}
		if len(datas2) > 0 {
			var secondLevelCategoryIds []int64
			for _, data := range datas2 {
				secondLevelCategoryIds = append(secondLevelCategoryIds, data.CategoryID)
			}
			datas3, err := mysql.SelectGoodsCategoryThree(LevelThree, secondLevelCategoryIds)
			if err != nil {
				zap.L().Error("查询三级分类失败了", zap.Error(err))
				return
			}
			thirdLevelCategoryMap := map[int][]*models.AdminGoodsCategory{}
			for _, data := range datas3 {
				thirdLevelCategoryMap[data.ParentId] = []*models.AdminGoodsCategory{}
			}
			for k, v := range thirdLevelCategoryMap {
				for _, third := range datas3 {
					if k == third.ParentId {
						v = append(v, third)
					}
					thirdLevelCategoryMap[k] = v
				}
			}
			var secondLevelCategory []models.SecondLevelCategory
			for _, secondLevelCategory1 := range datas2 {
				var secondLevelCategoryalone models.SecondLevelCategory
				err = copier.Copy(&secondLevelCategoryalone, &secondLevelCategory1)
				if _, ok := thirdLevelCategoryMap[int(secondLevelCategory1.CategoryID)]; ok {
					tempGoodsCategories := thirdLevelCategoryMap[int(secondLevelCategory1.CategoryID)]
					var thirdLevelCategory []models.ThirdLevelCategory
					err = copier.Copy(&thirdLevelCategory, &tempGoodsCategories)
					secondLevelCategoryalone.ThirdLevelCategory = thirdLevelCategory
					secondLevelCategory = append(secondLevelCategory, secondLevelCategoryalone)
				}
			}
			if len(secondLevelCategory) > 0 {
				secondLevelCategoryMap := map[int][]*models.SecondLevelCategory{}
				for _, data := range secondLevelCategory {
					secondLevelCategoryMap[data.ParentId] = []*models.SecondLevelCategory{}
				}
				for k, v := range secondLevelCategoryMap {
					for _, second := range datas2 {
						if k == second.ParentId {
							var secondLevelCategoryalone *models.SecondLevelCategory
							err = copier.Copy(&secondLevelCategoryalone, &second)
							v = append(v, secondLevelCategoryalone)
						}
						secondLevelCategoryMap[k] = v
					}
				}
				for _, first := range datas1 {
					var indexcategory *models.IndexCategory
					err = copier.Copy(&indexcategory, &first)
					if _, ok := secondLevelCategoryMap[int(first.CategoryID)]; ok {
						tempGoodsCategories := secondLevelCategoryMap[int(first.CategoryID)]
						var secondLevelCategorys []models.SecondLevelCategory
						err = copier.Copy(&secondLevelCategorys, &tempGoodsCategories)
						indexcategory.SecondLevelCategory = secondLevelCategorys
						datas = append(datas, indexcategory)
					}
				}
			}
		}
	}
	return
}
