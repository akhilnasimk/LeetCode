func sortSentence(s string) string {
    slice := []string{}
    tillnow := ""

    for _, val := range s {
        if val == ' ' {
            slice = append(slice, tillnow)
            tillnow = ""
        } else {
            tillnow += string(val)
        }
    }
    slice = append(slice, tillnow) 

    res := make([]string, len(slice))

    for _, val := range slice {
        index, _ := strconv.Atoi(string(val[len(val)-1])) 
        res[index-1] = val[:len(val)-1]                    
    }

    return strings.Join(res, " ")
}
