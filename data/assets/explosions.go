package assets

func Explosion1F1() ([][]rune) {

    matrix := [][]rune{}
    row1 := []rune{'█', '█', '█'}
    row2 := []rune{'█', '█', '█'}
    row3 := []rune{'█', '█', '█'}
    matrix = append(matrix, row1)
    matrix = append(matrix, row2)
    matrix = append(matrix, row3)

    return matrix
}

func Explosion1F2() ([][]rune) {

    matrix := [][]rune{}
    row1 := []rune{'▓', '▓', '▓'}
    row2 := []rune{'▓', '█', '▓'}
    row3 := []rune{'▓', '▓', '▓'}
    matrix = append(matrix, row1)
    matrix = append(matrix, row2)
    matrix = append(matrix, row3)

    return matrix
}

func Explosion1F3() ([][]rune) {

    matrix := [][]rune{}
    row1 := []rune{'▒', '▒', '▒', '▒'}
    row2 := []rune{'▒', '▓', '▓', '▒'}
    row3 := []rune{'▒', '▓', '▓', '▒'}
    row4 := []rune{'▒', '▒', '▒', '▒'}
    matrix = append(matrix, row1)
    matrix = append(matrix, row2)
    matrix = append(matrix, row3)
    matrix = append(matrix, row4)

    return matrix
}

func Explosion1F4() ([][]rune) {

    matrix := [][]rune{}
    row1 := []rune{'░', '░', '░', '░'}
    row2 := []rune{'░', '▓', '▓', '░'}
    row3 := []rune{'░', '▓', '▓', '░'}
    row4 := []rune{'░', '░', '░', '░'}
    matrix = append(matrix, row1)
    matrix = append(matrix, row2)
    matrix = append(matrix, row3)
    matrix = append(matrix, row4)

    return matrix
}

func Explosion1F5() ([][]rune) {

    matrix := [][]rune{}
    row1 := []rune{' ', ' ', ' ', '⢕'}
    row2 := []rune{' ', '⢵', '░', '░'}
    row3 := []rune{' ', '░', '▓', '⡧'}
    row4 := []rune{'░', '░', '░', ' '}
    matrix = append(matrix, row1)
    matrix = append(matrix, row2)
    matrix = append(matrix, row3)
    matrix = append(matrix, row4)

    return matrix
}