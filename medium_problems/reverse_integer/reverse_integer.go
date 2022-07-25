package reverse_integer

func Reverse(x int) int {
    num := x
    final :=0
    if x< 0{
        num *=-1
    }
    
    for num >0{
        final = (final *10 )+ (num%10)
        num /=10
    }
    
    if final > 1<<31{
        final = 0
    }
    
    if x < 0{
        final *=-1
    }
    
    return final
}
