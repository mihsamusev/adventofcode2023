package main

import (
	"common"
	"strings"
)

func ParseHands(str string) ([]Hand, error){
	rows := strings.Split(str, "\n")

	hands := make([]Hand, 0)
	for _, row := range rows {
		pair, err := common.ParseNameValuePair(row)
		if err != nil {
			return hands, err
		}

		cards := []Card(pair.Name)
		hands = append(hands, Hand{cards, pair.Value})
	}
	return hands, nil
}