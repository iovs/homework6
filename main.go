package main

import (
	"fmt"
	"homework6/internal/model"
)

//type Order struct {
//	ID             int64
//	Quantity       int64
//	DateTime       time.Time
//	Price          int64
//	Status         string
//	DeliveryAdress []string
//}

// Интерфейс Orders
type Orders interface {
	Add() string
}

//type Adminka struct {
//	Import string
//	File   string
//}
//func (a model.Adminka) add() string {
//	return "I am the Adminka"
//}

func main() {
	//	service.Func1()
	Adminka := model.Adminka{}
	add(Adminka)
}

func add(orders Orders) {
	fmt.Println(orders.Add())
}

//fmt.Println(NewFunc1)
//	println(func1.NewImport{s, d})
//
//	for i=0; i<100; i++ {
//		importfile := model.Adminka {...}
//		saveorder := model.Order {...}
//		addproduct := model.Products {...}
//		addproductcategory := model.Category {...}
//		adduser := model.User {...}
//		repository.CreateItem1(context, importfile)
//		repository.CreateItem1(context, saveorder)
//		repository.CreateItem1(context, addproduct)
//		repository.CreateItem1(context, addproductcategory)
//		repository.CreateItem1(context, adduser)

//println("Вывод чего-то )",  ) // вывод в лог информации о слайсах
