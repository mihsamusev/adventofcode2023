package main


type Range struct {
    start int
    length int
}

func (r *Range) Contains(other Range) bool {
    return (r.start <= other.start) && (r.start + r.length >= other.start + other.length)
}

func (r *Range) Union(other Range) Range {
    start := r.start
    end := r.start + r.length
    if start > other.start {
        start = other.start
    }
    if end < other.start + other.length {
        end = other.start + other.length
    }

    newLength := end - start
    if r.length + other.length > newLength {
        return Range{0, 0}
    }

    return Range{start, newLength}
}

func (r* Range) Intersection(other Range) Range {
    return Range{0, 0}
}

func (r* Range) Sub(other Range) Range {
    return Range{0, 0}
}

