package geometry

// Logger 定义日志接口
type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Errorf(format string, v ...interface{})
}

// NoopLogger 实现空日志记录器
type NoopLogger struct{}

func (n *NoopLogger) Debugf(format string, v ...interface{}) {}
func (n *NoopLogger) Infof(format string, v ...interface{})  {}
func (n *NoopLogger) Errorf(format string, v ...interface{}) {}
