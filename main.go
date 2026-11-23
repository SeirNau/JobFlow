package main

import (
	"fmt"
	"jobflow/pkg"
)

func main() {
	firstTitle := "Убраться в квартире"
	secondTitle := "Купить продукты"
	thirdTitle := "Повторить Go"

	firstMinutes := 25
	secondMinutes := 15
	thirdMinutes := 35
	const breakMinutes = 5
	availableMinutes := 90

	title := firstTitle
	priority := 2
	done := false
	reminders := 3

	workMinutes := firstMinutes + secondMinutes + thirdMinutes
	breaksTotal := 2 * breakMinutes
	planMinutes := totalMinutes(firstMinutes, secondMinutes, thirdMinutes, breakMinutes)

	fmt.Println(firstTitle, firstMinutes, "мин")
	fmt.Println(secondTitle, secondMinutes, "мин")
	fmt.Println(thirdTitle, thirdMinutes, "мин")
	fmt.Println("Работа, мин:", workMinutes)
	fmt.Println("Перерывы, мин:", breaksTotal)
	fmt.Printf("Итого: %d ч %d мин\n", planMinutes/60, planMinutes%60)
	fmt.Println("Всего минут:", planMinutes)
	if planMinutes <= availableMinutes {
		fmt.Println("План помещается")
	} else {
		fmt.Println("Перегрузка на", planMinutes-availableMinutes, "минут")
	}

	if title == "" {
		fmt.Println("Ошибка: пустое название дела")
		return
	}
	if !validPriority(priority) {
		fmt.Println("Ошибка: приоритет должен быть от 1 до 3")
		return
	}
	if reminders < 0 || reminders > 3 {
		fmt.Println("Ошибка: число напоминаний должно быть от 0 до 3")
		return
	}

	fmt.Println(describe(title, priority))
	if done {
		fmt.Println("Уже выполнено")
		return
	}
	if isUrgent(priority) {
		fmt.Println("Срочно")
	} else {
		fmt.Println("Обычное дело")
	}
	printReminders(title, reminders)

	runCollectionsDemo()

	task := pkg.Task{
		ID:       "a",
		Title:    "Повторить Go",
		Priority: 2,
	}

	fmt.Println("До:", task.Done)

	task.Complete()
	fmt.Println("После первого:", task.Done)

	task.Complete()
	fmt.Println("После второго:", task.Done)

	tasksByID := map[string]pkg.Task{
		"a": {
			ID:       "a",
			Title:    "Повторить Go",
			Priority: 2,
		},
	}

	copyTask, found := tasksByID["a"]

	if found {
		copyTask.Complete()

		fmt.Println("Копия:", copyTask.Done)
		fmt.Println("Хранилище до записи:", tasksByID["a"].Done)

		tasksByID["a"] = copyTask

		fmt.Println("Хранилище после записи:", tasksByID["a"].Done)
	}
	fmt.Println("\nПрактика 06.2 — CRUD")

	tasks06 := make(map[string]pkg.Task)

	fmt.Println("Создать a:", createTask(tasks06, pkg.Task{
		ID: "a", Title: "Повторить Go", Priority: 2,
	}))

	fmt.Println("Повторно создать a:", createTask(tasks06, pkg.Task{
		ID: "a", Title: "Дубликат", Priority: 2,
	}))

	local06, found06 := getTask(tasks06, "a")
	fmt.Println("Прочитать a:", local06.Title, found06)

	local06.Title = "Изменена копия"
	stored06, _ := getTask(tasks06, "a")
	fmt.Println("Без Update:", stored06.Title)

	fmt.Println("Изменить a:", updateTask(tasks06, "a", "Прочитать урок 6"))

	updated06, _ := getTask(tasks06, "a")
	fmt.Println("После Update:", updated06)

	fmt.Println("Завершить a:", completeTask(tasks06, "a"))
	fmt.Println("Завершить z:", completeTask(tasks06, "z"))

	completed06, _ := getTask(tasks06, "a")
	fmt.Println("После Complete:", completed06)

	fmt.Println("Удалить a:", removeTask(tasks06, "a"))
	_, exists06 := getTask(tasks06, "a")
	fmt.Println("После удаления найдена:", exists06)

	fmt.Println("\nПрактика 06.4 — бюджет задач")

	budgetTasks := map[string]pkg.Task{
		"a": {ID: "a", Title: "Повторить Go", Priority: 2, EstimateMinutes: 25},
		"b": {ID: "b", Title: "Купить продукты", Priority: 1, Done: true, EstimateMinutes: 15},
		"c": {ID: "c", Title: "Прочитать конспект", Priority: 2, EstimateMinutes: 30},
		"d": {ID: "d", Title: "Сделать заметки", Priority: 1, EstimateMinutes: 10},
	}

	budget := 40
	remaining := budget

	order := []string{"a", "b", "c", "d"}

	for _, id := range order {
		task := budgetTasks[id]

		if task.Done || task.EstimateMinutes <= 0 {
			continue
		}

		if task.EstimateMinutes > remaining {
			continue
		}

		fmt.Println(task.Title, task.EstimateMinutes, "мин")
		remaining -= task.EstimateMinutes
	}

	fmt.Println("Осталось минут:", remaining)
}
