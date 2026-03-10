package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	s := []float64{} //empty slice
	for i:=0; i<len(costs); i++{
		if costs[i].day == day {
			s = append(s, costs[i].value)
		}
	}
	return s
}
