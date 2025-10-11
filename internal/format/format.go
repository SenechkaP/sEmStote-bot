package format

import "strconv"

func FormatNumberWithDots(n int) string {
	str := strconv.Itoa(n)
	var result []byte

	count := 0
	for i := len(str) - 1; i >= 0; i-- {
		if count > 0 && count%3 == 0 {
			result = append([]byte{'.'}, result...)
		}
		result = append([]byte{str[i]}, result...)
		count++
	}

	return string(result)
}
