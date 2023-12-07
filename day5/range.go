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
    if sumLength < newLength {
        return Range{0, 0}
    }

    return Range{start, end}
}

func (r* Range) Intersection(other Range) Range {
    return Range{0, 0}
}

func (r* Range) Sub(other Range) Range {
    return Range{0, 0}
}

