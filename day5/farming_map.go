package main

import "fmt"

type Lookup struct {
	dstStart int
	srcRange Range
}

func (l *Lookup) Contains(src int) bool {
	diff := src - l.srcRange.start
    length := l.srcRange.end - l.srcRange.start
	return (diff >= 0) && (diff <= length)
}

func (l *Lookup) Convert(src int) int {
	if !l.Contains(src) {
		return src
	}
	diff := src - l.srcRange.start
	return l.dstStart + diff
}

func (l *Lookup) ConvertRange(r Range) []Range {
    ranges := make([]Range, 0, 3)

    // find itersection -> it gets re-mapped
    toRemap := r.Intersection(l.srcRange)

    if toRemap.Equal(Range{}) {
        ranges = append(ranges, r)
        return ranges
    }

    diffStart := r.start - l.srcRange.start
    diffEnd := r.end - l.srcRange.start
    remapped := Range{l.dstStart + diffStart, l.dstStart + diffEnd}
    ranges = append(ranges, remapped)
    // input range fully contained in the lookup range
    if toRemap.Equal(r) {
        return ranges
    }

    // input range is "hanging" from left and/or right of the lookup range
    diffRanges := toRemap.Difference(r)
    ranges = append(ranges, diffRanges...)
    return ranges
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
    prevBatch := make([]Range, 0, 1)
    prevBatch[0] = r

    for _, m := range maps {
        nextBatch := make([]Range, 0, 1)
        for _, p := range prevBatch {
            if p.Equal(Range{}) {
                continue
            }

            for _, lookup := range m.lookups {
                newRanges := lookup.ConvertRange(p)
                nextBatch = append(nextBatch, newRanges...)
            }
        }
        prevBatch = nextBatch
    }
    return prevBatch
}

func InterpretRanges(seedRanges []int) []Range {
    ranges := make([]Range, 0)
    for i := 0; i < len(seedRanges) / 2; i++ {
        seedStart := seedRanges[2 * i]
        length := seedRanges[2 * i + 1]
        ranges = append(ranges, Range{seedStart, length}) 
    }
    return ranges
}
