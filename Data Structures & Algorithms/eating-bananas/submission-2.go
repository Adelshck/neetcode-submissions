import "slices"
func calcTime(piles []int,k int) int{
	time := 0
	for _, v := range piles {
		time += (v + k - 1) / k
	}
	return time
}	
func minEatingSpeed(piles []int, h int) int {

	
	mx := slices.Max(piles)
	l, r:= 1, mx 
	k := 0
	for l <=r {
		mid := l + (r - l) / 2
		time := calcTime(piles,mid)
		if time <= h {
			k = mid
			r = mid - 1
		} else { 
			l = mid + 1
		}
	}

	return k
}