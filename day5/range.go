package main

type Range struct {
	start int
	end   int
}

func (r *Range) Equal(other Range) bool {
	return (r.start == other.start) && (r.end == other.end)
}

func (r *Range) IsEmpty() bool {
	return r.start == r.end
}

func (r *Range) HasValue(v int) bool {
	offset := v - r.start
	maxOffset := r.end - r.start
	return (offset >= 0) && (offset <= maxOffset)
}

func (r *Range) Offset(o int) Range {
	if r.Equal(Range{}) {
		return Range{}
	}
	return Range{r.start + o, r.end + o}
}

func (r *Range) RemapStart(s int) Range {
	return r.Offset(s - r.start)
}

func (r *Range) In(other Range) bool {
	return (r.start >= other.start) && (r.end <= other.end)
}

func (r *Range) Union(other Range) Range {
	start := r.start
	end := r.end
	if start > other.start {
		start = other.start
	}
	if end < other.end {
		end = other.end
	}

	newLength := end - start
	sumLength := r.end - r.start + other.end - other.start
	if newLength > sumLength+1 {
		return Range{}
	}

	return Range{start, end}
}

func (r *Range) Intersection(other Range) Range {
	newStart := Max(r.start, other.start)
	newEnd := Min(r.end, other.end)
	if newStart > newEnd {
		return Range{}
	}
	return Range{newStart, newEnd}
}

func (r *Range) Diff(other Range) (left, right Range) {
	if r.In(other) {
		return Range{}, Range{}

	}
	if (r.Intersection(other) == Range{}) {
		return Range{}, Range{}
	}

	if r.start < other.start {
		left = Range{r.start, other.start - 1}
	}

	if r.end > other.end {
		right = Range{other.end + 1, r.end}
	}

	return left, right
}

func Min(first, second int) int {
	if first < second {
		return first
	} else {
		return second
	}
}

func Max(first, second int) int {
	if first > second {
		return first
	} else {
		return second
	}
}
