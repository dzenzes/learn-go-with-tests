package iteration

func Repeat(character string, repeatCount int) string {
	if repeatCount == 0 {
		repeatCount = 5
	}
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
