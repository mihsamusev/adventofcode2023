package common

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type CliArgs struct {
	FileName  string
	LineCount int
}

func ReadCliArgs() CliArgs {
	args := CliArgs{"test.txt", -1}
	argParts := os.Args
	if len(argParts) > 1 {
		args.FileName = argParts[1]
	}

	if len(argParts) > 2 {
		lineCount, err := strconv.Atoi(argParts[2])
		if err == nil {
			args.LineCount = lineCount
		}
	}
	return args
}

func ParseId(str, prefix string) (int, error) {
	result, found := strings.CutPrefix(str, prefix)
	if !found {
		return -1, nil
	}
	result = strings.TrimSpace(result)
	return strconv.Atoi(result)
}

func ParseNamedSlice(str, prefix string) ([]int, error) {
	result, found := strings.CutPrefix(str, prefix)
	if !found {
		return make([]int, 0), errors.New("no prefix found")
	}
	return ParseSlice(result)
}

func ParseNameValuePair(str string) (
	struct {
		Name  string
		Value int
	}, error) {
	result := struct {
		Name  string
		Value int
	}{"", 0}
	parts := strings.Fields(str)
	if len(parts) != 2 {
		return result, errors.New("expected exactly 2 parts")
	}

	result.Name = parts[0]

	value, err := strconv.Atoi(parts[1])
	if err != nil {
		return result, err
	}

	result.Value = value
	return result, nil
}

func ParseSlice(str string) ([]int, error) {
	trimmed := strings.Fields(str)
	slice := make([]int, 0)
	for _, t := range trimmed {
		n, err := strconv.Atoi(t)
		if err != nil {
			return slice, err
		}
		slice = append(slice, n)
	}
	return slice, nil
}

func ParseDelimitedSlice(str string, delimiter string) ([]int, error) {
	split := strings.Split(str, delimiter)
	slice := make([]int, 0)
	for _, s := range split {
		t := strings.TrimSpace(s)
		n, err := strconv.Atoi(t)
		if err != nil {
			return slice, err
		}
		slice = append(slice, n)
	}
	return slice, nil
}

func ParseSlices(str, sep string) ([][]int, error) {
	rows := strings.Split(str, sep)
	slices := make([][]int, 0, len(rows))

	for _, row := range rows {
		slice, err := ParseSlice(row)
		if err != nil {
			return slices, err
		}
		slices = append(slices, slice)
	}
	return slices, nil
}