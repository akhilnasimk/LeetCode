func minDeletionSize(strs []string) int {
    count:=0
    for i:=range strs[0]{
        for j:=1;j<len(strs);j++{
            if strs[j][i]<strs[j-1][i]{
                count++
                break 
            }
        }
    }
    return count 
}