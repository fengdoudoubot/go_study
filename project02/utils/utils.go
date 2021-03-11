package utils
//package abc//也可以

func Cal(n1 float64,n2 float64,operator byte) float64{
	var res float64
	switch operator {
		case '+':
			res = n1 + n2
		case '-':
			res = n1 - n2
		case '*':
			res = n1 * n2
		case '/':
			res = n1 / n2
	}
	return res
}

// func Cal(){
// }//同一个包中不能有相同的函数名、全局变量名