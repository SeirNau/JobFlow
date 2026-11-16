package main

import "testing"

func TestPlan(t *testing.T) {
	if describe("Чай", 2) != "Дело: Чай | Приоритет: 2" {
		t.Fatal("describe")
	}
	if validPriority(0) || !validPriority(1) || !validPriority(2) || !validPriority(3) || validPriority(4) {
		t.Fatal("validPriority")
	}
	if isUrgent(1) || !isUrgent(2) || !isUrgent(3) {
		t.Fatal("isUrgent")
	}
	if totalMinutes(25, 15, 35, 5) != 85 || totalMinutes(40, 20, 30, 5) != 100 {
		t.Fatal("totalMinutes")
	}
	// isUrgent(9) == true, поэтому вызывающий код обязан вызвать validPriority сначала.
	if !isUrgent(9) || validPriority(9) {
		t.Fatal("граница проверки")
	}
}
