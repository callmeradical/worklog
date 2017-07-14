// Copyright © 2017 Lars Cromley lars@callmeradical.com
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

	"github.com/callmeradical/worklog/lib/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var activity string
var activityType string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new entry to the weekly worklog",
	Run: func(cmd *cobra.Command, args []string) {
		if activity == "" {
			fmt.Println("Must provide an activity to log")
			os.Exit(1)
		}

		filename := viper.GetString("CurrentLog")
		dir := viper.GetString("LogDir")
		if filename == "" || dir == "" {
			fmt.Println("Must create a new worklog first!")
			os.Exit(1)
		}

		path := data.WorkLogPath(dir, filename)

		//Write message to log
		err := data.WriteToWorkLog(path, activityType, activity)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

	},
}

func init() {
	RootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringVarP(&activity, "activity", "a", "", "the activity to add to the work log")
	addCmd.Flags().StringVarP(&activityType, "type", "t", "DONE", "type field for the activity")

}
