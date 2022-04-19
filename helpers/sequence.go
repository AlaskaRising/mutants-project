package helpers

func diagonalValuestoSlice(dna []string, lenDna, maxLen int) (diagonalDna []string) {
	var auxLen = lenDna - 1 //Max position in a array for len N
	var diagonalValue1, diagonalValue2 string = "", ""
	var auxMaxDiagonal int = 0

	//Diagonal Values Bottom to Top -> Top Values
	for i := auxLen; i >= 3; i-- {
		for j := 0; j < lenDna-auxMaxDiagonal; j++ {
			diagonalValue1 += string(dna[i-j][j])
		}
		diagonalDna = append(diagonalDna, diagonalValue1)
		diagonalValue1 = ""
		auxMaxDiagonal++
	}

	//Diagonal Values Bottom to Top -> Bottom Values
	auxMaxDiagonal = 1
	for i := auxLen; i >= 4; i-- {
		for j := 0; j < lenDna-auxMaxDiagonal; j++ {
			diagonalValue1 += string(dna[auxLen-j][j+auxMaxDiagonal])
		}
		diagonalDna = append(diagonalDna, diagonalValue1)
		diagonalValue1 = ""
		auxMaxDiagonal++
	}

	//Diagonal Values Top to Bottom -> Central Value
	for i := 0; i < lenDna; i++ {
		diagonalValue1 += string(dna[i][i])
	}
	diagonalDna = append(diagonalDna, diagonalValue1)

	//Diagonal Values Top to Bottom -> Top and Bottom Values
	diagonalValue1 = ""
	for i := 0; i <= maxLen-1; i++ {
		for j := 1; j <= lenDna-1-i; j++ {
			diagonalValue1 += string(dna[j-1][j+i])
			diagonalValue2 += string(dna[j+i][j-1])
		}
		diagonalDna = append(diagonalDna, diagonalValue1, diagonalValue2)
		diagonalValue1, diagonalValue2 = "", ""
	}
	return diagonalDna
}

func validateHorizontalDna(dna []string, lenDna, maxLen int) int {
	var auxSequence = 0
	for i := 0; i < lenDna; i++ {
		for j := 0; j <= maxLen; j++ {
			if (dna[i][j] == dna[i][j+1]) && (dna[i][j+1] == dna[i][j+2]) && (dna[i][j+2] == dna[i][j+3]) {
				auxSequence++
			}
		}
	}
	return auxSequence
}

func validateVerticalDna(dna []string, lenDna, maxLen int) int {
	var auxSequence = 0
	for i := 0; i < lenDna; i++ {
		for j := 0; j <= maxLen; j++ {
			if (dna[j][i] == dna[j+1][i]) && (dna[j+1][i] == dna[j+2][i]) && (dna[j+2][i] == dna[j+3][i]) {
				auxSequence++
			}
		}
	}
	return auxSequence
}

func validateDiagonalDna(diagDna []string) int {
	var auxSequence = 0
	var lenDiagonalDna = len(diagDna)
	for i := 0; i < lenDiagonalDna; i++ {
		var maxLen = len(diagDna[i]) - 4
		for j := 0; j <= maxLen; j++ {
			if (diagDna[i][j] == diagDna[i][j+1]) && (diagDna[i][j+1] == diagDna[i][j+2]) && (diagDna[i][j+2] == diagDna[i][j+3]) {
				auxSequence++
			}
		}
	}
	return auxSequence
}

func IsMutant(dna []string) bool {
	var totalSequence int = 0
	var lenDna = len(dna)   //Len N, Table NxN
	var maxLen = lenDna - 4 //Max position to allow a mutan DNA (Ej:GGGG)
	if lenDna < 4 {
		return false
	}
	totalSequence += validateHorizontalDna(dna, lenDna, maxLen)
	totalSequence += validateVerticalDna(dna, lenDna, maxLen)
	totalSequence += validateDiagonalDna(diagonalValuestoSlice(dna, lenDna, maxLen))
	if totalSequence > 1 {
		return true
	}
	return false
}
