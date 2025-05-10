package geometry

import (
	"errors"
	"math"
)

var ErrNegativeRadius = errors.New("圆形的半径不能为负数")

// Circle 实现圆形形状
type Circle struct {
	BaseShape
	Radius float64
}

// NewCircle 创建带日志功能的圆形
func NewCircle(radius float64, logger ...Logger) *Circle {
	c := &Circle{Radius: radius}
	if len(logger) > 0 {
		c.SetLogger(logger[0])
	}
	return c
}

// Area 计算圆形面积
func (c *Circle) Area() (float64, error) {
	logger := c.GetLogger()
	if c.Radius < 0 {
		logger.Errorf("错误: %v, 半径: %.2f", ErrNegativeRadius, c.Radius)
		return 0, ErrNegativeRadius
	}
	logger.Infof("计算圆形面积: 半径=%.2f, 面积=%.2f", c.Radius, math.Pi*c.Radius*c.Radius)
	return math.Pi * c.Radius * c.Radius, nil
}

// Perimeter 计算圆形周长
func (c *Circle) Perimeter() (float64, error) {
	logger := c.GetLogger()
	if c.Radius < 0 {
		logger.Errorf("错误: %v, 半径: %.2f", ErrNegativeRadius, c.Radius)
		return 0, ErrNegativeRadius
	}
	logger.Infof("计算圆形周长: 半径=%.2f, 周长=%.2f", c.Radius, 2*math.Pi*c.Radius)
	return 2 * math.Pi * c.Radius, nil
}

// Name 返回形状名称
func (c *Circle) Name() string {
	return "圆形"
}
