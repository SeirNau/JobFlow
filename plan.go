package main

import "fmt"

// describe получает проверенные данные и возвращает текст карточки без печати.
func describe(title string, priority int) string {
	return fmt.Sprintf("Дело: %s | Приоритет: %d", title, priority)
}

// validPriority принимает любое int и проверяет диапазон 1..3.
func validPriority(priority int) bool {
	return priority >= 1 && priority <= 3
}

// isUrgent требует предварительной проверки validPriority.
func isUrgent(priority int) bool {
	return priority >= 2
}

// totalMinutes требует неотрицательных длительностей; учитывает два перерыва.
func totalMinutes(first, second, third, breakMinutes int) int {
	return first + second + third + 2*breakMinutes
}

// printReminders печатает ровно count строк; вызывающий код проверяет count и done.
func printReminders(title string, count int) {
	for number := 1; number <= count; number++ {
		fmt.Printf("Напоминание %d: %s\n", number, title)
	}
}
