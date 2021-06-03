package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"

	"time"
)

type Product struct {
	Name     string
	Category string
	Price    int
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	category := [4]string{"fashion", "electronic", "sport", "food"}
	products := [20]Product{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(products); i++ {
		products[i] = Product{
			fmt.Sprintf("%s %d", "Product", i),
			category[rand.Intn(len(category))],
			randomInt(100, 200),
		}
	}
	for _, product := range products {
		fmt.Println(product)
	}

	sort.Slice(products[:], func(i, j int) bool {
		return products[i].Price > products[j].Price
	})
	for _, product := range products[:5] {
		fmt.Println(product)
	}

	myMap := make(map[string]int)
	for _, product := range products {
		myMap[product.Category]++
	}
	fmt.Println(myMap)
	//2.Gõ vào từ bàn phím tên một sản phẩm, trả về danh sách tìm được
	fmt.Println("Nhap vao ten product")
	scannerProductName := bufio.NewScanner(os.Stdin)
	scannerProductName.Scan()

	inputProductName := scannerProductName.Text()

	for _, product := range products {
		if product.Name == inputProductName {
			fmt.Println(product)

		}
	}
	//3.Gõ vào category, trả về danh sách tất cả các sản phẩm trong category đó
	fmt.Println("Nhap vao ten Category")
	scannerCategory := bufio.NewScanner(os.Stdin)
	scannerCategory.Scan()

	inputCategory := scannerCategory.Text()

	for _, product := range products {
		if product.Category == inputCategory {
			fmt.Println(product)
			break
		}
	}
	//4.Gõ vào price min, và price max, tìm tất cả các sản phẩm có giá trong dải min đến max
	fmt.Println("Nhap vao gia tri min")
	scannerMin := bufio.NewScanner(os.Stdin)
	scannerMin.Scan()

	inputMin := scannerMin.Text()
	min, _ := strconv.Atoi(inputMin)

	fmt.Println("Nhap vao gia tri max")
	scannerMax := bufio.NewScanner(os.Stdin)
	scannerMax.Scan()

	inputMax := scannerMax.Text()
	max, _ := strconv.Atoi(inputMax)

	for _, product := range products {
		if product.Price > min && product.Price < max {
			fmt.Println(product)

		}
	}

}

//nhập vào danh sách sản phẩm
