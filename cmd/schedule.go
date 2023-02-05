/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

func getJson3(url string, target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	if err != nil {
		log.Fatalln(err)
	}

	// fmt.Println(string(b))
	var data map[string]interface{}
	// Object.
	err2 := json.Unmarshal(b, &data)
	if err2 != nil {
		fmt.Println(err2)
	}
	// dates, _ := data.Get("dates")
	// vals, _ := data["dates"]
	dates := data["dates"].([]interface{})
	// fmt.Println(dates)
	for _, date := range dates {
		// fmt.Println(date)
		games := date.(map[string]interface{})["games"].([]interface{})
		for _, game := range games {
			fmt.Println(game.(map[string]interface{})["link"])
			fmt.Println(game.(map[string]interface{})["status"].(map[string]interface{})["detailedState"])
			// fmt.Println(game["link"])
			fmt.Println("--------------")
			// os.Exit(1)
		}
	}

	/*
		for _, val := range data["dates"].([]interface{}) {
			fmt.Println(val.(map[string]interface{})["date"])
			// fmt.Println(val.(map[string]interface{})["games"])
			for _, val2 := range val.(map[string]interface{})["games"].([]interface{}) {
				fmt.Println(val2.(map[string]interface{})["link"])
				// fmt.Println(val2)
			}
		}
	*/
	// dates := data["metaData"]
	// fmt.Println(dates)
	// fmt.Println(dates.(map[string]interface{})["timeStamp"])
	// fmt.Println(dates.(map[string]interface{}))
	return json.NewDecoder(resp.Body).Decode(target)
}

func getSCheduleUrl(start string, end string) string {
	scheme := viper.GetString("HttpScheme")
	host := viper.GetString("HttpHost")
	url := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   "api/v1/schedule",
	}

	q := url.Query()
	q.Set("startDate", start)
	q.Set("endDate", end)
	q.Set("gameType", "R")

	url.RawQuery = q.Encode()

	return url.String()
}

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Get a schedule of games",
	Long:  `Print out the schedule of a range of games. You will see teams, scores, and game status.`,
	Run: func(cmd *cobra.Command, args []string) {
		startDate, _ := cmd.Flags().GetString("start")
		endDate, _ := cmd.Flags().GetString("end")
		// fmt.Println(startDate, endDate)
		schedule := new(ScheduleList)
		scheduleUrl := getSCheduleUrl(startDate, endDate)
		getJson(scheduleUrl, schedule)
		// fmt.Printf("%+v\n", schedule)
		// fmt.Println(schedule.Dates[0].Games[0])
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Date", "ID", "Home Team", "Score", "Away Team", "Score", "Status"})
		// rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
		for _, Date := range schedule.Dates {
			// fmt.Println(Date.Date)

			t.AppendSeparator()
			for _, Game := range Date.Games {
				// fmt.Println(Game.Teams.Away.Team.Name, Game.Teams.Away.Score, " AT ", Game.Teams.Home.Team.Name, Game.Teams.Home.Score)
				t.AppendRow(table.Row{Date.Date, Game.ID, Game.Teams.Home.Team.Name, Game.Teams.Home.Score, Game.Teams.Away.Team.Name, Game.Teams.Away.Score, Game.Status.AbstractGameState})
				// tbl.AddRow(Date.Date, Game.Id, Game.Teams.Home.Team.Name, Game.Teams.Home.Score, Game.Teams.Away.Team.Name, Game.Teams.Away.Score)
			}

		}
		t.Render()
	},
}

func init() {
	getCmd.AddCommand(scheduleCmd)

	now := time.Now().Format("2006-01-02")
	yest := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	scheduleCmd.PersistentFlags().String("start", yest, "Schedule start date")
	scheduleCmd.PersistentFlags().String("end", now, "Schedule end date")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scheduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
