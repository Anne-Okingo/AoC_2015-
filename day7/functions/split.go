package functions

import (

)

func Split(s string)[]string{
	sliced := []string{}
	result := ""
for _, ch := range s{
	if ch != ' '{
		result += string(ch)
	}else if result != ""{
		sliced = append(sliced,result)
		result = ""
	}
}
if result != ""{
	sliced = append(sliced,result)
}
	return sliced
}