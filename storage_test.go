package main

import (
	"reflect"
	"testing"
)

func TestCRUDAndRefusals(t *testing.T) {
	tasks := make(map[string]string)
	if !create(tasks, "a", "Чай") || create(tasks, "a", "Стереть") {
		t.Fatal("create")
	}
	if title, found := get(tasks, "a"); !found || title != "Чай" {
		t.Fatal("duplicate changed entry")
	}
	if update(tasks, "z", "Неизвестно") || !update(tasks, "a", "Улун") {
		t.Fatal("update")
	}
	if title, _ := get(tasks, "a"); title != "Улун" {
		t.Fatal("update missing")
	}
	if remove(tasks, "z") || !remove(tasks, "a") {
		t.Fatal("remove")
	}
	if _, found := get(tasks, "a"); found {
		t.Fatal("get after remove")
	}
}

func TestBudget(t *testing.T) {
	durations := []int{10, 30, 25}
	indices, left, ok := selectByBudget(durations, 50)
	if !ok || !reflect.DeepEqual(indices, []int{0, 1}) || left != 10 {
		t.Fatalf("bad selection: %v %d %t", indices, left, ok)
	}
	before := []int{10, -1, 5}
	indices, left, ok = selectByBudget(before, 50)
	if ok || len(indices) > 0 || left != 50 || !reflect.DeepEqual(before, []int{10, -1, 5}) {
		t.Fatal("negative duration mutated input")
	}
	indices, left, ok = selectByBudget([]int{}, 50)
	if !ok || len(indices) != 0 || left != 50 {
		t.Fatal("empty slice")
	}
	indices, left, ok = selectByBudget([]int{40, 25, 25}, 50)
	if !ok || !reflect.DeepEqual(indices, []int{0}) || left != 10 {
		t.Fatal("greedy example")
	}
	indices, left, ok = selectByBudget([]int{0, 10}, 10)
	if !ok || !reflect.DeepEqual(indices, []int{0, 1}) || left != 0 {
		t.Fatal("zero duration")
	}
}
