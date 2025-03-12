package uniqueemailaddresses

import "strings"

func numUniqueEmails(emails []string) int {
	uniqueEmail := make(map[string]bool, len(emails))
	for _, email := range emails {
		uniqueEmail[parseEmail(email)] = true
	}
	return len(uniqueEmail)
}

func parseEmail(email string) string {
	stringBuilder := strings.Builder{}
	stringBuilder.Grow(len(email))
	i := 0
	foundPlus := false
	for ; i < len(email); i++ {
		if email[i] == '@' {
			break
		}
		if email[i] == '.' {
			continue
		}
		if email[i] == '+' || foundPlus {
			foundPlus = true
			continue
		}
		stringBuilder.WriteByte(email[i])
	}
	stringBuilder.WriteString(string(email[i:]))
	return stringBuilder.String()
}
