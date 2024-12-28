package main
import (
    "fmt"
    "os"
)
func main() {
    args:=os.Args
    if len(args)!=2 {
        return
    }
    d:=LoadFile(args[1])
    fmt.Print(OPEN_ALT_BUFFER,HOME,CLEAR)
    cur:=[]int{0,0}
    offset:=[]int{0,0}
    mode:=1
    for mode!=0 {
        if mode==1 {
            fmt.Print(BLINKING_BLOCK_CURSOR)
            Print(d,cur,offset,args[1],mode)
            d,cur,offset,mode=NormalMode(d,cur,offset,mode,args[1])
        } else if mode==2 {
            fmt.Print(BLINKING_THIN_CURSOR)
            Print(d,cur,offset,args[1],mode)
            d,cur,offset,mode=InsertMode(d,cur,offset,mode)
        }
    }
    fmt.Print(CLOSE_ALT_BUFFER,BLINKING_BLOCK_CURSOR)
}
