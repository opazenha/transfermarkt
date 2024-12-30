package main

// Response represents the top-level response structure.
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

// Data represents the data structure containing player information and performance.
type Data struct {
	Player      Player        `json:"player"`
	Performance []Performance `json:"performance"`
}

// Player represents the player's details.
type Player struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	DateOfBirth   string      `json:"dateOfBirth"`
	Age           int         `json:"age"`
	Gender        string      `json:"gender"`
	Nationalities Nationality `json:"nationalities"`
	Attributes    Attributes  `json:"attributes"`
	Assignment    Assignment  `json:"assignment"`
	PortraitURL   string      `json:"portraitUrl"`
}

// Assignment represents the player's assignment details.
type Assignment struct {
	Title      string `json:"title"`
	PreviewURL string `json:"previewUrl"`
}

// Nationality represents the player's nationality information.
type Nationality struct {
	NationalityID       int     `json:"nationalityId"`
	SecondNationalityID int     `json:"secondNationalityId"`
	Nationality         Country `json:"nationality"`
}

// Country represents the country details.
type Country struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	ConfederationID int    `json:"confederationId"`
	FlagURL         string `json:"flagUrl"`
}

// Attributes represents the player's attributes.
type Attributes struct {
	PositionID int      `json:"positionId"`
	Position   Position `json:"position"`
}

// Position represents the position details of the player.
type Position struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

// Performance represents the player's performance statistics.
type Performance struct {
	PlayerInformation     PlayerInformation     `json:"playerInformation"`
	GoalStatistics        GoalStatistics        `json:"goalStatistics"`
	CardStatistics        CardStatistics        `json:"cardStatistics"`
	PlayingTimeStatistics PlayingTimeStatistics `json:"playingTimeStatistics"`
	Competition           Competition           `json:"competition"`
}

// PlayerInformation contains information about the player in a specific season.
type PlayerInformation struct {
	PlayerID      string `json:"playerId"`
	Season        int    `json:"season"`
	CompetitionID string `json:"competitionId"`
	ClubID        string `json:"clubId"`
}

// GoalStatistics holds statistics related to goals.
type GoalStatistics struct {
	GoalsCount     int `json:"goalsCount"`
	AssistsCount   int `json:"assistsCount"`
	ScorersCount   int `json:"scorersCount"`
	PenaltiesCount int `json:"penaltiesCount"`
	OwnGoalsCount  int `json:"ownGoalsCount"`
}

// CardStatistics holds statistics related to cards.
type CardStatistics struct {
	YellowCardsCount    int `json:"yellowCardsCount"`
	YellowRedCardsCount int `json:"yellowRedCardsCount"`
	RedCardsCount       int `json:"redCardsCount"`
}

// PlayingTimeStatistics holds statistics related to playing time.
type PlayingTimeStatistics struct {
	StartingElevenCount            int `json:"startingElevenCount"`
	NotPlayedMinutesSentOff        int `json:"notPlayedMinutesSentOff"`
	SubstitutedInCount             int `json:"substitutedInCount"`
	NotPlayedMinutesSubstitutedIn  int `json:"notPlayedMinutesSubstitutedIn"`
	SubstitutedOutCount            int `json:"substitutedOutCount"`
	NotPlayedMinutesSubstitutedOut int `json:"notPlayedMinutesSubstitutedOut"`
	OvertimesCount                 int `json:"overtimesCount"`
	AppearancesCount               int `json:"appearancesCount"`
	TotalMinutesPlayed             int `json:"totalMinutesPlayed"`
}

// Competition holds details about the competition.
type Competition struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	LogoURL          string           `json:"logoUrl"`
	OriginDetails    OriginDetails    `json:"originDetails"`
	BaseDetails      BaseDetails      `json:"baseDetails"`
	TotalMarketValue TotalMarketValue `json:"totalMarketValue"`
	Type             CompetitionType  `json:"type"`
	Label            string           `json:"label"`
	Link             string           `json:"link"`
}

// OriginDetails holds details about the competition's origin.
type OriginDetails struct {
	CountryID   int `json:"countryId"`
	ContinentID int `json:"continentId"`
}

// BaseDetails holds basic details about the competition.
type BaseDetails struct {
	StandardGameDurationMinutes int  `json:"standardGameDurationMinutes"`
	GameDayCount                int  `json:"gameDayCount"`
	IsOneOff                    bool `json:"isOneOff"`
	IsOngoing                   bool `json:"isOngoing"`
	IsTournament                bool `json:"isTournament"`
}

// TotalMarketValue holds the market value details.
type TotalMarketValue struct {
	Value          int    `json:"value"`
	Currency       string `json:"currency"`
	Display        string `json:"display"`
	CompactDisplay string `json:"compactDisplay"`
}

// CompetitionType holds the competition type details.
type CompetitionType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Link    string `json:"link"`
	Image1x string `json:"image1x"`
	Image2x string `json:"image2x"`
	IsNT    bool   `json:"isNT"`
}

type Match struct {
	Competition  Competition  `json:"competition"`
	ID           int          `json:"id"`
	Match        MatchDetails `json:"match"`
	Integrations []string     `json:"integrations"`
}

type MatchDetails struct {
	Home    int    `json:"home"`
	Away    int    `json:"away"`
	Day     string `json:"day"`
	DayLink string `json:"dayLink"`
	Injury  struct {
	} `json:"injury"`
	Group struct {
	} `json:"group"`
	Link            string `json:"link"`
	Result          string `json:"result"`
	ResultExtension string `json:"resultExtension"`
	State           string `json:"state"`
	Suspension      struct {
	} `json:"suspension"`
	Time int64 `json:"time"`
}

type TeamMap map[int]Team

type Club struct {
	Name         string `json:"clubName"`
	Emblem1x     string `json:"clubEmblem-1x"`
	Emblem2x     string `json:"clubEmblem-2x"`
	EmblemMobile string `json:"clubEmblemMobile"`
	CountryFlag  string `json:"countryFlag"`
}

type Transfer struct {
	ID                 int    `json:"id"`
	URL                string `json:"url"`
	From               Club   `json:"from"`
	To                 Club   `json:"to"`
	FutureTransfer     int    `json:"futureTransfer"`
	Date               string `json:"date"`
	DateUnformatted    string `json:"dateUnformatted"`
	Upcoming           bool   `json:"upcoming"`
	Season             string `json:"season"`
	MarketValue        string `json:"marketValue"`
	Fee                string `json:"fee"`
	ShowUpcomingHeader bool   `json:"showUpcomingHeader"`
	ShowResetHeader    bool   `json:"showResetHeader"`
}

type Transfers struct {
	Transfers       []Transfer `json:"transfers"`
	FeeSum          int        `json:"feeSum"`
	URL             string     `json:"url"`
	FormattedFeeSum string     `json:"formattedFeeSum"`
	Translations    struct {
		Headline            string `json:"headline"`
		TransferHistory     string `json:"transfer history"`
		UpcomingTransfers   string `json:"upcoming transfers"`
		Feesum              string `json:"feesum"`
		Season              string `json:"season"`
		Date                string `json:"date"`
		ClubFrom            string `json:"club from"`
		ClubTo              string `json:"club to"`
		MarketValue         string `json:"market value"`
		TransferFee         string `json:"transfer fee"`
		TransferMapLinkText string `json:"transferMapLinkText"`
	} `json:"translations"`
}

type MarketValue struct {
	List []struct {
		X       int64  `json:"x"`
		Y       int    `json:"y"`
		MW      string `json:"mw"`
		DatumMW string `json:"datum_mw"`
		Verein  string `json:"verein"`
		Age     string `json:"age"`
		Wappen  string `json:"wappen"`
	} `json:"list"`
	Current     string `json:"current"`
	Highest     string `json:"highest"`
	HighestDate string `json:"highest_date"`
	LastChange  string `json:"last_change"`
	DetailsURL  string `json:"details_url"`
	Thread      struct {
		URL          string `json:"url"`
		ThreadTitle  string `json:"thread_title"`
		CountReplies int    `json:"count_replies"`
	} `json:"thread"`
	Translations struct {
		MarketValue string `json:"market value"`
		Team        string `json:"team"`
		Age         string `json:"age"`
		ResetZoom   string `json:"resetZoom"`
		Current     string `json:"current"`
		Highest     string `json:"highest"`
		Thread      string `json:"thread"`
		Forum       string `json:"forum"`
		Details     string `json:"details"`
		Headline    string `json:"headline"`
		LastChange  string `json:"lastChange"`
	} `json:"translations"`
}
