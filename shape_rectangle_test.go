package geometry

import (
	"errors"
	"fmt"
	"testing"
)

// TestRectangle_Area 测试矩形面积计算
func TestRectangle_Area(t *testing.T) {
	tests := []struct {
		name    string
		width   float64
		height  float64
		want    float64
		wantErr bool
	}{
		{"正常尺寸", 5, 4, 5 * 4, false},
		{"零宽", 0, 4, 0, false},
		{"零高", 5, 0, 0, false},
		{"零宽高", 0, 0, 0, false},
		{"负宽", -5, 4, 0, true},
		{"负高", 5, -4, 0, true},
		{"负宽高", -5, -4, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRectangle(tt.width, tt.height)

			// 测试无日志情况
			area, err := r.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && area != tt.want {
				t.Errorf("Rectangle.Area() = %v, want %v", area, tt.want)
			}

			// 测试带日志情况
			mockLogger := &MockLogger{}
			r.SetLogger(mockLogger)
			area, err = r.Area()

			if tt.wantErr {
				if !mockLogger.HasErrorMessage("矩形的宽和高不能为负数") {
					t.Errorf("Rectangle.Area() 错误时未记录正确的错误日志")
				}
			} else {
				if !mockLogger.HasInfoMessage("计算矩形面积") {
					t.Errorf("Rectangle.Area() 成功时未记录信息日志")
				}
			}
		})
	}
}

// TestRectangle_Perimeter 测试矩形周长计算
func TestRectangle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		width   float64
		height  float64
		want    float64
		wantErr bool
	}{
		{"正常尺寸", 5, 4, 2 * (5 + 4), false},
		{"零宽", 0, 4, 2 * (0 + 4), false},
		{"零高", 5, 0, 2 * (5 + 0), false},
		{"零宽高", 0, 0, 0, false},
		{"负宽", -5, 4, 0, true},
		{"负高", 5, -4, 0, true},
		{"负宽高", -5, -4, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRectangle(tt.width, tt.height)

			// 测试无日志情况
			perimeter, err := r.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Rectangle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && perimeter != tt.want {
				t.Errorf("Rectangle.Perimeter() = %v, want %v", perimeter, tt.want)
			}

			// 测试带日志情况
			mockLogger := &MockLogger{}
			r.SetLogger(mockLogger)
			perimeter, err = r.Perimeter()

			if tt.wantErr {
				if !mockLogger.HasErrorMessage("矩形的宽和高不能为负数") {
					t.Errorf("Rectangle.Perimeter() 错误时未记录正确的错误日志")
				}
			} else {
				if !mockLogger.HasInfoMessage("计算矩形周长") {
					t.Errorf("Rectangle.Perimeter() 成功时未记录信息日志")
				}
			}
		})
	}
}

// TestRectangle_Name 测试矩形名称
func TestRectangle_Name(t *testing.T) {
	r := NewRectangle(5, 4)
	if got := r.Name(); got != "矩形" {
		t.Errorf("Rectangle.Name() = %q, want %q", got, "矩形")
	}
}

// ExampleRectangle_Area 矩形面积计算示例
func ExampleRectangle_Area() {
	r := NewRectangle(5, 4)
	area, err := r.Area()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("矩形面积: %.2f\n", area)
	// Output: 矩形面积: 20.00
}

// TestRectangle_Area_NegativeWidth 测试负宽度的情况
func TestRectangle_Area_NegativeWidth(t *testing.T) {
	r := NewRectangle(-5, 4)
	area, err := r.Area()

	if err == nil {
		t.Errorf("Rectangle.Area() 期望错误，但未返回错误")
		return
	}

	if !errors.Is(err, ErrInvalidDimensions) {
		t.Errorf("Rectangle.Area() 错误类型不匹配，got %v, want %v", err, ErrInvalidDimensions)
	}

	if area != 0 {
		t.Errorf("Rectangle.Area() 负宽度时面积应为 0，got %v", area)
	}
}

// TestRectangle_Area_NegativeHeight 测试负高度的情况
func TestRectangle_Area_NegativeHeight(t *testing.T) {
	r := NewRectangle(5, -4)
	area, err := r.Area()

	if err == nil {
		t.Errorf("Rectangle.Area() 期望错误，但未返回错误")
		return
	}

	if !errors.Is(err, ErrInvalidDimensions) {
		t.Errorf("Rectangle.Area() 错误类型不匹配，got %v, want %v", err, ErrInvalidDimensions)
	}

	if area != 0 {
		t.Errorf("Rectangle.Area() 负高度时面积应为 0，got %v", area)
	}
}

// TestRectangle_Perimeter_NegativeDimensions 测试负尺寸的周长计算
func TestRectangle_Perimeter_NegativeDimensions(t *testing.T) {
	r := NewRectangle(-5, -4)
	perimeter, err := r.Perimeter()

	if err == nil {
		t.Errorf("Rectangle.Perimeter() 期望错误，但未返回错误")
		return
	}

	if !errors.Is(err, ErrInvalidDimensions) {
		t.Errorf("Rectangle.Perimeter() 错误类型不匹配，got %v, want %v", err, ErrInvalidDimensions)
	}

	if perimeter != 0 {
		t.Errorf("Rectangle.Perimeter() 负尺寸时周长应为 0，got %v", perimeter)
	}
}

// TestRectangle_Area_WithLogrus 测试与 logrus 集成的情况
// func TestRectangle_Area_WithLogrus(t *testing.T) {
// 	logger := logrus.New()
// 	logger.SetLevel(logrus.InfoLevel)
// 	logger.Out = &bytes.Buffer{}
// 	logger.SetFormatter(&logrus.TextFormatter{})

// 	r := NewRectangle(5, 4, logger)
// 	area, err := r.Area()

// 	if err != nil {
// 		t.Errorf("Rectangle.Area() error = %v", err)
// 	}

// 	if area != 5*4 {
// 		t.Errorf("Rectangle.Area() = %v, want %v", area, 5*4)
// 	}
// }
