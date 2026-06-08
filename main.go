package main

import "fmt"

func main() {
	for {
		fmt.Println("Please Choose and option: ")
		options := [5]string{"1. Get tasks", "2. Add Task", "3. Delete Task", "4. Update Task", "5. Exit"}

		for _, v := range options {
			fmt.Println(v)
		}
		var choice int
		fmt.Print("Enter you choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			HandlePagination()
		case 2:
			Add()
		case 3:
			HandleDelete()
		case 4:
			Update()
		case 5:
			fmt.Println("Have a Nice Day!")
			return
		}
	}

}

func GetTasks(page int, pageSize int) ([]Tasks, int) {

	allTasks := FindAll()

	total := len(allTasks)

	totalPages := (total + pageSize - 1) / pageSize

	offset := (page - 1) * pageSize

	if offset >= total {
		return []Tasks{}, totalPages
	}

	end := offset + pageSize
	if end > total {
		end = total
	}

	return allTasks[offset:end], totalPages
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
