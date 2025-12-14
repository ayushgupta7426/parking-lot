package parkingfees

type FeeStrategy interface {
	Calculate(hours int) int
}
