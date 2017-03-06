package assets

func Laser1() ([][]rune) {

	matrix := [][]rune{}
    row1 := []rune{'|'}
    matrix = append(matrix, row1)

    return matrix
}

func Laser2() ([][]rune) {

	matrix := [][]rune{}
    row1 := []rune{'║'}
    row2 := []rune{'║'}
    matrix = append(matrix, row1)
    matrix = append(matrix, row2)

    return matrix
}