package functions

import(

)

func Floor(str string)string{
	steps := 0

	for _, ch := range str{
		if ch == '('{
			steps++
		}else{
			steps--
		}
		
	}
return Itoa(steps)
}

func Itoa(n int)string{
sign := ""
result := ""
if n < 0{
	sign = "-"
	n = -n
}
if n == 0{
	return "0"
}
for n != 0{
	start := '0'
for i := 0; i < (n%10); i++{
	start++
}
result = string(start) + result
n /= 10
}
return sign + result
}