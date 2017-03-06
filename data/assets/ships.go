package assets

func Ship1() ([][]rune) {

	matrix := [][]rune{}
    row1 := []rune{'▚', '█', '▞'}
    row2 := []rune{' ', '▓', ' '}
    row3 := []rune{'■', ' ', '■'}
    matrix = append(matrix, row1)
    matrix = append(matrix, row2)
    matrix = append(matrix, row3)

    return matrix
}

func Ship2() ([][]rune) {

	matrix := [][]rune{}
    row1 := []rune{'█', '█'}
    row2 := []rune{'▜', '▛'}
    matrix = append(matrix, row1)
    matrix = append(matrix, row2)

    return matrix
}

func Ship3() ([][]rune) {

    matrix := [][]rune{}
    row1 := []rune{'╱', '╲'}
    row2 := []rune{'▩', '▩'}
    matrix = append(matrix, row1)
    matrix = append(matrix, row2)

    return matrix
}