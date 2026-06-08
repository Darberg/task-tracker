package main

import (
	"fmt"
	"time"
)

type Tasks struct {
	Id          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

func Add() {
	status, description := HandleAddTasks()
	tasks := FindAll()
	newTask := new(Tasks)
	newTask.Id = generateId(tasks)
	newTask.Description = description
	newTask.Status = status
	newTask.CreatedAt = time.Now().String()
	newTask.UpdatedAt = time.Now().String()
	tasks = append(tasks, *newTask)
	Save(tasks)
	PrintTaskCards(*newTask)
}

func Delete(id int) {

	task, index := FindOne(id)
	VerifyTaskExist(index)

	PrintTaskCards(task)

	fmt.Println("1. Yes")
	fmt.Println("2. No")

	var confirm int
	fmt.Scanln(&confirm)

	if confirm == 1 {

		tasks := FindAll()
		newTask := make([]Tasks, len(tasks))

		newLength := 0
		for i, v := range tasks {
			if i != index {
				newTask[newLength] = v
				newLength++
			}
		}

		Save(newTask)

		fmt.Println("Record has Been Removed")
	} else {
		fmt.Println("Invalid Option received")
	}
}

func Update() {
	fmt.Println("Please Send the Id of task you want to update :")
	var id int
	fmt.Scanln(&id)
	task, index := FindOne(id)
	if index == -1 {
		err := fmt.Errorf("Task With this id: %d not found", index)
		check(err)
	}
	PrintTaskCards(task)

	updatedTask := new(Tasks)
	updatedTask.Id = task.Id
	updatedTask.Status = HandleStatuses()
	updatedTask.CreatedAt = task.CreatedAt
	fmt.Println("Please chose the new Description: ")
	fmt.Scanln(&updatedTask.Description)
	if updatedTask.Description == "" {
		updatedTask.Description = task.Description
	}

	updatedTask.UpdatedAt = time.Now().String()

	tasks := FindAll()
	tasks[index] = *updatedTask

	Save(tasks)
}

func GetAll(page int, pageSize int) ([]Tasks, int) {

	tasks := FindAll()
	offset := (page - 1) * pageSize

	if offset > len(tasks) {
		return tasks, offset
	}

	max := offset + pageSize
	if max > len(tasks) {
		max = len(tasks)
	}

	return tasks, pageSize
}

func Get(id int) (Tasks, int) {
	tasks := FindAll()
	for i, v := range tasks {
		if v.Id == id {
			return v, i
		}
	}

	return Tasks{}, -1
}

func generateId(tasks []Tasks) int {
	id := 1

	sortTasks(tasks)

	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].Id + 1
	}

	return id
}

func sortTasks(tasks []Tasks) []Tasks {
	for i := 0; i < len(tasks); i++ {
		for j := i + 1; j < len(tasks); j++ {

			if tasks[i].Id > tasks[j].Id {
				tasks[i], tasks[j] = tasks[j], tasks[i]
			}
		}
	}
	return tasks
}

func VerifyTaskExist(index int) {
	if index == -1 {
		fmt.Errorf("Task With this id: %d not found", index)
	}
}
