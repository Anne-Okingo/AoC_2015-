package functions

import (

)

func Separator(s, sep string)[]string{
	lensep := len(sep)
	result := []string{}
	start := 0

	for i := 0; i < len(s)-lensep; i++{
		if s[i:i+lensep] == sep{
			result = append(result,s[start:i+lensep])
			start = i + lensep
		}
		result = append(result,s[start:])
	}
	return result
}