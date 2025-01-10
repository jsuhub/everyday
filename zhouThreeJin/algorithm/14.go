
func fraction(cont []int) []int {
    if len(cont)==1{
        return []int{cont[0],1}
    }
src :=make ([]int,len(cont))
copy(src,cont)
return operateCont(src)

}

func operateCont(cont []int ) []int {
 if(len(cont)==2){
    n := cont[1]*cont[0]+1
   number := []int{n,cont[1]}    //n是分子b是分母
    return  number
 }
 a1 := cont[0]
cont=cont[1:]
m := operateCont(cont)
a2:=a1*m[0]+m[1]
number1:= []int{a2,m[0]}
return number1
}
// a + 1 / b 必定是最简分数，所以不用求GCD了。 （前提：a是整数，b是一个最简分数） 因为b是最简分数，所以 1 / b肯定也是一个最简分数，加上一个整数仍然是最简分数：（ab + 1）/ b = a …… 1