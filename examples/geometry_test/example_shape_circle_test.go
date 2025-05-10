package geometry_test

import (
	"fmt"

	"github.com/humble-zh/geometry"
)

func ExampleCircle_Area() {
	c := geometry.NewCircle(5)
	area, err := c.Area()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("圆形面积: %.2f\n", area)
	// Output: 圆形面积: 78.54
}

func ExampleCircle_Perimeter() {
	c := geometry.NewCircle(5)
	perimeter, err := c.Perimeter()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("圆形周长: %.2f\n", perimeter)
	// Output: 圆形周长: 31.42
}
