package day1

func FindFirstDigit(s []byte, ch chan int) {
	zs := []byte("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	zs_index := 0
	prev_level := 0
	go func() {
		for _, r := range s {
			zs[zs_index] = toLower(r)
			level := forwardNumberParser(zs)
			if level > 100 {

			}
			if isDigit(r) {
				ch <- int(r - '0')
				close(ch)
				return
			}
			if level > 100 {
				ch <- level - 100
				close(ch)
				return
			}
			if level == 0 || level == prev_level {
				resetToZ(zs)
				zs_index = 0
			} else {
				zs_index++
			}
			prev_level = level
		}
		close(ch)
	}()
}

func FindLastDigit(s []byte, ch chan int) {
	zs := []byte("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	zs_index := 0
	prev_level := 0

	go func() {
		for i := len(s) - 1; i >= 0; i-- {
			zs[zs_index] = toLower(s[i])
			level := reverseNumberParser(zs)

			if isDigit(s[i]) {
				ch <- int(s[i] - '0')
				close(ch)
				return
			}

			if level > 100 {
				ch <- level - 100
				close(ch)
				return
			}

			if level == 0 || level == prev_level {
				resetToZ(zs)
				zs_index = 0
			} else {
				zs_index++

			}

			prev_level = level
		}
		close(ch)
	}()
}
