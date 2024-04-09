package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	commonPrefix := strs[0]

	for i := 0; i < len(commonPrefix); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != commonPrefix[i] {
				return commonPrefix[:i]
			}
		}
	}

	return commonPrefix
}
