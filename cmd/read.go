// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read the current worklog",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		filename := viper.GetString("CurrentLog")
		dir := viper.GetString("LogDir")
		if filename == "" || dir == "" {
			fmt.Println("Must create a new worklog first!")
			os.Exit(1)
		}

		out, err := ReadLog(dir, filename)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(out)

	},
}

func init() {
	RootCmd.AddCommand(readCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
