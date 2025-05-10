package geometry_test

import (
	"fmt"

	"github.com/humble-zh/geometry"
)

func ExampleMeasurer_MeasureShape() {
	measurer := geometry.NewMeasurer()

	circle := geometry.NewCircle(5)
	circleArea, circlePerimeter, err := measurer.MeasureShape(circle)
	if err == nil {
		fmt.Printf("圆形 - 测量 失败: %v\n", err)
	} else {
		fmt.Printf("圆形 - 面积: %.2f, 周长: %.2f\n", circleArea, circlePerimeter)
	}

	rect := geometry.NewRectangle(4, 6)
	rectArea, rectPerimeter, err := measurer.MeasureShape(rect)
	if err != nil {
		fmt.Printf("矩形(%f,%f) - 测量 失败: %v\n", rect.Width, rect.Height, err)
	} else {
		fmt.Printf("矩形(%f,%f) - 面积: %.2f, 周长: %.2f\n", rect.Width, rect.Height, rectArea, rectPerimeter)
	}
}
