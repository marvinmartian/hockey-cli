/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getTeamUrl(TeamId string) string {
	scheme := viper.GetString("HttpScheme")
	host := viper.GetString("HttpHost")
	// http://localhost/api/v1/game/2022020754/feed/live
	// fmt.Println(host)
	url := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   "api/v1/teams/" + TeamId,
	}

	q := url.Query()
	// q.Set("expand", "team.roster,team.stats")
	q.Set("expand", "team.roster")

	url.RawQuery = q.Encode()

	return url.String()
}

// playerCmd represents the player command
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Get details for a specific team",
	Long:  `Return details for a specific team. `,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		teamId := args[0]
		team := new(TeamList)
		teamUrl := getTeamUrl(teamId)
		getJson(teamUrl, team)
		teamDetails := team.Teams[0]

		// fmt.Printf("%+v\n", teamDetails.Roster.Roster)

		fmt.Println("------------")
		fmt.Printf("%s\n", teamDetails.Name)
		// fmt.Println("------------")
		fmt.Printf("Venue: %s\n", teamDetails.Venue.Name)
		// fmt.Println("------------")

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Name", "Jersey", "Position"})
		for _, player := range teamDetails.Roster.Roster {
			// fmt.Printf("%+v\n", player.Person.FullName)
			t.AppendRow(table.Row{player.Person.FullName, player.JerseyNumber, player.Position.Name})
		}
		t.Render()
	},
}

func init() {
	getCmd.AddCommand(teamCmd)

}
