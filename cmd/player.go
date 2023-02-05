/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getPlayerUrl(PlayerId string) string {
	scheme := viper.GetString("HttpScheme")
	host := viper.GetString("HttpHost")
	// http://localhost/api/v1/game/2022020754/feed/live
	// fmt.Println(host)
	url := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   "api/v1/people/" + PlayerId,
	}
	return url.String()
}

func playerPrint(title string, content string, prefix string) {
	fmt.Printf("%s:\n", title)
	fmt.Println(strings.Repeat("-", len(title)+1))
	for _, line := range strings.Split(content, "\n") {
		fmt.Printf("%s%s\n", prefix, line)
	}
	fmt.Println()
}

// playerCmd represents the player command
var playerCmd = &cobra.Command{
	Use:   "player",
	Short: "Get details for a specific player",
	Long:  `Return details for a specific player. `,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(args[0])
		playerId := args[0]
		player := new(SinglePlayer)
		playerUrl := getPlayerUrl(playerId)
		// fmt.Println(playerUrl)
		getJson(playerUrl, player)
		playerDetails := player.Player[0]
		// fmt.Println(game.LiveData.Boxscore.Officials)

		// fmt.Printf("%+v\n", player.Player[0])

		fmt.Println("------------")
		fmt.Printf("%s:\n", playerDetails.FullName)
		fmt.Println("------------")
		fmt.Printf("%s: %s, %s\n", "Birth City", playerDetails.BirthCity, playerDetails.BirthStateProvince)
		fmt.Printf("%s: %s\n", "Birth Country", playerDetails.BirthCountry)
		fmt.Printf("%s: %s\n", "Birthday", playerDetails.BirthDate)
		fmt.Printf("%s: %v\n", "Age", playerDetails.CurrentAge)
		fmt.Println("------------")
		fmt.Printf("%s: %s\n", "Height", playerDetails.Height)
		fmt.Printf("%s: %v\n", "Weight", playerDetails.Weight)
		fmt.Println("------------")
		fmt.Printf("%s: %s\n", "Shoots/Catches", playerDetails.ShootsCatches)
		fmt.Println("------------")
		fmt.Printf("%s: %s\n", "Current Team", playerDetails.CurrentTeam.Name)
		fmt.Println("------------")
		fmt.Printf("%s: %s\n", "Position", playerDetails.PrimaryPosition.Name)

		// t := table.NewWriter()
		// t.SetOutputMirror(os.Stdout)
		// t.AppendHeader(table.Row{"Period", "Time", "Remaining", "Event", "Description", "Coordinates"})
		// for _, play := range game.LiveData.Plays.AllPlays {
		// 	// fmt.Printf("%+v\n", play)
		// 	// fmt.Printf("%+v\n", play)

		// 	coords := ""
		// 	if play.Coordinates.X != 0 && play.Coordinates.Y != 0 {
		// 		// strconv.FormatFloat(10.900, 'f', -1, 64)
		// 		coords = fmt.Sprintf("%s, %s", strconv.FormatFloat(play.Coordinates.X, 'f', -1, 64), strconv.FormatFloat(play.Coordinates.Y, 'f', -1, 64))
		// 	}

		// 	t.AppendRow(table.Row{play.About.Period, play.About.PeriodTime, play.About.PeriodTimeRemaining, play.Result.Event, play.Result.Description, coords})
		// }
		// t.Render()
	},
}

func init() {
	getCmd.AddCommand(playerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// playerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// playerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
