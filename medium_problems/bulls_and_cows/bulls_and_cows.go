package bulls_and_cows

import "fmt"

func GetHint(secret, guess string) string {
	arr := make([]int, 10)
	cows, bulls := 0, 0
	for i := range secret {
		a, b := secret[i]-'0', guess[i]-'0'
		if a == b {
			bulls++
		} else {
			if arr[a] < 0 {
				cows++
			}
			if arr[b] > 0 {
				cows++
			}
		}
		arr[a]++
		arr[b]--
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func GetHintApproch1(secret, guess string) string {
	bull, cow := 0, 0
	mapSecret, mapGuess := convertString(secret), convertString(guess)
	for key, sIndexes := range mapSecret {
		if gIndexes, gok := mapGuess[key]; gok {
			if len(sIndexes) > len(gIndexes) {
				count(gIndexes, sIndexes, &bull, &cow)
			} else {
				count(sIndexes, gIndexes, &bull, &cow)
			}
		}
	}

	return fmt.Sprintf("%dA%dB", bull, cow)
}

func count(a, b map[int]bool, bull, cow *int) {
	for k := range a {
		if _, ok := b[k]; ok {
			(*bull)++
		} else {
			(*cow)++
		}
	}
}

func convertString(secret string) map[byte]map[int]bool {
	mapData := make(map[byte]map[int]bool)

	for i := 0; i < len(secret); i++ {
		if val, ok := mapData[secret[i]]; ok {
			val[i] = true
			mapData[secret[i]] = val
		} else {
			val = make(map[int]bool)
			val[i] = true
			mapData[secret[i]] = val
		}
	}

	return mapData
}
