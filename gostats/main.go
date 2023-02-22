package main

import (
	//"example.com/gostats/calculations"
	"fmt"

	"example.com/gostats/calculations"
)

func main() {
	firstPlayer := calculations.Batter{PlayerID: "troutmi01", YearID: 2014, Stint: 1, TeamID: "LAA", LeagueID: "AL", Games: 157, AtBats: 602, Runs: 115, Hits: 173, Doubles: 39, Triples: 9, Homeruns: 36, RunsBattedIn: 111, StolenBases: 16, CaughtStealing: 2, Walks: 83, Strikeouts: 184, IntentionalWalks: 6, HitByPitch: 10, SacrificeBunt: 0, SacrificeFly: 10, GroundIntoDoublePlay: 6, BattingAverage: 0, OnBasePercentage: 0, SluggingPercentage: 0, OnBasePlusSlugging: 0}
	fmt.Println(firstPlayer)
	fmt.Println("Play Ball!")
}
