package geometry

import "errors"

// Measurer 测量器，用于测量形状
type Measurer struct{}

// NewMeasurer 创建测量器
func NewMeasurer() *Measurer {
	return &Measurer{}
}

// MeasureShape 测量形状的面积和周长
func (m *Measurer) MeasureShape(s Shape) (area, perimeter float64, err error) {
	if s == nil {
		return 0, 0, errors.New("形状不能为空")
	}

	area, err = s.Area()
	if err != nil {
		return 0, 0, err
	}

	perimeter, err = s.Perimeter()
	if err != nil {
		return area, 0, err
	}

	return area, perimeter, nil
}
