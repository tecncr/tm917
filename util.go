package tm917

// Returns the first index of seq in s, or -1 if not found.
func findSequence(s, seq string) int {
	for i := 0; i+len(seq) <= len(s); i++ {
		if s[i:i+len(seq)] == seq {
			return i
		}
	}
	return -1
}
