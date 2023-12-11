package main

import "fmt"

type Lookup struct {
	dstStart int
	srcRange Range
}

func (l *Lookup) Contains(src int) bool {
	return l.srcRange.HasValue(src)
}

func (l *Lookup) Offset() int {
	return l.dstStart - l.srcRange.start
}

func (l *Lookup) Convert(src int) int {
	if !l.Contains(src) {
		return src
	}
	offset := src - l.srcRange.start
	return l.dstStart + offset
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

func (f *FarmingMap) ConvertRange(r Range) []Range {
	ranges := make([]Range, 0)

	stack := make([]Range, 0)
	stack = append(stack, r)

	found := false
	for len(stack) > 0 {
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		found = false

		for _, l := range f.lookups {
			inter := popped.Intersection(l.srcRange)

			fmt.Printf("stack: %v popped: %v lookup %v\n", stack, popped, l.srcRange)
			if inter.IsEmpty() {
				continue
			}

			found = true
			mapped := inter.Offset(l.Offset())
			ranges = append(ranges, mapped)

			// check if there are parts left
			if popped.In(l.srcRange) {
				fmt.Printf("popped %v completely in lookup %v\n", inter, l.srcRange)
				break
			}

			left, right := popped.Diff(inter)
			fmt.Printf("left: %v, inter: %v, right %v\n", left, inter, right)
			if !left.IsEmpty() {
				stack = append(stack, left)
			}
			if !right.IsEmpty() {
				stack = append(stack, right)
			}
			break
		}

		if !found {
			ranges = append(ranges, popped)
		}
	}
	return ranges
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

func TraceRanges(ranges []Range, maps []FarmingMap) []Range {
	mappingStages := make([][]Range, 0)
	maxStages := len(maps)
	mappingStages = append(mappingStages, ranges)

	for i, m := range maps {
		nextMappingStage := make([]Range, 0)
		for _, r := range mappingStages[i] {
			nextMappingStage = append(nextMappingStage, m.ConvertRange(r)...)
		}
		mappingStages = append(mappingStages, nextMappingStage)
	}

	return mappingStages[maxStages]
}

func SliceAsRanges(seedRanges []int) []Range {
	ranges := make([]Range, 0)
	for i := 0; i < len(seedRanges)/2; i++ {
		start := seedRanges[2*i]
		length := seedRanges[2*i+1]
		end := start + length - 1
		ranges = append(ranges, Range{start, end})
	}
	return ranges
}
