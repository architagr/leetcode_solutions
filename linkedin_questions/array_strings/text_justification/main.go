package textjustification

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	resultJustifiedParagraph := make([]string, 0, len(words))
	startIndex := 0
	currentSentenceLength := 0
	addPending := func() {
		str := buildJustifySentence(words[startIndex:], maxWidth, true)
		resultJustifiedParagraph = append(resultJustifiedParagraph, str)
	}
	for lastIndex := 0; lastIndex < len(words); lastIndex++ {
		word := words[lastIndex]
		// for each sentence
		if currentSentenceLength == 0 {
			// first word in the sentence
			currentSentenceLength = len(word)
			startIndex = lastIndex
			if lastIndex == len(words)-1 {
				addPending()
			}
			continue
		}
		pridictedLength := lastIndex - startIndex + currentSentenceLength + len(word)

		if pridictedLength >= maxWidth {
			subSet := words[startIndex:lastIndex]
			lastIndex--
			if pridictedLength == maxWidth {
				lastIndex++
				subSet = words[startIndex : lastIndex+1]
			}
			str := buildJustifySentence(subSet, maxWidth, false)
			resultJustifiedParagraph = append(resultJustifiedParagraph, str)
			currentSentenceLength = 0
			continue
		} else if lastIndex == len(words)-1 {
			addPending()
		}
		currentSentenceLength += len(word)
	}
	return resultJustifiedParagraph
}

func buildJustifySentence(words []string, maxWidth int, isLeftJustified bool) string {
	sb := strings.Builder{}
	sb.Grow(maxWidth)
	spaceBlocks := len(words) - 1
	unjustifiedWidth := spaceBlocks
	for _, word := range words {
		unjustifiedWidth += len(word)
	}
	additionalSpace := 0
	if spaceBlocks > 0 {
		additionalSpace = (maxWidth - unjustifiedWidth) / spaceBlocks
	}
	spaceBlockBuilder := strings.Builder{}
	spaceBlockBuilder.Grow(additionalSpace)
	spaceBlockBuilder.WriteRune(' ')

	if !isLeftJustified {
		for i := 0; i < additionalSpace; i++ {
			spaceBlockBuilder.WriteRune(' ')
		}
	}
	spaceBlock := spaceBlockBuilder.String()
	extraSpaces := (maxWidth - unjustifiedWidth)
	if spaceBlocks > 0 {
		extraSpaces = extraSpaces % spaceBlocks
	}
	for _, word := range words[:spaceBlocks] {
		sb.WriteString(word)
		sb.WriteString(spaceBlock)
		if extraSpaces > 0 && !isLeftJustified {
			sb.WriteRune(' ')
			extraSpaces--
		}
	}

	sb.WriteString(words[spaceBlocks])
	for i := sb.Len(); i < maxWidth; i++ {
		sb.WriteRune(' ')
	}
	return sb.String()
}
