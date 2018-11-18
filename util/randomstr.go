package util

import "math/rand"

/*
|--------------------------------------------------------------------------
| Function main()
|--------------------------------------------------------------------------
| @author 	: Thomas
| @return 	: string
| @RandStringBytes 	: set uuid data
|
*/

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
