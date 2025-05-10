package geometry_test

import (
	"fmt"

	"github.com/humble-zh/geometry"
)

func ExampleRectangle_Area() {
	r := geometry.NewRectangle(5, 4)
	area, err := r.Area()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("矩形面积: %.2f\n", area)
	// Output: 矩形面积: 20.00
}

func ExampleRectangle_Perimeter() {
	r := geometry.NewRectangle(5, 4)
	perimeter, err := r.Perimeter()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("矩形周长: %.2f\n", perimeter)
	// Output: 矩形周长: 18.00
}
