package helpers

import (
	"go-yao/pkg/global"
	"time"
)

func TimenowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(global.Conf.Application.Timezone)
	return time.Now().In(chinaTimezone)
}
