package food_file_service

import (
	"FoodService/interface/repository"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type foodFileService struct {
	IFoodRepository repository.IFoodRepository
}

func NewFoodFileService(foodRepository repository.IFoodRepository) *foodFileService {
	return &foodFileService{
		IFoodRepository: foodRepository,
	}
}

func (sv *foodFileService) AddFoodByFileExcel() error {
	f, err := excelize.OpenFile("resource/tmp.xlsx")
	if err != nil {
		return err
	}

	rows := f.GetRows("Trang tính1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
	}
	return nil
}
