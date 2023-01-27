package main


func main() {


}


func LongestSubstring(word string) {
	var longest string
	var count int

	for i:=0; i < len(word); i++ {

	
	}
}


func FindCharacter(s string, sub string) bool{
	flag := false
	for i:= 0; i< len(sub); i++ {
		if string(sub[i]) == s{
			flag = true
			return flag
		}
	}
	return flag
}