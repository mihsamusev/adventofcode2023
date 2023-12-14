package main

import "strings"

func ParseMap(str string) Map {
	rows := strings.Split(str, "\n")

	m := make(Map, 0)
	for i := range rows {
		m = append(m, []rune(rows[i]))
	}

	return m
}