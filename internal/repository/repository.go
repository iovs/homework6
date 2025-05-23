package repository

import "homework6/internal/model"

var OrderSlice []model.Order // слайс структур model.Order

// Описание подходящего нам интерфейса
type ID interface {
	GetID() int64
}

// функция принимает любые структуры, если они имплементируют (предоставляют) метод GetID() int64
func Store(input ID) {
	switch typed := input.(type) { // с помощью .(type) получаем реальный тип input, по сути снимаем мантию интерфейса

	case model.Order:
		//append в нужный слайс
		OrderSlice = append(OrderSlice, typed)
		// и так далее
	}
}
