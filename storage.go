package main

import (
	"fmt"
	"sort"
)

func create(tasks map[string]string, id, title string) bool {
	if _, exists := tasks[id]; exists {
		return false
	}
	tasks[id] = title
	return true
}

func get(tasks map[string]string, id string) (string, bool) {
	title, found := tasks[id]
	return title, found
}

func update(tasks map[string]string, id, title string) bool {
	if _, exists := tasks[id]; !exists {
		return false
	}
	tasks[id] = title
	return true
}

func remove(tasks map[string]string, id string) bool {
	if _, exists := tasks[id]; !exists {
		return false
	}
	delete(tasks, id)
	return true
}

func printSorted(tasks map[string]string) {
	keys := make([]string, 0, len(tasks))
	for key := range tasks {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Print("Состояние: [")
	for i, key := range keys {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%s=%q", key, tasks[key])
	}
	fmt.Println("]")
}

func runCollectionsDemo() {
	fmt.Println("\nПрактика 05: выбор дел по порядку")
	durations := []int{10, 30, 25}
	indices, left, ok := selectByBudget(durations, 50)
	fmt.Printf("Длительности: %v; выбраны индексы: %v; остаток: %d; успех: %t\n", durations, indices, left, ok)
	fmt.Println("Пустой список:")
	emptyIndices, emptyLeft, emptyOK := selectByBudget([]int{}, 50)
	fmt.Printf("Выбрано: %v; остаток: %d; успех: %t\n", emptyIndices, emptyLeft, emptyOK)
	fmt.Println("Отрицательная длительность:")
	bad := []int{10, -2, 5}
	_, badLeft, badOK := selectByBudget(bad, 50)
	fmt.Printf("Успех: %t; остаток: %d; исходные данные: %v\n", badOK, badLeft, bad)

	fmt.Println("\nПрактика 05: CRUD по ID")
	tasks := make(map[string]string)
	fmt.Println("Создать a:", create(tasks, "a", "Купить чай"))
	printSorted(tasks)
	title, found := get(tasks, "a")
	fmt.Printf("Прочитать a: название=%q, найдено=%t\n", title, found)
	printSorted(tasks)
	fmt.Println("Повторно создать a:", create(tasks, "a", "Другое название"))
	printSorted(tasks)
	fmt.Println("Изменить z:", update(tasks, "z", "Нет такого дела"))
	printSorted(tasks)
	fmt.Println("Изменить a:", update(tasks, "a", "Заварить улун"))
	printSorted(tasks)
	fmt.Println("Удалить a:", remove(tasks, "a"))
	printSorted(tasks)
	title, found = get(tasks, "a")
	fmt.Printf("Прочитать a: название=%q, найдено=%t\n", title, found)
	printSorted(tasks)
}
