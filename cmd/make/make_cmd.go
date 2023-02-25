package make

import (
	"fmt"
	"go-yao/pkg/console"

	"github.com/spf13/cobra"
)

var CmdMakeCMD = &cobra.Command{
	Use:   "cmd",
	Short: "Create a command, should be snake_case, example: make cmd buckup_database",
	Run:   runMakeCMD,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeCMD(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])                                         // 格式化模型名称，返回一个 Model 对象
	createFileFromStub(fmt.Sprintf("cmd/%s.go", model.PackageName), "cmd", model) // 从模板中创建文件（做好变量替换）

	console.Success("make cmd success ...")
}
