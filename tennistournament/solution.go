package tennistournament

func Solution(P, C int) int {
	halfPlayers := P/2
	if halfPlayers > C {
		return C
	}
	return halfPlayers
}
