package array


//1094  .  拼车
func carPooling(trips [][]int, capacity int) bool {
	diffs := NewDiffArr(1001)

	for _,trip := range trips {
		diffs.Incr(trip[1] ,trip[2] -1 ,trip[0])
	}

	res := diffs.Result()
	for _,v := range res {
		if v>capacity {
			return false
		}
	}
	return true
}



