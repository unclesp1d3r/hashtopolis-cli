/*
Copyright © 2022 UncleSp1d3r <unclespider@protonmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/unclesp1d3r/hashtopolis-cli/api"
	"github.com/unclesp1d3r/hashtopolis-cli/utils"
	"strconv"
)

// listTasksCmd represents the listTasks command
var listTasksCmd = &cobra.Command{
	Use:   "listTasks",
	Short: "List all tasks on the server",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := api.ListTasks(viper.GetString("apiUrl"), viper.GetString("apiKey"))
		if err != nil {
			cobra.CheckErr(err)
		}

		var data = [][]string{}
		for _, element := range result.Tasks {
			data = append(data, []string{
				strconv.Itoa(element.TaskId),
				element.Name,
				strconv.Itoa(element.HashlistId),
				strconv.Itoa(element.Priority),
				typeIdToText(element.Type),
			})
		}
		utils.PrintTable([]string{"ID", "Name", "Hashlist ID", "Priority", "Type"}, data)
	},
}

func init() {
	rootCmd.AddCommand(listTasksCmd)
}

func typeIdToText(id int) string {
	if id == 0 {
		return "Normal Task"
	} else {
		return "Supertask"
	}
}
