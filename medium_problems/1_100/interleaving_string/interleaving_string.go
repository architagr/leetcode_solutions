package interleaving_string

func isInterleave(s1, s2, s3 string) bool {
	// length of all 3 strings
	ls1, ls2, ls3 := len(s1), len(s2), len(s3)

	if ls1 == 0 && ls2 == 0 && ls3 == 0 {
		return true
	}
	if (ls1 + ls2) != ls3 {
		return false
	}

	// ittr index for s1, s2 and s3 respectively
	i, j, k := ls1-1, ls2-1, ls3-1

	if i >= 0 && s1[i] == s3[k] && j >= 0 && s2[j] == s3[k] {
		s3 = s3[:k]
		return isInterleave(s1[:i], s2, s3) || isInterleave(s1, s2[:j], s3)
	}
	if i >= 0 && k >= 0 && s1[i] == s3[k] {
		return isInterleave(s1[:i], s2, s3[:k])
	}
	if j >= 0 && k >= 0 && s2[j] == s3[k] {
		return isInterleave(s1, s2[:j], s3[:k])
	}

	return false
}
