package helpers

import (
	"go-yao/common/global"
	"time"
)

func IsLocal() bool {
	return global.Conf.Application.Mode == "local"
}
func TimeNowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(global.Conf.Application.Timezone)

	return time.Now().In(chinaTimezone)
}
