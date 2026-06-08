package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func FindAll() []Tasks {
	path := GetFilePath()
	data, err := os.ReadFile(path)
	check(err)
	var tasks []Tasks
	json.Unmarshal(data, &tasks)

	return tasks
}

func Save(tasks []Tasks) {
	newData, errr := json.Marshal(tasks)
	check(errr)
	os.WriteFile(GetFilePath(), newData, os.ModeAppend)
}

func FindOne(id int) (Tasks, int) {
	tasks := FindAll()
	for i, v := range tasks {
		if v.Id == id {
			return v, i
		}
	}

	return Tasks{}, -1
}

func GetFilePath() string {
	return filepath.Join(".", "tasks.json")
}
