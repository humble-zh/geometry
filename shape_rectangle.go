package geometry

import "errors"

var ErrInvalidDimensions = errors.New("矩形的宽和高不能为负数")

// Rectangle 实现矩形形状
type Rectangle struct {
	BaseShape
	Width  float64
	Height float64
}

// NewRectangle 创建带日志功能的矩形
func NewRectangle(width, height float64, logger ...Logger) *Rectangle {
	r := &Rectangle{Width: width, Height: height}
	if len(logger) > 0 {
		r.SetLogger(logger[0])
	}
	return r
}

// Area 计算矩形面积
func (r *Rectangle) Area() (float64, error) {
	logger := r.GetLogger()
	if r.Width < 0 || r.Height < 0 {
		logger.Errorf("错误: %v, 宽度: %.2f, 高度: %.2f", ErrInvalidDimensions, r.Width, r.Height)
		return 0, ErrInvalidDimensions
	}
	logger.Infof("计算矩形面积: 宽度=%.2f, 高度=%.2f, 面积=%.2f", r.Width, r.Height, r.Width*r.Height)
	return r.Width * r.Height, nil
}

// Perimeter 计算矩形周长
func (r *Rectangle) Perimeter() (float64, error) {
	logger := r.GetLogger()
	if r.Width < 0 || r.Height < 0 {
		logger.Errorf("错误: %v, 宽度: %.2f, 高度: %.2f", ErrInvalidDimensions, r.Width, r.Height)
		return 0, ErrInvalidDimensions
	}
	logger.Infof("计算矩形周长: 宽度=%.2f, 高度=%.2f, 周长=%.2f", r.Width, r.Height, 2*(r.Width+r.Height))
	return 2 * (r.Width + r.Height), nil
}

// Name 返回形状名称
func (r *Rectangle) Name() string {
	return "矩形"
}
