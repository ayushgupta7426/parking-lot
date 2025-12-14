package parkingfees

type BikeFee struct {
}

func (b *BikeFee) Calculate(hours int) int {
	return hours * 50
}
