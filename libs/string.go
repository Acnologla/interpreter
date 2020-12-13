package libs

func init(){
	Libs["string"] = `fn stringIncludes(strArr, str){
		for i:=0; i < len(strArr);i++{
			if strArr[i] == str{
				ret true
			}
		}
		false
	}
	fn stringSlice(str, start, end){
		finalStr := ""
		for i:= start; i < end;i++{
			finalStr = finalStr + str[i]
		}
		finalStr
	}
	
	fn isChar(str){
		if str == null{
			ret false
		}
		charlist := ["a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z","_"]
		stringIncludes(charlist, toLower(str))
	}`
}