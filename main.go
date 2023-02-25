package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-yao/boot"
	"go-yao/cmd"
	"go-yao/cmd/make"
	"go-yao/pkg/console"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "GoYao",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			boot.SetupConfig(cmd.Env)
			boot.SetupLogger()
			boot.SetupDB()
			boot.SetupRedis()

			// 初始化缓存
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		make.CmdMake,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
