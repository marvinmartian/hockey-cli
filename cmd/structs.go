/*
Copyright © 2023 Melvin Wiens
*/
package cmd

import "time"

type ScheduleList struct {
	Copyright    string `json:"copyright,omitempty"`
	TotalItems   int    `json:"totalItems,omitempty"`
	TotalEvents  int    `json:"totalEvents,omitempty"`
	TotalGames   int    `json:"totalGames,omitempty"`
	TotalMatches int    `json:"totalMatches,omitempty"`
	MetaData     struct {
		TimeStamp string `json:"timeStamp,omitempty"`
	} `json:"metaData,omitempty"`
	Wait  int         `json:"wait,omitempty"`
	Dates []DatesList `json:"dates,omitempty"`
}

type DatesList struct {
	Date         string        `json:"date,omitempty"`
	TotalItems   int           `json:"totalItems,omitempty"`
	TotalEvents  int           `json:"totalEvents,omitempty"`
	TotalGames   int           `json:"totalGames,omitempty"`
	TotalMatches int           `json:"totalMatches,omitempty"`
	Games        []Game        `json:"games,omitempty"`
	Events       []interface{} `json:"events,omitempty"`
	Matches      []interface{} `json:"matches,omitempty"`
}

type Game struct {
	ID       int       `json:"gamePk,omitempty"`
	Link     string    `json:"link,omitempty"`
	GameType string    `json:"gameType,omitempty"`
	Season   string    `json:"season,omitempty"`
	GameDate time.Time `json:"gameDate,omitempty"`
	Status   Status    `json:"status,omitempty"`
	Teams    struct {
		Away TeamDetails `json:"away,omitempty"`
		Home TeamDetails `json:"home,omitempty"`
	} `json:"teams,omitempty"`
	Venue   Venue `json:"venue,omitempty,omitempty"`
	Content struct {
		Link string `json:"link,omitempty"`
	} `json:"content,omitempty"`
}

type Status struct {
	AbstractGameState string `json:"abstractGameState,omitempty"`
	CodedGameState    string `json:"codedGameState,omitempty"`
	DetailedState     string `json:"detailedState,omitempty"`
	StatusCode        string `json:"statusCode,omitempty"`
	StartTimeTBD      bool   `json:"startTimeTBD,omitempty"`
}

type TeamDetails struct {
	LeagueRecord struct {
		Wins   int    `json:"wins,omitempty"`
		Losses int    `json:"losses,omitempty"`
		Ot     int    `json:"ot,omitempty"`
		Type   string `json:"type,omitempty"`
	} `json:"leagueRecord,omitempty"`
	Score int  `json:"score,omitempty"`
	Team  Team `json:"team,omitempty"`
}

type Team struct {
	ID              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Link            string `json:"link,omitempty"`
	Venue           Venue  `json:"venue,omitempty"`
	Abbreviation    string `json:"abbreviation,omitempty"`
	TriCode         string `json:"triCode,omitempty"`
	TeamName        string `json:"teamName,omitempty"`
	LocationName    string `json:"locationName,omitempty"`
	FirstYearOfPlay string `json:"firstYearOfPlay,omitempty"`
	Division        struct {
		ID           int    `json:"id,omitempty"`
		Name         string `json:"name,omitempty"`
		NameShort    string `json:"nameShort,omitempty"`
		Link         string `json:"link,omitempty"`
		Abbreviation string `json:"abbreviation,omitempty"`
	} `json:"division,omitempty"`
	Conference struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Link string `json:"link,omitempty"`
	} `json:"conference,omitempty"`
	Franchise struct {
		FranchiseID int    `json:"franchiseId,omitempty"`
		TeamName    string `json:"teamName,omitempty"`
		Link        string `json:"link,omitempty"`
	} `json:"franchise,omitempty"`
	ShortName       string `json:"shortName,omitempty"`
	OfficialSiteURL string `json:"officialSiteUrl,omitempty"`
	FranchiseID     int    `json:"franchiseId,omitempty"`
	Active          bool   `json:"active,omitempty"`
	Roster          struct {
		Roster []struct {
			Person       Player `json:"person,omitempty"`
			JerseyNumber string `json:"jerseyNumber,omitempty"`
			Position     struct {
				Code         string `json:"code,omitempty"`
				Name         string `json:"name,omitempty"`
				Type         string `json:"type,omitempty"`
				Abbreviation string `json:"abbreviation,omitempty"`
			} `json:"position,omitempty"`
		} `json:"roster,omitempty"`
		Link string `json:"link,omitempty"`
	} `json:"roster,omitempty"`
}

// Game Structs

type SingleGame struct {
	Link     string   `json:"link"`
	GameData GameData `json:"gameData"`
	LiveData LiveData `json:"liveData"`
}

type GameData struct {
	Game struct {
		Pk     int    `json:"pk,omitempty"`
		Season string `json:"season,omitempty"`
		Type   string `json:"type,omitempty"`
	} `json:"game,omitempty"`
	Datetime struct {
		DateTime time.Time `json:"dateTime,omitempty"`
	} `json:"datetime,omitempty"`
	Status Status `json:"status"`
	Teams  struct {
		Away Team `json:"away"`
		Home Team `json:"home"`
	} `json:"teams"`
	Venue Venue `json:"venue,omitempty"`
}

type Venue struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Link     string `json:"link,omitempty"`
	City     string `json:"city,omitempty"`
	TimeZone struct {
		ID     string `json:"id,omitempty"`
		Offset int    `json:"offset,omitempty"`
		Tz     string `json:"tz,omitempty"`
	} `json:"timeZone,omitempty"`
}

type LiveData struct {
	Plays struct {
		AllPlays []Plays `json:"allPlays"`
	} `json:"plays"`
	LineScore Linescore `json:"linescore"`
	Boxscore  Boxscore  `json:"boxscore"`
}

type EventResult struct {
	Event       string `json:"event,omitempty"`
	EventCode   string `json:"eventCode,omitempty"`
	EventTypeID string `json:"eventTypeId,omitempty"`
	Description string `json:"description,omitempty"`
}

type EventAbout struct {
	EventIdx            int       `json:"eventIdx,omitempty"`
	EventID             int       `json:"eventId,omitempty"`
	Period              int       `json:"period,omitempty"`
	PeriodType          string    `json:"periodType,omitempty"`
	OrdinalNum          string    `json:"ordinalNum,omitempty"`
	PeriodTime          string    `json:"periodTime,omitempty"`
	PeriodTimeRemaining string    `json:"periodTimeRemaining,omitempty"`
	DateTime            time.Time `json:"dateTime,omitempty"`
	Goals               struct {
		Away int `json:"away,omitempty"`
		Home int `json:"home,omitempty"`
	} `json:"goals,omitempty"`
}

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type EventPlayers struct {
	Players    []Player `json:"player,omitempty"`
	PlayerType string   `json:"playerType,omitempty"`
}

type Player struct {
	ID                 int    `json:"id,omitempty"`
	FullName           string `json:"fullName,omitempty"`
	Link               string `json:"link,omitempty"`
	FirstName          string `json:"firstName,omitempty"`
	LastName           string `json:"lastName,omitempty"`
	PrimaryNumber      string `json:"primaryNumber,omitempty"`
	BirthDate          string `json:"birthDate,omitempty"`
	CurrentAge         int    `json:"currentAge,omitempty"`
	BirthCity          string `json:"birthCity,omitempty"`
	BirthStateProvince string `json:"birthStateProvince,omitempty"`
	BirthCountry       string `json:"birthCountry,omitempty"`
	Nationality        string `json:"nationality,omitempty"`
	Height             string `json:"height,omitempty"`
	Weight             int    `json:"weight,omitempty"`
	Active             bool   `json:"active,omitempty"`
	AlternateCaptain   bool   `json:"alternateCaptain,omitempty"`
	Captain            bool   `json:"captain,omitempty"`
	Rookie             bool   `json:"rookie,omitempty"`
	ShootsCatches      string `json:"shootsCatches,omitempty"`
	RosterStatus       string `json:"rosterStatus,omitempty"`
	CurrentTeam        struct {
		ID      int    `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
		Link    string `json:"link,omitempty"`
		TriCode string `json:"triCode,omitempty"`
	} `json:"currentTeam,omitempty"`
	PrimaryPosition struct {
		Code         string `json:"code,omitempty"`
		Name         string `json:"name,omitempty"`
		Type         string `json:"type,omitempty"`
		Abbreviation string `json:"abbreviation,omitempty"`
	} `json:"primaryPosition,omitempty"`
}

type Plays struct {
	Result      EventResult    `json:"result"`
	About       EventAbout     `json:"about"`
	Coordinates Coordinates    `json:"coordinates"`
	Players     []EventPlayers `json:"players,omitempty"`
	Team        Team           `json:"team,omitempty"`
}

type Linescore struct {
	CurrentPeriod              int    `json:"currentPeriod,omitempty"`
	CurrentPeriodOrdinal       string `json:"currentPeriodOrdinal,omitempty"`
	CurrentPeriodTimeRemaining string `json:"currentPeriodTimeRemaining,omitempty"`
	ShootoutInfo               struct {
		Away struct {
			Scores   int `json:"scores,omitempty"`
			Attempts int `json:"attempts,omitempty"`
		} `json:"away,omitempty"`
		Home struct {
			Scores   int `json:"scores,omitempty"`
			Attempts int `json:"attempts,omitempty"`
		} `json:"home,omitempty"`
	} `json:"shootoutInfo,omitempty"`
	Teams struct {
		Home struct {
			Team         Team `json:"team,omitempty"`
			Goals        int  `json:"goals,omitempty"`
			ShotsOnGoal  int  `json:"shotsOnGoal,omitempty"`
			GoaliePulled bool `json:"goaliePulled,omitempty"`
			NumSkaters   int  `json:"numSkaters,omitempty"`
			PowerPlay    bool `json:"powerPlay,omitempty"`
		} `json:"home,omitempty"`
		Away struct {
			Team         Team `json:"team,omitempty"`
			Goals        int  `json:"goals,omitempty"`
			ShotsOnGoal  int  `json:"shotsOnGoal,omitempty"`
			GoaliePulled bool `json:"goaliePulled,omitempty"`
			NumSkaters   int  `json:"numSkaters,omitempty"`
			PowerPlay    bool `json:"powerPlay,omitempty"`
		} `json:"away,omitempty"`
	} `json:"teams,omitempty"`
	PowerPlayStrength string `json:"powerPlayStrength,omitempty"`
	HasShootout       bool   `json:"hasShootout,omitempty"`
	IntermissionInfo  struct {
		IntermissionTimeRemaining int  `json:"intermissionTimeRemaining,omitempty"`
		IntermissionTimeElapsed   int  `json:"intermissionTimeElapsed,omitempty"`
		InIntermission            bool `json:"inIntermission,omitempty"`
	} `json:"intermissionInfo,omitempty"`
	PowerPlayInfo struct {
		SituationTimeRemaining int  `json:"situationTimeRemaining,omitempty"`
		SituationTimeElapsed   int  `json:"situationTimeElapsed,omitempty"`
		InSituation            bool `json:"inSituation,omitempty"`
	} `json:"powerPlayInfo,omitempty"`
}

type Boxscore struct {
	Teams struct {
		Away struct {
			Team      Team      `json:"team,omitempty"`
			TeamStats TeamStats `json:"teamStats,omitempty"`
			Goalies   []int     `json:"goalies,omitempty"`
			Skaters   []int     `json:"skaters,omitempty"`
			OnIce     []int     `json:"onIce,omitempty"`
			Scratches []int     `json:"scratches,omitempty"`
			Coaches   []struct {
				Person   Person `json:"person,omitempty"`
				Position struct {
					Code         string `json:"code,omitempty"`
					Name         string `json:"name,omitempty"`
					Type         string `json:"type,omitempty"`
					Abbreviation string `json:"abbreviation,omitempty"`
				} `json:"position,omitempty"`
			} `json:"coaches,omitempty"`
		} `json:"away,omitempty"`
		Home struct {
			Team      Team      `json:"team,omitempty"`
			TeamStats TeamStats `json:"teamStats,omitempty"`
			Goalies   []int     `json:"goalies,omitempty"`
			Skaters   []int     `json:"skaters,omitempty"`
			OnIce     []int     `json:"onIce,omitempty"`
			Scratches []int     `json:"scratches,omitempty"`
			Coaches   []struct {
				Person   Person `json:"person,omitempty"`
				Position struct {
					Code         string `json:"code,omitempty"`
					Name         string `json:"name,omitempty"`
					Type         string `json:"type,omitempty"`
					Abbreviation string `json:"abbreviation,omitempty"`
				} `json:"position,omitempty"`
			} `json:"coaches,omitempty"`
		} `json:"home,omitempty"`
	} `json:"teams,omitempty"`
	Officials []struct {
		Official     Person `json:"official,omitempty"`
		OfficialType string `json:"officialType,omitempty"`
	} `json:"officials,omitempty"`
}

type Person struct {
	ID       int    `json:"id,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Link     string `json:"link,omitempty"`
}

type TeamStats struct {
	TeamSkaterStats struct {
		Goals                  int     `json:"goals,omitempty"`
		Pim                    int     `json:"pim,omitempty"`
		Shots                  int     `json:"shots,omitempty"`
		PowerPlayPercentage    string  `json:"powerPlayPercentage,omitempty"`
		PowerPlayGoals         float64 `json:"powerPlayGoals,omitempty"`
		PowerPlayOpportunities float64 `json:"powerPlayOpportunities,omitempty"`
		FaceOffWinPercentage   string  `json:"faceOffWinPercentage,omitempty"`
		Blocked                int     `json:"blocked,omitempty"`
		Takeaways              int     `json:"takeaways,omitempty"`
		Giveaways              int     `json:"giveaways,omitempty"`
		Hits                   int     `json:"hits,omitempty"`
	} `json:"teamSkaterStats,omitempty"`
}

type SinglePlayer struct {
	Copyright string   `json:"copyright,omitempty"`
	Player    []Player `json:"people,omitempty"`
}

type TeamList struct {
	Copyright string `json:"copyright,omitempty"`
	Teams     []Team `json:"teams,omitempty"`
}
