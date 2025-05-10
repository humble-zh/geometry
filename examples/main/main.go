package main

import (
	"fmt"
	"log"

	"github.com/humble-zh/geometry"
	// "github.com/sirupsen/logrus"
)

// StdLoggerWrapper 包装标准库 log 以实现 geometry.Logger 接口
type StdLoggerWrapper struct {
	logger *log.Logger
}

func (w *StdLoggerWrapper) Debugf(format string, v ...interface{}) {
	w.logger.Printf("[DEBUG] "+format, v...)
}
func (w *StdLoggerWrapper) Infof(format string, v ...interface{}) {
	w.logger.Printf("[INFO] "+format, v...)
}
func (w *StdLoggerWrapper) Errorf(format string, v ...interface{}) {
	w.logger.Printf("[ERROR] "+format, v...)
}

func main() {
	// 创建标准库 logger
	stdLogger := log.New(log.Writer(), "", log.LstdFlags)
	stdWrapper := &StdLoggerWrapper{logger: stdLogger}

	// logrusLogger := logrus.New() // 创建 logrus logger
	// logrusLogger.SetLevel(logrus.InfoLevel)

	// 情况1: 创建矩形时不传入 logger（使用默认 NoopLogger）
	rect1 := geometry.NewRectangle(4, 6)

	// 情况2: 创建矩形时传入标准库 logger
	rect2 := geometry.NewRectangle(5, 7, stdWrapper)

	// 情况3: 创建矩形时传入 logrus logger
	// rect3 := geometry.NewRectangle(3, 8, logrusLogger)

	// 情况4: 创建圆形时不传入 logger
	circle1 := geometry.NewCircle(5)

	// 情况5: 创建圆形时传入 logrus logger
	// circle2 := geometry.NewCircle(3, logrusLogger)

	// 创建测量器（测量器本身不需要 logger）
	measurer := geometry.NewMeasurer()

	// 测量各种形状
	measureShape(measurer, rect1, "矩形1（无日志）")
	measureShape(measurer, rect2, "矩形2（标准库日志）")
	// measureShape(measurer, rect3, "矩形3（logrus日志）")
	measureShape(measurer, circle1, "圆形1（无日志）")
	// measureShape(measurer, circle2, "圆形2（logrus日志）")
}

func measureShape(measurer *geometry.Measurer, shape geometry.Shape, name string) {
	area, perimeter, err := measurer.MeasureShape(shape)
	if err != nil {
		fmt.Printf("测量 %s 失败: %v\n", name, err)
		return
	}
	fmt.Printf("%s - 面积: %.2f, 周长: %.2f\n", name, area, perimeter)
}
