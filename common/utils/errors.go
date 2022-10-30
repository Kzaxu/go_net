package utils

func Check(cond bool, err string) {
	if cond {
		panic(err)
	}
}
