package main

import "fmt"


func HandleDelete() {
	fmt.Println("Please Enter the Id of task you want to delete: ")
	var id int
	fmt.Scanln(&id)
	fmt.Println("Are you sure you want to delete this record:")
	Delete(id)
}

func PrintTaskCards(task Tasks) {
	fmt.Println("---------------------------------")
	fmt.Printf(
		"Task Id: %d\nDescription: %s\nStatus: %s\nCreatedAt: %v\nUpdatedAt: %v\n",
		task.Id, task.Description, task.Status, task.CreatedAt, task.UpdatedAt,
	)
	fmt.Println("---------------------------------")
}

func HandleStatuses() string {
	statues := [4]string{"1. DONE", "2. PENDING", "3. SUSPEND", "4. CANCEL"}
	fmt.Println("Plaese Choose the Task Status: ")
	for _, v := range statues {
		fmt.Println(v)
	}
	var choiceStatus int
	var status string

	fmt.Scanln(&choiceStatus)
	switch choiceStatus {
	case 1:
		status = "DONE"
	case 2:
		status = "PENDING"
	case 3:
		status = "SUSPEND"
	case 4:
		status = "CANCEL"
	}

	return status

}

func HandlePagination() {
	page := 1
	pageSize := 5

	for {
		tasks, totalPages := GetTasks(page, pageSize)

		for _, v := range tasks {
			PrintTaskCards(v)
		}

		fmt.Printf("\nPage %d / %d\n", page, totalPages)

		fmt.Println("1. Next Page")
		fmt.Println("2. Prev Page")
		fmt.Println("3. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			if page < totalPages {
				page++
			}
		case 2:
			if page > 1 {
				page--
			}
		case 3:
			return
		}
	}
}

func HandleAddTasks() (string, string) {
	fmt.Println("Plaese Enter the Task Description: ")
	var description string
	fmt.Scanln(&description)

	status := HandleStatuses()

	return status, description
}
