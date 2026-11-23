package main

import "jobflow/pkg"

func createTask(tasks map[string]pkg.Task, task pkg.Task) bool {
	if tasks == nil || task.ID == "" || task.Title == "" ||
		task.Priority < 1 || task.Priority > 3 {
		return false
	}

	if _, exists := tasks[task.ID]; exists {
		return false
	}

	tasks[task.ID] = task
	return true
}

func getTask(tasks map[string]pkg.Task, id string) (pkg.Task, bool) {
	task, found := tasks[id]
	return task, found
}

func updateTask(tasks map[string]pkg.Task, id, title string) bool {
	if title == "" {
		return false
	}

	task, found := tasks[id]
	if !found {
		return false
	}

	task.Title = title
	tasks[id] = task
	return true
}

func removeTask(tasks map[string]pkg.Task, id string) bool {
	if _, found := tasks[id]; !found {
		return false
	}

	delete(tasks, id)
	return true
}

func completeTask(tasks map[string]pkg.Task, id string) bool {
	task, found := tasks[id]
	if !found {
		return false
	}

	task.Complete()
	tasks[id] = task
	return true
}
