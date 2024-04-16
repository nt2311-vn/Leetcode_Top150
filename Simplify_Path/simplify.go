package simplifypath

func simplifyPath(path string) string {
	var stack []string
	var res string
	var i int
	for i < len(path) {
		if path[i] == '/' {
			i++
			continue
		}
		start := i
		for i < len(path) && path[i] != '/' {
			i++
		}
		substr := path[start:i]
		if substr == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if substr != "." {
			stack = append(stack, substr)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	for _, v := range stack {
		res += "/" + v
	}
	return res
}
