package main

import "fmt"

type Lookup struct {
	dstStart int
	srcRange Range
}

func (l *Lookup) Contains(src int) bool {
	diff := src - l.srcRange.start
	return (diff >= 0) && (diff <= l.srcRange.length-1)
}

func (l *Lookup) Convert(src int) int {
	if !l.Contains(src) {
		return src
	}
	diff := src - l.srcRange.start
	return l.dstStart + diff
}

func (l *Lookup) ConvertRange(r Range) ([]Range) {
    // find itersection -> it gets re-mapped

    // find other parts -> they continue unchanged
    return []Range{{0,0}, {0,0}, {0,0}}
}

type FarmingMap struct {
	lookups []Lookup
}

func (f *FarmingMap) Convert(src int) int {
	for _, lookup := range f.lookups {
		if lookup.Contains(src) {
			return lookup.Convert(src)
		}
	}
	return src
}

func Trace(src int, maps []FarmingMap) int {
    dst := src
    fmt.Printf("%d", src)
    for _, m := range maps {
        dst = m.Convert(dst)
        fmt.Printf(" -> %d", dst)
    }
    fmt.Printf("\n")
    return dst
}

func TraceRange(r Range, maps[]FarmingMap) []Range {
    stack := make([]Range, 0, 1)
    stack[0] = r

    for _, m := range maps {
        newRanges := m.Convert()

        
    }
 
}

func ExpandSeeds(seedRanges []int) []int {
    seeds := make([]int, 0)
    for i := 0; i < len(seedRanges) / 2; i++ {
        seedStart := seedRanges[2 * i]
        length := seedRanges[2 * i + 1]
        for s := seedStart; s < seedStart + length; s++ {
            seeds = append(seeds, s)
        }
    }
    return seeds
}