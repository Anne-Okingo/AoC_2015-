package functions

import(

)

func IsDigit(s string)bool{
	for _, ch := range s{
		if ch >= '0' && ch <= '9'{
			return true
		}
	}
	return false
}