package common

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type CliArgs struct {
	FileName string
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