package main

import (
	"fmt"
	"database/sql"
	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return sales, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT product, volume, date FROM sales WHERE client = :id", sql.Named("id", client))
	if err != nil {
		return sales, err
	}
	defer rows.Close()
	var (
		product int
		volume int
		date string
	)
    
	for rows.Next() {
		err = rows.Scan(&product, &volume, &date)
		if err != nil {
			return sales, err
		}
		sales = append(sales, Sale {
			product, volume, date,
		})
	}
	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
