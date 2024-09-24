package main

func Ft_min_window(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	window := make(map[byte]int)
	left, right := 0, 0
	valid := 0
	start := 0
	minLength := len(s) + 1

	for right < len(s) {
		c := s[right]
		right++

		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < minLength {
				start = left
				minLength = right - left
			}

			d := s[left]
			left++

			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if minLength == len(s)+1 {
		return ""
	}

	return s[start : start+minLength]
}

func main() {
	Ft_min_window("ADOBECODEBANC", "ABC") // resultat : "BANC"
	Ft_min_window("a", "aa")              // resultat : ""
}
