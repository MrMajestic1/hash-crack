/*
Copyright © 2024 Quetzalcoatl

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
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help [COMMAND]",
	Short: "Displays information about program or a specific command.",
	Long: `Usage:
	hashcrack [COMMAND] [FLAGS] [HASH...]
Commands:
	crack      try to reverse find hash function inputs that produce the input hash values
	list       list supported hash functions
	type       try to determine the hash function(s) used to produce the input hash values
	buildtable build a rainbow table from provided file
	help       display this screen
Flags: 
	--functions="FUNCTION[ FUNCTION...]"  Specify a list of hash functions to try out
	--max-threads=NUMBER                  Maximum number of threads allowed. Defaults to number of CPU cores`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Long)
			return
		}

		if args[0] == "crack" {
			fmt.Println(crackCmd.Long)
		} else if args[0] == "buildtable" {
			fmt.Println(buildtableCmd.Long)
		}
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
