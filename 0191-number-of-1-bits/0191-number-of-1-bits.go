
import "strconv"
func hammingWeight(n int) int {
    var res string 
    var count int 
    for n>0{
        res+=strconv.Itoa(n%2);
        n/=2;
    }
    // fmt.Println(res,len(res))
    for _,val := range res{
        // fmt.Println(i,string(val))
        if string(val)=="1"{
            count++
        }
    }
    return count 
}