package main

import "fmt"

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
}
