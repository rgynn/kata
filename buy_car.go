package kata

import (
	"math"
)

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	oldf, newf := float64(startPriceOld), float64(startPriceNew)
	month, savings, available := 0, 0.0, float64(startPriceOld-startPriceNew)
	for available < 0 {
		if month%2 == 1 {
			percentLossByMonth += 0.5
		}
		oldf = oldf * (1.0 - (percentLossByMonth / 100))
		newf = newf * (1.0 - (percentLossByMonth / 100))
		savings += float64(savingperMonth)
		available = oldf + savings - newf
		month++
	}
	return [2]int{month, int(math.Round(available))}
}
