package main

import "math/rand"

type log struct {
	timestamp int
	length    int
}

func main() {
	pc1 := pcCreateData()
	pc2 := pcCreateData()
	total := logMerge(pc1, pc2)

	pcPrint(pc1)
	println()
	pcPrint(pc2)
	println()
	pcPrint(total)

}

func logMerge(pc1 []log, pc2 []log) []log {
	res := make([]log, 1)
	var c1, c2 int

	for c1 < len(pc1) && c2 < len(pc2) {
		if pc1[c1].timestamp < pc2[c2].timestamp {
			res = append(res, log{timestamp: pc1[c1].timestamp, length: res[len(res)-1].length + pc1[c1].length})
			c1++
		} else if pc1[c1].timestamp > pc2[c2].timestamp {
			res = append(res, log{timestamp: pc2[c2].timestamp, length: res[len(res)-1].length + pc2[c2].length})
			c2++
		} else if pc1[c1].timestamp == pc2[c2].timestamp {
			res = append(res, log{timestamp: pc2[c2].timestamp, length: res[len(res)-1].length + pc2[c2].length + pc1[c1].length})
			c1++
			c2++
		}
	}

	if c2 < len(pc2) && c1 == len(pc1) {
		for c2 < len(pc2) {
			res = append(res, log{timestamp: pc2[c2].timestamp, length: res[len(res)-1].length + pc2[c2].length})
			c2++
		}
	} else if c1 < len(pc1) && c2 == len(pc2) {
		for c1 < len(pc1) {
			res = append(res, log{timestamp: pc1[c1].timestamp, length: res[len(res)-1].length + pc1[c1].length})
			c1++
		}
	}
	return res
}

func pcCreateData() []log {
	pc := make([]log, 0)
	pc = append(pc, log{timestamp: 0, length: 0})
	for i := 0; i < 100; i++ {
		if rand.Intn(5) == 1 {
			tmp := rand.Intn(4)
			if tmp != 0 {
				pc = append(pc, log{timestamp: i, length: pc[len(pc)-1].length + tmp})
			}
		}
	}
	return pc
}

func pcPrint(pc []log) {
	for _, l := range pc {
		println(l.timestamp, " | ", l.length)
	}
}
