package matrix

type Matrix interface {
	Dim() (int, int)
	DimN() int
	DimM() int
}

type SquareMatrix interface {
	Matrix
	Det() float64
}

type DimError error
type NotSquareError DimError
type MultDimError DimError

type NoInverseError error

// func Mult(a, b Matrix) (Matrix, MultDimError) {
//
// }
//
// func Inverse(a Matrix) (Matrix, NoInverseError) {
//
// }
//
// func Add(a, b Matrix) (Matrix, error) {
//
// }

func Init() {

}
