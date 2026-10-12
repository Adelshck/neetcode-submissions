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

	stack := make([]float64, 0)
	for i := range cars {
		time := float64(target - cars[i].position) / float64(cars[i].speed)
		if len(stack) == 0 {
			cnt++
			stack = append(stack, time)
			continue
		}
		if time > stack[len(stack) - 1] {
			cnt++
			stack = append(stack, time)
		} 
		
	}
	return cnt
}
