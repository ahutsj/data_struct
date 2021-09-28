package ToOffer

func longestPalindrome(s string) string {
	res := ""
	for i := 0;i< len(s);i ++{
		for l,r := i,i;l >= 0 && r < len(s) && s[l] == s[r];l,r = l-1,r+1 {
			if len(res) < r-l +1{
				res = s[l :r+1]
			}
		}

		for l,r := i,i+1;l >= 0 && r <len(s) && s[l] == s[r];l,r = l-1,r+1 {
			if len(res) < r-l+1 {
				res = s[l:r+1]
			}
		}
	}
	return res
}

