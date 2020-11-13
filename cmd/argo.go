/*
Copyright © 2020 Ambor <saltbo@foxmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/saltbo/zpan/config"
	"github.com/spf13/cobra"
	"os/exec"
)

// argoCmd represents the argo command
var argoCmd = &cobra.Command{
	Use:   "argo",
	Short: "Launch local aria2c daemon.",
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Parse()
		if err := LaunchAria2cDaemon(c.Aria2.RPCSecret); err != nil {
			fmt.Println("Aria2c daemon start failed.")
			fmt.Println(err)
		}
		fmt.Println("Aria2c daemon start success.")
	},
}

func init() {
	rootCmd.AddCommand(argoCmd)
}

// LaunchAria2cDaemon Launch aria2 daemon to listen for RPC calls, locally.
func LaunchAria2cDaemon(secret string) (err error) {
	args := []string{"--enable-rpc", "--rpc-listen-all"}
	if secret != "" {
		args = append(args, "--rpc-secret="+secret)
	}
	cmd := exec.Command("aria2c", args...)
	if err = cmd.Start(); err != nil {
		return
	}
	return cmd.Process.Release()
}
