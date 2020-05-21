package main

func reversString(s string) string  {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(s string) bool {
	var low, high int
	high = len(s) - 1
	for low < high {
		if s[low] != s[high] {
			return false
		}
		low ++
		high --
	}
	return true
}

func longestPalindrome(s string) string {
	ans := ""
	//fmt.Println(s)
	rs := reversString(s)
	//fmt.Println(rs)
	for i := 1; i < len(s); i ++ {
		//fmt.Println(s[:i])
		//fmt.Println(rs[len(rs) - i:])
		ts := s[:i]
		trs := rs[len(rs) - i:]

		tans := ""
		tempt := ""
		for j := 0; j < len(ts); j ++ {
			if ts[j] == trs[j] {
				tempt += string(ts[j])
			} else {
				if len(tempt) > len(tans) {
					tans = tempt
				}
				tempt = ""
			}
		}
		if len(tempt) > len(tans) {
			tans = tempt
		}
		if len(tans) > len(ans) {
			ans = tans
		}
	}
	for i := 1; i < len(s); i ++ {
		//fmt.Println(s[:i])
		//fmt.Println(rs[len(rs) - i:])

		ts := s[i:]
		trs := rs[:len(rs) - i]
		tans := ""
		tempt := ""
		for j := 0; j < len(ts); j ++ {
			if ts[j] == trs[j] {
				tempt += string(ts[j])
			} else {
				if len(tempt) > len(tans) {
					tans = tempt
				}
				tempt = ""
			}
		}
		if len(tempt) > len(tans) {
			tans = tempt
		}
		if len(tans) > len(ans) {
			ans = tans
		}
	}
	tempt := ""
	for j := 0; j < len(s); j ++ {
		if s[j] == rs[j] {
			tempt += string(s[j])
		} else {
			if len(tempt) > len(ans) && isPalindrome(tempt) {
				ans = tempt
			}
			tempt = ""
		}
	}
	if len(tempt) > len(ans) && isPalindrome(tempt) {
		ans = tempt
	}
	if len(ans) == 0 && len(s) > 0 {
		return string(s[0])
	}
	return ans
}

func main() {
	//fmt.Println(longestPalindrome("babad"))
	//fmt.Println(longestPalindrome("cbbd"))
	//fmt.Println(longestPalindrome("aacdefcaa"))
	//fmt.Println(longestPalindrome("esbtzjaaijqkgmtaajpsdfiqtvxsgfvijpxrvxgfumsuprzlyvhclgkhccmcnquukivlpnjlfteljvykbddtrpmxzcrdqinsnlsteonhcegtkoszzonkwjevlasgjlcquzuhdmmkhfniozhuphcfkeobturbuoefhmtgcvhlsezvkpgfebbdbhiuwdcftenihseorykdguoqotqyscwymtjejpdzqepjkadtftzwebxwyuqwyeegwxhroaaymusddwnjkvsvrwwsmolmidoybsotaqufhepinkkxicvzrgbgsarmizugbvtzfxghkhthzpuetufqvigmyhmlsgfaaqmmlblxbqxpluhaawqkdluwfirfngbhdkjjyfsxglsnakskcbsyafqpwmwmoxjwlhjduayqyzmpkmrjhbqyhongfdxmuwaqgjkcpatgbrqdllbzodnrifvhcfvgbixbwywanivsdjnbrgskyifgvksadvgzzzuogzcukskjxbohofdimkmyqypyuexypwnjlrfpbtkqyngvxjcwvngmilgwbpcsseoywetatfjijsbcekaixvqreelnlmdonknmxerjjhvmqiztsgjkijjtcyetuygqgsikxctvpxrqtuhxreidhwcklkkjayvqdzqqapgdqaapefzjfngdvjsiiivnkfimqkkucltgavwlakcfyhnpgmqxgfyjziliyqhugphhjtlllgtlcsibfdktzhcfuallqlonbsgyyvvyarvaxmchtyrtkgekkmhejwvsuumhcfcyncgeqtltfmhtlsfswaqpmwpjwgvksvazhwyrzwhyjjdbphhjcmurdcgtbvpkhbkpirhysrpcrntetacyfvgjivhaxgpqhbjahruuejdmaghoaquhiafjqaionbrjbjksxaezosxqmncejjptcksnoq"))
}
