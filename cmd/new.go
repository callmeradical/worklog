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

	"github.com/callmeradical/worklog/lib/data"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new weekly worklog",
	Run: func(cmd *cobra.Command, args []string) {
		path := viper.GetString("LogDir")
		if path == "" {
			fmt.Println("LogDir is not set")
		}

		filename := viper.GetString("CurrentLog")
		if filename == "" {
			fmt.Println("No file found, generating new file...")
			err := UpdateConfig("CurrentLog", "current_log.json")
			if err != nil {
				panic(err)
			}
		}

		if data.LogExists(data.WorkLogPath(path, filename)) {
			archiveLogPath := viper.GetString("ArchiveLogDir")
			if archiveLogPath == "" {
				archiveLogPath = path
				err := UpdateConfig("ArchiveLogDir", path)
				if err != nil {
					panic(err)
				}
			}

			err := data.ArchiveLog(path, filename)
			if err != nil {
				panic(err)
			}
		}

		err := data.CreateLogFile(data.WorkLogPath(path, filename))
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
