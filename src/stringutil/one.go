/*
# date 2020-01-08 16:44:44
# author calllivecn <c-all@qq.com>
*/

package stringutil

func Reverse(s string) string {
	var r []rune

    r_rune := []rune(s)

	l := len(r_rune)

	for i := l-1; i >= 0; i-- {
        r = append(r, rune(r_rune[i]))
	}

	return string(r)
}
