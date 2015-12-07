package util

type Direction int

const (
	None Direction = iota
	Up
	Down
	Left
	Right
)

const Timeconst = 100000
const ReferenceSpeed = 100

func GenerateKey(primeKey, attribute string) string {
	return primeKey + "_" + attribute
}
