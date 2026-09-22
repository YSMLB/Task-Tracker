package service

import (
	"fmt"
	"os"
	"time"

	"task-tracker/internal/storage"
)

func AddTask(description string) error{
	maxID := 0
	argWithProg := os.Args
	if len(argWithProg) < 3{
		fmt.Println("описания нет")
		return fmt.Errorf("error")
	}

	read, err:= storage.ReadTask()
	if err != nil{
		return err
	}else{
		for _, task := range read{
			if task.ID > maxID{
				maxID = task.ID
			}
		}
	}
	id := maxID + 1

	newTask := storage.Task{
		ID: id,
		Description: description,
		Status: "todo",
		CreatedAt: time.Now(),
		UpdateAt: time.Now(),
	}

	read = append(read, newTask)
	
	storage.WriteTask(read)

	return nil
}

func ListTask() error{
	read, err := storage.ReadTask()


	if err != nil{
		fmt.Println("неудалось прочитать список")
		return err
	}

	if len(read) == 0{
		fmt.Println("Список задач пуст")
	}else{
		for _, task := range read{
			fmt.Println("ID: ", task.ID, 
			"\n Задача: ", task.Description,
			"\n Статус:", task.Status,
			"\n Время создания задачи: ", task.CreatedAt, 
			"\n Время обновления: ", task.UpdateAt,
		)

		}
	}
	return nil
}

func listStatTask(){

}

func markInProgressTask(){

}

func markDoneTask(){

}

func updateTask(){

}

func deleteTask(){

}

//Команда add:
//
//Проверить, передано ли описание задачи. - +
//Прочитать текущий список задач. - +
//Сгенерировать новый ID. - +
//Зафиксировать текущее время для createdAt и updatedAt. - 
//Установить статус todo. - 
//Добавить задачу в список и сохранить в файл. - 
//Вывести подтверждение: Task added successfully (ID: X). - 
//
//Команда list (без фильтра):
//
//Прочитать задачи.
//Если задач нет — вывести информативное сообщение (например, «Список задач пуст»).
//Если задачи есть — вывести их на экран в понятном виде (список или таблица с колонками ID, Статус, Описание, Дата).
//Команда list <status>:
//
//Проверить, что переданный статус валиден (todo, in-progress или done).
//Отфильтровать задачи по этому статусу и вывести.
//
//Команды mark-in-progress и mark-done:
//
//Проверить, передан ли ID.
//Найти задачу с таким ID.
//Изменить её статус и обновить поле updatedAt текущим временем.
//Сохранить изменения.
//
//Команда update:
//
//Проверить, переданы ли оба обязательных аргумента: ID и новое описание.
//Найти задачу, обновить её description и updatedAt.
//Сохранить изменения.
//
//Команда delete:
//
//Найти задачу по ID.
//Удалить её из списка и сохранить обновлённый список.