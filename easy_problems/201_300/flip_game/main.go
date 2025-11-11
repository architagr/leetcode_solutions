package flipgame

func generatePossibleNextMoves(currentState string) []string {
	output := make([]string, 0)
	if len(currentState) <= 1 {
		return []string{}
	}

	for i := 0; i < len(currentState); i++ {
		if currentState[i] == '+' && i+1 <= len(currentState)-1 && currentState[i+1] == '+' {
			newState := currentState[:i] + "--" + currentState[i+2:]
			output = append(output, newState)
		}
	}
	return output
}
