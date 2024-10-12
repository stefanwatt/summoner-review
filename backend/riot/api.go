package riot

import (
	"context"
	"fmt"

	"github.com/Kyagara/equinox/v2"
	"github.com/Kyagara/equinox/v2/api"
	"github.com/Kyagara/equinox/v2/clients/lol"
)

var client, err = equinox.NewClient(RIOT_API_KEY)

func GetMatches(puuid string) ([]string, error) {
	ctx := context.Background()
	matchIds, err := client.LOL.MatchV5.ListByPUUID(
		ctx,
		api.EUROPE,
		puuid,
		-1,
		-1,
		-1,
		"",
		-1,
		20,
	)
	if err != nil {
		fmt.Println("error retrieving champion rotation: ", err)
		return nil, err
	}
	return matchIds, nil
}

func GetMatch(id string) (Match, error) {
	ctx := context.Background()
	match, err := client.LOL.MatchV5.ByID(ctx, api.EUROPE, id)
	if err != nil {
		fmt.Println("error retrieving match: ", err)
		return Match{}, err
	}
	var (
		redTeam  Team
		blueTeam Team
	)

	if match.Info.Teams[0].TeamID == lol.RED {
		redTeam = mapTeam(match.Info.Teams[0])
		blueTeam = mapTeam(match.Info.Teams[1])
	} else {
		redTeam = mapTeam(match.Info.Teams[1])
		blueTeam = mapTeam(match.Info.Teams[0])
	}
	redTeam.Participants = mapPlayersOfTeam(match.Info.Participants, lol.RED)
	blueTeam.Participants = mapPlayersOfTeam(match.Info.Participants, lol.BLUE)

	return Match{
		DurationSeconds: int(match.Info.GameDuration),
		RedTeam:         redTeam,
		BlueTeam:        blueTeam,
	}, nil
}

func mapSummonerIconUrl(id int32) string {
	return fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/14.20.1/img/profileicon/%d.png", id)
}

func mapSummonerSpellIconUrl(id int32) string {
	return ""
}

func mapChampionIconUrl(championName string) string {
	return fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/14.20.1/img/champion/%s.png", championName)
}

func mapItems(input lol.MatchParticipantV5DTO) [6]string {
	itemBaseUrl := "https://ddragon.leagueoflegends.com/cdn/14.20.1/img/item/"
	return [6]string{
		fmt.Sprintf("%s%d.png", itemBaseUrl, input.Item1),
		fmt.Sprintf("%s%d.png", itemBaseUrl, input.Item2),
		fmt.Sprintf("%s%d.png", itemBaseUrl, input.Item3),
		fmt.Sprintf("%s%d.png", itemBaseUrl, input.Item4),
		fmt.Sprintf("%s%d.png", itemBaseUrl, input.Item5),
		fmt.Sprintf("%s%d.png", itemBaseUrl, input.Item6),
	}
}

func mapTeam(input lol.MatchTeamV5DTO) Team {
	return Team{
		Victorious:          input.Win,
		TotalKills:          1,
		TotalGold:           1,
		BaronsKilled:        int(input.Objectives.Baron.Kills),
		DragonsKilled:       int(input.Objectives.Dragon.Kills),
		HeraldsKilled:       int(input.Objectives.RiftHerald.Kills),
		VoidGrubsKilled:     int(input.Objectives.Horde.Kills),
		TowersDestroyed:     int(input.Objectives.Tower.Kills),
		InhibitorsDestroyed: int(input.Objectives.Inhibitor.Kills),
	}
}

func mapPlayersOfTeam(input []lol.MatchParticipantV5DTO, teamId lol.Team) [5]Participant {
	var participants [5]Participant
	if len(input) != 10 {
		panic("a game must have exactly 10 participants")
	}
	count := 0
	for _, participant := range input {
		if count > 4 {
			break
		}
		if participant.TeamID == teamId {
			participants[count] = mapParticipant(participant)
			count = count + 1
		}
	}
	return participants
}

func mapParticipant(input lol.MatchParticipantV5DTO) Participant {
	creepScore := int(input.TotalMinionsKilled) + int(input.TotalAllyJungleMinionsKilled) + int(input.TotalEnemyJungleMinionsKilled)
	return Participant{
		SummonerName:    input.SummonerName,
		SummonerIconUrl: mapSummonerIconUrl(input.ProfileIcon),
		SummonerSpells: [2]string{
			mapSummonerSpellIconUrl(input.Summoner1ID),
			mapSummonerSpellIconUrl(input.Summoner2ID),
		},
		ChampionName:                input.ChampionName,
		ChampionLevel:               int(input.ChampLevel),
		ChampionIconUrl:             mapChampionIconUrl(input.ChampionName),
		Lane:                        input.Lane,
		Kills:                       int(input.Kills),
		Assists:                     int(input.Assists),
		Deaths:                      int(input.Deaths),
		TotalDamageDealtToChampions: int(input.TotalDamageDealtToChampions),
		TotalDamageTaken:            int(input.TotalDamageTaken),
		WardsPlaced:                 int(input.WardsPlaced),
		WardsDestroyed:              int(input.WardsKilled),
		CreepScore:                  creepScore,
		Items:                       mapItems(input),
	}
}
