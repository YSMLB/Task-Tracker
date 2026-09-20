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
		return nil, err
	}

	var u []Task

	if err := json.Unmarshal(data, &u); err != nil{
		return nil, nil
	}
	return u, nil
}

func writeTask(tasksAdd []Task) error{
	//принимаем задачу и кодируем ее в json
	data, er := json.MarshalIndent(tasksAdd, "", "    ")
	if er != nil{//обработка ошибки
		fmt.Println("ошибка при сериализации: ", er)
		return er
	}
	//делаем запись в файл
	err := os.WriteFile("task.json", data, 0644)
	if err != nil{//обработка ошибки again
		fmt.Println("Ошибка записи файла:", err)
		return err
	}

	return nil
}