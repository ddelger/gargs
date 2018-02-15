package gargs

var m = make(map[string]interface{})

func Load(args map[string]interface{}) {
	for k, v := range args {
		m[k] = v
	}
}

func Exists(k string) bool {
	return m[k] != nil
}
