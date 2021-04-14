package food_consume_service

import (
	"FoodService/interface/service"
	"FoodService/model/api_model/request_model"
	"FoodService/repository/food_repository"
	"FoodService/repository/recipe_repository"
	"FoodService/service/food_service"
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
)

type foodConsumeService struct {
	IFoodServiceInterface service.IFoodServiceInterface
}

func NewFoodConsumerService(
	IFoodServiceInterface service.IFoodServiceInterface) *foodConsumeService {
	return &foodConsumeService{
		IFoodServiceInterface: IFoodServiceInterface,
	}
}

func (sv *foodConsumeService) InsertManyFood(bodyConsume []request_model.AddFoodBodyRequestDto) error {

	jobs := make(chan request_model.AddFoodBodyRequestDto, len(bodyConsume))
	errChan := make(chan error, len(bodyConsume))
	worker := 3
	var wg sync.WaitGroup
	for w := 1; w <= worker; w++ {
		wg.Add(1)
		go sv.workerInsert(&wg, w, jobs, errChan)
	}

	for _, item := range bodyConsume {
		jobs <- item
	}

	go func() {
		wg.Wait()
		close(errChan)
		close(jobs)
	}()

	for item := range errChan {
		if item != nil {
			fmt.Printf("Error Handling")
		}
	}

	return nil
}

func (sv *foodConsumeService) workerInsert(wg *sync.WaitGroup, worker int, jobs <-chan request_model.AddFoodBodyRequestDto, errChan chan<- error) {
	wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d start insert food with name %s \n", worker, j.Name)
		err := sv.IFoodServiceInterface.AddFoodService(j)
		if err != nil {
			fmt.Printf("Insert Fail : Food Name %s , Error thrown %s \n ", j.Name, err.Error())
			errChan <- err
		}
		errChan <- nil
	}
}

type FoodConsume struct {
	db *sql.DB
}

func (fs FoodConsume) Run(body []byte) error {
	var request []request_model.AddFoodBodyRequestDto
	if err := json.Unmarshal(body, &request); err != nil {
		return err
	}
	NewFoodRepository := food_repository.NewFoodRepository(fs.db)
	NewRecipeRepository := recipe_repository.NewRecipeRepository(fs.db)
	NewFoodService := food_service.NewFoodService(NewFoodRepository, NewRecipeRepository)
	NewFoodConsumerService := NewFoodConsumerService(NewFoodService)
	err := NewFoodConsumerService.InsertManyFood(request)
	if err != nil {
		return err
	}
	return nil
}
