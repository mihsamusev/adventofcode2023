package main

type Lookup struct {
    dstStart int
    srcStart int
    srcWidth int
}

func (this *Lookup) Contains(src int) bool {
    diff := src - this.srcStart
    return diff >= 0 && diff <= this.srcWidth
}


func (this *Lookup) Convert(src int) int {
     if !this.Contains(src) {
        return src
     }
    diff := src - this.srcStart
     return this.dstStart + diff
}

type FarmingMap struct {
    fields []Lookup
}
