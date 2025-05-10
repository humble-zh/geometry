// shape_circle_test.go
package geometry

import (
	"errors"
	"fmt"
	"math"
	"strings"
	"testing"
)

// MockLogger 用于测试的模拟日志记录器
type MockLogger struct {
	DebugMessages []string
	InfoMessages  []string
	ErrorMessages []string
}

func (m *MockLogger) Debugf(format string, v ...interface{}) {
	m.DebugMessages = append(m.DebugMessages, fmt.Sprintf(format, v...))
}
func (m *MockLogger) Infof(format string, v ...interface{}) {
	m.InfoMessages = append(m.InfoMessages, fmt.Sprintf(format, v...))
}
func (m *MockLogger) Errorf(format string, v ...interface{}) {
	m.ErrorMessages = append(m.ErrorMessages, fmt.Sprintf(format, v...))
}
func (m *MockLogger) HasInfoMessage(msg string) bool {
	for _, m := range m.InfoMessages {
		if strings.Contains(m, msg) {
			return true
		}
	}
	return false
}
func (m *MockLogger) HasErrorMessage(msg string) bool {
	for _, m := range m.ErrorMessages {
		if strings.Contains(m, msg) {
			return true
		}
	}
	return false
}

func TestCircle_Area(t *testing.T) {
	tests := []struct {
		name    string
		radius  float64
		want    float64
		wantErr bool
	}{
		{"正常半径", 5, math.Pi * 5 * 5, false},
		{"零半径", 0, 0, false},
		{"负半径", -3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCircle(tt.radius)

			// 测试无日志情况
			area, err := c.Area()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Area() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && area != tt.want {
				t.Errorf("Circle.Area() = %v, want %v", area, tt.want)
			}

			// 测试带日志情况
			mockLogger := &MockLogger{}
			c.SetLogger(mockLogger)
			area, err = c.Area()

			if tt.wantErr {
				if !mockLogger.HasErrorMessage("圆形的半径不能为负数") {
					t.Errorf("Circle.Area() 错误时未记录正确的错误日志")
				}
			} else {
				if !mockLogger.HasInfoMessage("计算圆形面积") {
					t.Errorf("Circle.Area() 成功时未记录信息日志")
				}
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	tests := []struct {
		name    string
		radius  float64
		want    float64
		wantErr bool
	}{
		{"正常半径", 5, 2 * math.Pi * 5, false},
		{"零半径", 0, 0, false},
		{"负半径", -3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCircle(tt.radius)

			// 测试无日志情况
			perimeter, err := c.Perimeter()
			if (err != nil) != tt.wantErr {
				t.Errorf("Circle.Perimeter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && perimeter != tt.want {
				t.Errorf("Circle.Perimeter() = %v, want %v", perimeter, tt.want)
			}

			// 测试带日志情况
			mockLogger := &MockLogger{}
			c.SetLogger(mockLogger)
			perimeter, err = c.Perimeter()

			if tt.wantErr {
				if !mockLogger.HasErrorMessage("圆形的半径不能为负数") {
					t.Errorf("Circle.Perimeter() 错误时未记录正确的错误日志")
				}
			} else {
				if !mockLogger.HasInfoMessage("计算圆形周长") {
					t.Errorf("Circle.Perimeter() 成功时未记录信息日志")
				}
			}
		})
	}
}

func TestCircle_Name(t *testing.T) {
	c := NewCircle(5)
	if got := c.Name(); got != "圆形" {
		t.Errorf("Circle.Name() = %q, want %q", got, "圆形")
	}
}

func ExampleCircle_Area() {
	c := NewCircle(5)
	area, err := c.Area()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("圆形面积: %.2f\n", area)
	// Output: 圆形面积: 78.54
}

// func TestCircle_Area_WithLogrus(t *testing.T) {
// 	logger := logrus.New()
// 	logger.SetLevel(logrus.InfoLevel)
// 	logger.Out = &bytes.Buffer{}                 // 使用缓冲区代替 nil
// 	logger.SetFormatter(&logrus.TextFormatter{}) // 显式设置格式化器

// 	c := NewCircle(5, logger)
// 	area, err := c.Area()

// 	if err != nil {
// 		t.Errorf("Circle.Area() error = %v", err)
// 	}

// 	if area != math.Pi*5*5 {
// 		t.Errorf("Circle.Area() = %v, want %v", area, math.Pi*5*5)
// 	}
// }

func TestCircle_Area_NegativeRadius(t *testing.T) {
	c := NewCircle(-5)
	area, err := c.Area()

	if err == nil {
		t.Errorf("Circle.Area() 期望错误，但未返回错误")
		return
	}

	// 使用 errors.Is 比较错误类型
	if !errors.Is(err, ErrNegativeRadius) {
		t.Errorf("Circle.Area() 错误类型不匹配，got %v, want %v", err, ErrNegativeRadius)
	}

	if area != 0 {
		t.Errorf("Circle.Area() 负数半径时面积应为 0，got %v", area)
	}
}
