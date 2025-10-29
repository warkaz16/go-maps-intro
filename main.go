package main

import (
	"fmt"
	"strings"
)

func getVisits(m map[string]int, day string) (int, bool) {
	v, ok := m[day]
	return v, ok
}

func clone(src map[string]int) map[string]int {
	m := make(map[string]int)
	// maps.Copy(m, src)
	for k, v := range src {
		m[k] = v
	}
	return m
}

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	m := map[string]int{}

	for _, v := range words {
		m[v]++
	}

	return m
}

func Add(phones map[string]string, name, number string) {
	phones[name] = number

}

func Get(phones map[string]string, name string) (string, bool) {
	v, ok := phones[name]
	return v, ok
}

func Remove(phones map[string]string, name string) {
	delete(phones, name)
}

func main() {
	// Подзадача 1 — Первое знакомство

	m := map[string]int{
		"Mon": 10,
		"Tue": 20,
		"Wen": 25,
		"Thu": 21,
	}

	fmt.Println("====\\ Подзадача 1 /====")
	fmt.Println(m, len(m))

	// Подзадача 2 — Проверка наличия ключа

	v, ok := m["Mon"]
	fmt.Println("\n====\\ Подзадача 2 /====")
	fmt.Println(v, ok)

	fmt.Println(getVisits(m, "Wen"))

	// Подзадача 3 — Удаление и длина
	fmt.Println("\n====\\ Подзадача 3 /====")

	delete(m, "Tue") // удаление элемента
	fmt.Println(m, len(m))

	delete(m, "Sun") // удаление несуществующего элемента

	//Подзадача 4 — Обход карты

	city := map[string]string{
		"Грозный":     "Беркат Рынок",
		"Урус-Мартан": "Мащен Базар",
		"Аргун":       "Мечеть",
		"Гудермес":    "Мост",
		"Закан Юрт":   "Бурсагов",
	}

	fmt.Println("\n====\\ Подзадача 4 /====")
	for k, v := range city {
		fmt.Println(k, v)
	}

	// Подзадача 5 — assignment to entry in nil map

	m1 := make(map[string]int)
	m1["visits"] = 1
	fmt.Println("\n====\\ Подзадача 5 /====")
	fmt.Println(m1)

	// Подзадача 6 — некорректный тип ключа

	// m2 := make(map[[]int]int) // невозможно создать ключ из несравниваемых значений

	// Подзадача 7 — Карта значений-структур

	type Student struct {
		Name   string
		Active bool
	}

	m2 := map[string]Student{
		"st1": {Name: "Adam", Active: true},
		"st2": {Name: "Zubayr", Active: true},
	}

	new := m2["st2"]
	m2["st2"] = Student{Name: new.Name, Active: false}

	fmt.Println("\n====\\ Подзадача 7 /====")
	fmt.Println(m2)

	// Подзадача 8 — Ссылочная семантика карт

	a := map[string]int{"Mon": 10}
	b := a
	b["Mon"] = 99

	fmt.Println("\n====\\ Подзадача 8 /====")
	fmt.Println(a)
	fmt.Println(b)
	// в обеих картах 99 потому что значение в картах не копируется а напрямую ссылается в источник

	c := clone(b)
	c["Mon"] = 9900 // независимая копия - не меняет оригинал
	fmt.Println("клон -", c)
	fmt.Println("старый -",a)
	// Подзадача 9 — счётчик слов

	fmt.Println("\n====\\ Подзадача 9 /====")
	word := WordCount("go go togo golang")
	fmt.Println(word)

	// Подзадача 10 — простая база телефонов

	phones := map[string]string{}

	Add(phones, "adam", "89380009614")
	Add(phones, "zubayr", "89298884211")
	Add(phones, "otez", "89224567586")
	Add(phones, "jabrail", "89238889090")

	fmt.Println("\n====\\ Подзадача 10 /====")
	for i, v := range phones {
		fmt.Println(i, v)
	}

	fmt.Println("\nполучить номер по имени:")
	fmt.Println(Get(phones, "adam"))

	Remove(phones, "jabrail")

	fmt.Println()
	fmt.Println("карта после удаления:", phones, "\nДлина карты после удаления:", len(phones))
}
