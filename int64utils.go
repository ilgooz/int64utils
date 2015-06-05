package int64utils

func Difference(a, b []int64) []int64 {
	out := make([]int64, 0)

	for _, v := range a {
		in := false

		for _, v2 := range b {
			if v2 == v {
				in = true
				break
			}
		}

		if !in {
			out = append(out, v)
		}
	}

	return out
}
