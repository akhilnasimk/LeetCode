func isValid(s string) bool {
    if len(s)==1{
        return false 
    }
    stack:=[]rune{}
    for _,val:=range s {
        if string(val)=="{" || string(val)=="(" || string(val)=="["{
            stack= append(stack,val)
        }else{
            if len(stack) == 0 {
				return false
			}
            last:=stack[len(stack)-1]
            stack=stack[:len(stack)-1]
            switch string(last) {
                 case "(":
                    if string(val)!=")"{
                        return false 
                    }
                case "{" :
                    if string(val)!="}"{
                        return false 
                    }
                case "[":
                    if string(val)!="]"{
                        return false 
                    }
                default :
                    continue     
            }    
            } 
    }
    return len(stack) == 0 
}