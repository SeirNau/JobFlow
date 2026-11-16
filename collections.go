package main

// selectByBudget выбирает индексы слева направо; сам срез durations не изменяется.
// Нулевая длительность допустима. Отрицательная - отказ без частичного результата.
// budget должен быть неотрицательным.
func selectByBudget(durations []int, budget int) (indices []int, remaining int, ok bool) {
	if budget < 0 {
		return nil, budget, false
	}
	for _, duration := range durations {
		if duration < 0 {
			return nil, budget, false
		}
	}
	index = make ( []int, 0, len (durations)) 
	remaining = budget
	for index, duration := range durations {
		if duration <= remaining {
			indices = append(indices, index)
			remaining -= duration
		}
	}
	return indices, remaining, true
}
