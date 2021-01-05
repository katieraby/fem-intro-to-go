package main

func average (arg1, arg2, arg3 float64) (averageOfNums float64) {
	averageOfNums = (arg1 + arg2 + arg3) / 3
	return
}

func main() {
	average(6.1, 1.2, 3.3)
}