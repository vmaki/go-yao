package make

import (
	"fmt"
	"go-yao/pkg/console"
	"strings"

	"github.com/spf13/cobra"
)

var CmdMakeController = &cobra.Command{
	Use:   "controller",
	Short: "Create api controller，example: make controller api/v1/user",
	Run:   runMakeController,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeController(cmd *cobra.Command, args []string) {
	// 处理参数，要求附带 API 版本（v1 或者 v2）
	array := strings.Split(args[0], "/")
	if len(array) != 3 {
		console.Exit("controller name format: api/v1/user")
	}

	controllerName, apiVersion, name := array[0], array[1], array[2]
	model := makeModelFromString(name)

	filePath := fmt.Sprintf("app/http/controllers/%s/%s/%s.go", controllerName, apiVersion, model.TableName)
	createFileFromStub(filePath, "controller", model)

	console.Success("make controller success ...")
}
