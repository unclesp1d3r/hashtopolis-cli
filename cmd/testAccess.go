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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/unclesp1d3r/hashtopolis-cli/api"
)

// testAccessCmd represents the testAccess command
var testAccessCmd = &cobra.Command{
	Use:   "testAccess",
	Short: "Tests connectivity with the server and verifies API key",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := api.TestConnection(viper.GetString("apiUrl"), viper.GetString("apiKey"))
		cobra.CheckErr(err)
		fmt.Println("Conection Successful")
	},
}

func init() {
	rootCmd.AddCommand(testAccessCmd)
}
