package make

import (
	"fmt"
	"go-yao/pkg/console"

	"github.com/spf13/cobra"
)

var CmdMakeDto = &cobra.Command{
	Use:   "dto",
	Short: "Create dto file, example make dto user",
	Run:   runMakeDto,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeDto(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])
	createFileFromStub(fmt.Sprintf("app/http/dto/%s.go", model.PackageName), "dto", model)

	console.Success("make dto success ...")
}
