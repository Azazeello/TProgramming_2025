package main

import "fmt"

type Fox struct {
	Name       string
	Age        int
	IsSleeping bool
}

func NewFox(name string, age int) *Fox {
	return &Fox{
		Name:       name,
		Age:        age,
		IsSleeping: false,
	}
}

// ASCII представление лисицы
func (f *Fox) GetView() string {
	return `
	  /\_/\
	 ( o.o )
	  > ^ <
	`
}

// Метод заставить лисицу спать
func (f *Fox) Sleep() {
	f.IsSleeping = true
}

// Метод разбудить лисицу
func (f *Fox) WakeUp() {
	f.IsSleeping = false
}

// Метод проверки состояния
func (f *Fox) Status() string {
	if f.IsSleeping {
		return fmt.Sprintf("%s спит (zzz)", f.Name)
	}
	return fmt.Sprintf("%s бодрствует!", f.Name)
}

func main() {

	fox := NewFox("Рыжик", 3)

	fmt.Println(fox.GetView())
	fmt.Println(fox.Status())

	fox.Sleep()
	fmt.Println(fox.Status())

	fox.WakeUp()
	fmt.Println(fox.Status())
}