package main

func btoi(value bool) int {
	if value {
		return 1
	}

	return 0
}

func setNthBit(num *int, ix int, value bool) {
	bitmask := 1 << ix
	*num = (*num &^ bitmask) | (-btoi(value) & bitmask)
}
