func maxFreqSum(s string) int {
    vowel:=make(map[string]int)
    cons:=make(map[string]int)
    for _,val:=range s{
        if string(val)=="a" || string(val)=="e" || string(val)=="i" || string(val)=="o" || string(val)=="u"{
            vowel[string(val)]++
        }else{
            cons[string(val)]++
        }
    }
    fmt.Println(vowel,"  ",cons)
    BV:=0
    for _,val :=range vowel{
        if val>BV{
            BV=val
        }
    }
    CV:=0
    for _,val :=range cons{
        if val>CV{
            CV=val
        }
    }
    return BV+CV
}