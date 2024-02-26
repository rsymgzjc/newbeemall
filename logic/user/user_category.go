package user

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
)

const (
	LevelOne   = 1
	LevelTwo   = 2
	LevelThree = 3
)

//func GetGoodsCategory() (datas []*models.IndexCategory, err error) {
//	datas1, err := mysql.SelectGoodsCategoryOne(LevelOne)
//	if err != nil {
//		zap.L().Error("查询一级分类失败了", zap.Error(err))
//		return
//	}
//	if len(datas1) > 0 {
//		var firstLevelCategoryIds []int64
//		for _, data := range datas1 {
//			firstLevelCategoryIds = append(firstLevelCategoryIds, data.CategoryID)
//		}
//		datas2, err := mysql.SelectGoodsCategoryTwo(LevelTwo, firstLevelCategoryIds)
//		if err != nil {
//			zap.L().Error("查询二级分类失败了", zap.Error(err))
//			return
//		}
//		if len(datas2) > 0 {
//			var secondLevelCategoryIds []int64
//			for _, data := range datas2 {
//				secondLevelCategoryIds = append(secondLevelCategoryIds, data.CategoryID)
//			}
//			datas3, err := mysql.SelectGoodsCategoryThree(LevelThree, secondLevelCategoryIds)
//			if err != nil {
//				zap.L().Error("查询三级分类失败了", zap.Error(err))
//				return
//			}
//			thirdLevelCategoryMap := map[int][]*models.AdminGoodsCategory{}
//			for _, data := range datas3 {
//				thirdLevelCategoryMap[data.ParentId] = []*models.AdminGoodsCategory{}
//			}
//			for k, v := range thirdLevelCategoryMap {
//				for _, third := range datas3 {
//					if k == third.ParentId {
//						v = append(v, third)
//					}
//					thirdLevelCategoryMap[k] = v
//				}
//			}
//			var secondLevelCategory []models.SecondLevelCategory
//			for _, secondLevelCategory1 := range datas2 {
//				var secondLevelCategoryalone models.SecondLevelCategory
//				err = copier.Copy(&secondLevelCategoryalone, &secondLevelCategory1)
//				if _, ok := thirdLevelCategoryMap[int(secondLevelCategory1.CategoryID)]; ok {
//					tempGoodsCategories := thirdLevelCategoryMap[int(secondLevelCategory1.CategoryID)]
//					var thirdLevelCategory []models.ThirdLevelCategory
//					err = copier.Copy(&thirdLevelCategory, &tempGoodsCategories)
//					secondLevelCategoryalone.ThirdLevelCategory = thirdLevelCategory
//					secondLevelCategory = append(secondLevelCategory, secondLevelCategoryalone)
//				}
//			}
//			if len(secondLevelCategory) > 0 {
//				secondLevelCategoryMap := map[int][]*models.SecondLevelCategory{}
//				for _, data := range secondLevelCategory {
//					secondLevelCategoryMap[data.ParentId] = []*models.SecondLevelCategory{}
//				}
//				for k, v := range secondLevelCategoryMap {
//					for _, second := range datas2 {
//						if k == second.ParentId {
//							var secondLevelCategoryalone *models.SecondLevelCategory
//							err = copier.Copy(&secondLevelCategoryalone, &second)
//							v = append(v, secondLevelCategoryalone)
//						}
//						secondLevelCategoryMap[k] = v
//					}
//				}
//				for _, first := range datas1 {
//					var indexcategory *models.IndexCategory
//					err = copier.Copy(&indexcategory, &first)
//					if _, ok := secondLevelCategoryMap[int(first.CategoryID)]; ok {
//						tempGoodsCategories := secondLevelCategoryMap[int(first.CategoryID)]
//						var secondLevelCategorys []models.SecondLevelCategory
//						err = copier.Copy(&secondLevelCategorys, &tempGoodsCategories)
//						indexcategory.SecondLevelCategory = secondLevelCategorys
//						datas = append(datas, indexcategory)
//					}
//				}
//			}
//		}
//	}
//	return
//}

func GetGoodsCategory() (datas []*models.IndexCategory, err error) {
	datas = make([]*models.IndexCategory, 0)
	data1 := make([]*models.SecondLevelCategory, 0)
	data2 := make([]*models.ThirdLevelCategory, 0)
	datas1, err := mysql.SelectGoodsCategoryOne(LevelOne)
	if err != nil {
		zap.L().Error("查询一级分类失败了", zap.Error(err))
		return
	}
	datas2, err := mysql.SelectGoodsCategoryOne(LevelTwo)
	if err != nil {
		zap.L().Error("查询二级分类失败了", zap.Error(err))
		return
	}
	datas3, err := mysql.SelectGoodsCategoryOne(LevelThree)
	if err != nil {
		zap.L().Error("查询三级分类失败了", zap.Error(err))
		return
	}
	for _, data := range datas1 {
		datag := &models.IndexCategory{
			CategoryID:          data.CategoryID,
			CategoryLevel:       data.CategoryLevel,
			CategoryName:        data.CategoryName,
			SecondLevelCategory: nil,
		}
		datas = append(datas, datag)
	}
	for _, data22 := range datas2 {
		data := &models.SecondLevelCategory{
			CategoryID:         data22.CategoryID,
			ParentId:           data22.ParentId,
			CategoryLevel:      data22.CategoryLevel,
			CategoryName:       data22.CategoryName,
			ThirdLevelCategory: nil,
		}
		data1 = append(data1, data)
	}
	for _, data33 := range datas3 {
		data := &models.ThirdLevelCategory{
			CategoryID:    data33.CategoryID,
			ParentId:      data33.ParentId,
			CategoryLevel: data33.CategoryLevel,
			CategoryName:  data33.CategoryName,
		}
		data2 = append(data2, data)
	}
	for _, d1 := range data1 {
		for _, d2 := range data2 {
			if d1.CategoryID == int64(d2.ParentId) {
				d1.ThirdLevelCategory = append(d1.ThirdLevelCategory, d2)
			}
		}
	}

	for _, d0 := range datas {
		for _, d1 := range data1 {
			if d0.CategoryID == int64(d1.ParentId) {
				d0.SecondLevelCategory = append(d0.SecondLevelCategory, d1)
			}
		}
	}

	return datas, err
}
