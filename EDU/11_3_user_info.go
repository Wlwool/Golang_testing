package main

import (
	"fmt"
	"strings"
)

// User представляет пользователя
type User struct {
	ID      int
	Name    string
	Email   string
	Phone   string
	Address Address
	Cart    []CartItem
}

// Address представляет адрес пользователя
type Address struct {
	Street     string
	City       string
	PostalCode string
}

// CartItem представляет элемент в корзине
type CartItem struct {
	Product  Product
	Quantity int
}

// Product представляет продукт в корзине
type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Category    string
	Brand       string
	Rating      float64
	Reviews     int
}

func printInfo(user User) {
	fmt.Printf("Покупатель %s. Телефон: %s. Адрес: г. %s, %s.\n", 
	user.Name, 
	user.Phone, 
	user.Address.City, 
	user.Address.Street)
	
	// проверка на покупателя электроники
	isBuyer := false
	for _, item := range user.Cart {
		if item.Product.Category == "Электроника"{
			isBuyer = true
			break
		}
	}

	// статус поупателя
	status := "не является"
	if isBuyer {
		status = "является"
	}
	fmt.Printf("Пользователь %s покупателем электроники.\n", status)

	// поиск товаров с ценой более 10000
	var expensiveGoodsList []string
	for _, item := range user.Cart {
		if item.Product.Price >= 10000{
			expensiveGoodsList = append(expensiveGoodsList, item.Product.Name)
		}
	}

	// формирование списка товаров
	goodsList := "нет"
	if len(expensiveGoodsList) > 0 {
		goodsList = strings.Join(expensiveGoodsList, ", ")
	}
	fmt.Printf("Товары в корзине, где цена 10000 и более: %s.\n", goodsList)

	total := 0
	for _, i := range user.Cart {
		total += i.Product.Price * i.Quantity
	}
	fmt.Printf("Общая сумма покупки: %d руб.\n", total)

}

func main() {
	user := User{

        ID:    1,

        Name:  "Иван Петров",

        Email: "ivan.petrov@example.com",

        Phone: "+7 999 123-45-67",

        Address: Address{

            Street:     "Улица Ленина",

            City:       "Москва",

            PostalCode: "101000",

        },

        Cart: []CartItem{

            {

                Product: Product{

                    ID:          1,

                    Name:        "Ноутбук",

                    Description: "Мощный ноутбук для работы и игр",

                    Price:       59990,

                    Category:    "Электроника",

                    Brand:       "Brand A",

                    Rating:      4.5,

                    Reviews:     120,

                },

                Quantity: 1,

            },

            {

                Product: Product{

                    ID:          2,

                    Name:        "Смартфон",

                    Description: "Современный смартфон с отличной камерой",

                    Price:       29990,

                    Category:    "Электроника",

                    Brand:       "Brand B",

                    Rating:      4.7,

                    Reviews:     200,

                },

                Quantity: 2,

            },

            {

                Product: Product{

                    ID:          3,

                    Name:        "Наушники",

                    Description: "Беспроводные наушники с шумоподавлением",

                    Price:       7990,

                    Category:    "Аудио",

                    Brand:       "Brand C",

                    Rating:      4.3,

                    Reviews:     80,

                },

                Quantity: 1,

            },

        },

    }

	printInfo(user)

}