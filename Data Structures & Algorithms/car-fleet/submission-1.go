import "slices"

func carFleet(target int, position []int, speed []int) int {
	type Car struct {
    position int
    speed    int
	}
	cnt := 0
	cars := make([]Car, 0, len(position))
	for i := range position {
		cars = append(cars, Car{
			position: position[i],
			speed:    speed[i],
		})
	}
	slices.SortFunc(cars, func(a, b Car) int {
    	return b.position - a.position
	})
	s := 0.0
	for i := range cars {
		time := float64(target - cars[i].position) / float64(cars[i].speed)
		if i == 0 {
			cnt++
			s = time
			continue
		}
		if time > s{
			cnt++
			s = time
		} 
		
	}
	return cnt
}
