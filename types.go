package gargs

import "strconv"

func Int(k string) (int, error) {
	return strconv.Atoi(String(k))
}

func Bool(k string) bool {
	return Exists(k) && m[k].(bool)
}

func String(k string) string {
	return m[k].(string)
}
