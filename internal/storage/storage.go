package storage

import (
	"encoding/json"
	"fmt"
	"os"
	
	"time"
)

//Стркутура под второй метод
type Task struct{
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createat"`
	UpdateAt time.Time `json:"updateat"`
}

func readTask() ([]Task, error){
	data, err := os.ReadFile("task.json")
	if err != nil{
		fmt.Println("Список пуст")
		return nil, nil
	}

	var u []Task

	if err := json.Unmarshal(data, &u); err != nil{
		return nil, nil
	}
	return u, nil
}

func writeTask(){

}