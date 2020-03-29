package world

func convertRange(v, oldMin, oldMax, newMin, newMax int) int {
	oldRange, newRange := oldMax-oldMin, newMax-newMin
	return ((v-oldMin)*newRange)/oldRange + newMin
}
