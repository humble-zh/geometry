package geometry

// Shape 定义形状接口
type Shape interface {
	Area() (float64, error)
	Perimeter() (float64, error)
	Name() string
}

// BaseShape 基础形状，包含日志记录器
type BaseShape struct {
	logger Logger
}

// SetLogger 设置日志记录器
func (b *BaseShape) SetLogger(logger Logger) {
	if logger == nil {
		logger = &NoopLogger{}
	}
	b.logger = logger
}

// GetLogger 获取日志记录器
func (b *BaseShape) GetLogger() Logger {
	if b.logger == nil {
		return &NoopLogger{}
	}
	return b.logger
}
