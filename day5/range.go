package main


type Range struct {
    start int
    end int
}

func (r *Range) Equal(other Range) bool {
   return (r.start == other.start) && (r.end == other.end) 
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
    if newLength > sumLength + 1{
        return Range{}
    }

    return Range{start, end}
}

func (r* Range) Intersection(other Range) Range {
    newStart := Max(r.start, other.start)
    newEnd := Min(r.end, other.end)
    if newStart > newEnd {
        return Range{}
    }
    return Range{newStart, newEnd}
}

func (r* Range) Difference(other Range) (left, right Range) {
    return Range{}, Range{0, 0}
}

func (r* Range) Sub(other Range) Range {
    return Range{}
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
