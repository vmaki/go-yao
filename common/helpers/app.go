package helpers

import "go-yao/pkg/global"

func IsLocal() bool {
	return global.Conf.Application.Mode == "local"
}
