package food_file_service

import (
	"FoodService/core/enums/food_enum/queue_name_enum"
	"FoodService/core/type_utils"
	"FoodService/interface/provider"
	"FoodService/interface/repository"
	"FoodService/model/api_model/request_model"
	"bufio"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo/v4"
	"strings"
)

type foodFileService struct {
	IFoodRepository  repository.IFoodRepository
	IRabbitMQSupport provider.IRabbitMQSupport
}

func NewFoodFileService(
	foodRepository repository.IFoodRepository,
	iRabbitMQSupport provider.IRabbitMQSupport) *foodFileService {
	return &foodFileService{
		IFoodRepository:  foodRepository,
		IRabbitMQSupport: iRabbitMQSupport,
	}
}

func (sv *foodFileService) AddFoodByFileExcel(c echo.Context) error {

	fileReq, errReq := c.FormFile("file")
	if errReq != nil {
		return errReq
	}
	file, errFile := fileReq.Open()
	if errFile != nil {
		return errFile
	}
	reader := bufio.NewReader(file)
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return err
	}
	foodList := make([]request_model.AddFoodBodyRequestDto, 0)
	//cellVal := f.GetCellValue("Sheet1" , "A1")
	rows := f.GetRows("Sheet1")
	if len(rows) <= 1 {
		return nil
	}
	foodRecipe := make([]request_model.AddRecipeRequestDto, 0)
	previous := rows[1][0]
	var itemFood request_model.AddFoodBodyRequestDto
	for i := 1; i < len(rows); i++ {
		if len(rows[i][0]) > 0 {
			if previous != rows[i][0] {
				itemFood.Recipes = foodRecipe
				foodList = append(foodList, itemFood)
				foodRecipe = nil
			}
			itemFood.Name = rows[i][0]
			itemFood.TypeFood = type_utils.ConvertStringToInt(rows[i][1])
			itemFood.Status = type_utils.ConvertStringToInt(rows[i][2])
			itemFood.Description = rows[i][3]
			previous = itemFood.Name
		}
		foodRecipe = append(foodRecipe, request_model.AddRecipeRequestDto{
			Name:        rows[i][4],
			Description: rows[i][5],
			Keyword:     strings.Split(rows[i][6], ","),
			Price:       type_utils.ConvertStringToFloat(rows[i][7]),
			Level:       type_utils.ConvertStringToInt(rows[i][8]),
			Images:      strings.Split(rows[i][9], ","),
		})
	}
	itemFood.Recipes = foodRecipe
	foodList = append(foodList, itemFood)
	dataString, errConvert := type_utils.ConvertStructToJSONString(foodList)
	if errConvert != nil {
		return errConvert
	}
	errPubQueue := sv.IRabbitMQSupport.Publish(queue_name_enum.InsertMany().Name, dataString)
	if errPubQueue != nil {
		return errPubQueue
	}
	return nil
}
