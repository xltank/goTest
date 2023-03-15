package utils

func Reverse(str string) string {
    list := []rune(str)
    length := len(list)
    for i := range list {
        if(i < length/2){
            list[i], list[length-i-1] = list[length-i-1], list[i]
        }
    }
    return string(list)
}