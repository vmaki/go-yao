package make

import (
	"fmt"
	"go-yao/pkg/console"
	"os"

	"github.com/spf13/cobra"
)

var CmdMakeModel = &cobra.Command{
	Use:   "model",
	Short: "Crate model file, example: make model user",
	Run:   runMakeModel,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeModel(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])                   // 格式化模型名称，返回一个 Model 对象
	dir := fmt.Sprintf("app/models/%s/", model.PackageName) // 确保模型的目录存在，例如 `app/models/user`
	os.MkdirAll(dir, os.ModePerm)                           // os.MkdirAll 会确保父目录和子目录都会创建，第二个参数是目录权限，使用 0777

	createFileFromStub(dir+model.PackageName+".go", "model", model)
	createFileFromStub(dir+"util.go", "model_util", model)
	createFileFromStub(dir+"hooks.go", "model_hooks", model)

	console.Success("make model success ...")
}
