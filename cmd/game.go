/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func getGameUrl(GameId string) string {
	scheme := viper.GetString("HttpScheme")
	host := viper.GetString("HttpHost")
	// http://localhost/api/v1/game/2022020754/feed/live
	// fmt.Println(host)
	url := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   "api/v1/game/" + GameId + "/feed/live",
	}

	return url.String()
}

// gameCmd represents the game command
var gameCmd = &cobra.Command{
	Use:   "game",
	Short: "Get a specific game details",
	Long:  `Print out all of the information from a single game. `,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(args[0])
		gameId := args[0]
		game := new(SingleGame)
		gameUrl := getGameUrl(gameId)
		getJson(gameUrl, game)
		// fmt.Println(game.LiveData.Boxscore.Officials)

		// fmt.Printf("%+v\n", game.LiveData.Boxscore.Officials)
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Period", "Time", "Remaining", "Event", "Description", "Coordinates"})
		for _, play := range game.LiveData.Plays.AllPlays {
			// fmt.Printf("%+v\n", play)
			// fmt.Printf("%+v\n", play)

			coords := ""
			if play.Coordinates.X != 0 && play.Coordinates.Y != 0 {
				// strconv.FormatFloat(10.900, 'f', -1, 64)
				coords = fmt.Sprintf("%s, %s", strconv.FormatFloat(play.Coordinates.X, 'f', -1, 64), strconv.FormatFloat(play.Coordinates.Y, 'f', -1, 64))
			}

			t.AppendRow(table.Row{play.About.Period, play.About.PeriodTime, play.About.PeriodTimeRemaining, play.Result.Event, play.Result.Description, coords})
		}
		t.Render()
	},
}

func init() {
	getCmd.AddCommand(gameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
