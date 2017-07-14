// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	"github.com/callmeradical/worklog/lib/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var newType string
var index int

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the status of a previous entry",
	Run: func(cmd *cobra.Command, args []string) {
		if index < 1 {
			fmt.Println("Must provide an index to update")
			os.Exit(1)
		}
		filename := viper.GetString("CurrentLog")
		dir := viper.GetString("LogDir")
		if filename == "" || dir == "" {
			fmt.Println("Must create a new worklog first!")
			os.Exit(1)
		}

		err := data.UpdateActivity(dir, filename, newType, index)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	updateCmd.Flags().StringVarP(&newType, "type", "t", "DONE", "new type field for the activity")
	updateCmd.Flags().IntVarP(&index, "item", "i", 0, "the activity to update")
}
