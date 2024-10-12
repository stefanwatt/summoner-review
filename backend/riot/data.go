package riot

type Participant struct {
	SummonerName                string    `json:"summonerName"`
	SummonerIconUrl             string    `json:"summonerIconUrl"`
	SummonerSpells              [2]string `json:"summonerSpells"`
	ChampionName                string    `json:"championName"`
	ChampionLevel               int       `json:"championLevel"`
	ChampionIconUrl             string    `json:"championIconUrl"`
	Lane                        string    `json:"lane"`
	Kills                       int       `json:"kills"`
	Assists                     int       `json:"assists"`
	Deaths                      int       `json:"deaths"`
	TotalDamageDealtToChampions int       `json:"totalDamageDealtToChampions"`
	TotalDamageTaken            int       `json:"totalDamageTaken"`
	WardsPlaced                 int       `json:"wardsPlaced"`
	ControlWardsPlaced          int       `json:"controlWardsPlaced"`
	WardsDestroyed              int       `json:"wardsDestroyed"`
	CreepScore                  int       `json:"creepScore"`
	Items                       [6]string `json:"items"`
}

type Team struct {
	Victorious          bool           `json:"victorious"`
	TotalKills          int            `json:"totalKills"`
	TotalGold           int            `json:"totalGold"`
	BaronsKilled        int            `json:"baronsKilled"`
	DragonsKilled       int            `json:"dragonsKilled"`
	HeraldsKilled       int            `json:"heraldsKilled"`
	VoidGrubsKilled     int            `json:"voidGrubsKilled"`
	TowersDestroyed     int            `json:"towersDestroyed"`
	InhibitorsDestroyed int            `json:"inhibitorsDestroyed"`
	Participants        [5]Participant `json:"participants"`
}

type Match struct {
	RedTeam         Team `json:"redTeam"`
	BlueTeam        Team `json:"blueTeam"`
	DurationSeconds int  `json:"durationSeconds"`
}
