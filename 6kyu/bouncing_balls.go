package codewars


func BouncingBall(h, bounce, window float64) int {
	if h <= float64(0) || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}
	timesSeen := 1
	bounceHeight := h
	for bounceHeight > window {
		bounceHeight *= bounce
		if bounceHeight > window {
			timesSeen += 2
		}
	}
	return timesSeen
}