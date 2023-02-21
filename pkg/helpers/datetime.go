package helpers

import (
	"fmt"
	"time"
)

// MicrosecondsStr 将 time.Duration 类型输出为小数点后 3 位的 ms
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}
