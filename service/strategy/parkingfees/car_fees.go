package parkingfees

type CarFee struct {
}

func (b *CarFee) Calculate(hours int) int {
	return hours * 100
}
