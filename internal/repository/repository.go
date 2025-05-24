package repository

import (
	"fmt"
	"homework6/internal/model"
)

var adminka []model.Adminka
var order []model.Order
var product []model.Product

func AddData(data model.Orders) {
	switch v := data.(type) {
	case model.Order:
		order = append(order, v)
		fmt.Println("add Order")
	case model.Adminka:
		adminka = append(adminka, v)
		fmt.Println("add Adminka")
	default:
		fmt.Printf("Не определенный тип данных: %T/n", v)
	}

}

//var OrderSlice []model.Order // слайс структур model.Order

// Описание подходящего нам интерфейса
//type ID interface {
//	GetID() int64
//}

// функция принимает любые структуры, если они имплементируют (предоставляют) метод GetID() int64
//func Store(input ID) {
//	switch typed := input.(type) { // с помощью .(type) получаем реальный тип input, по сути снимаем мантию интерфейса

//	case model.Order:
//
// append в нужный слайс
//
//	OrderSlice = append(OrderSlice, typed)
//
// и так далее
//
//	}
//}
