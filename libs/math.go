package libs

func init(){
	Libs["math"] = `fn pow(x,y){
		val := x
		for i :=1; i < y;i++{
			val := val * x
		}
		ret val
	}`
}