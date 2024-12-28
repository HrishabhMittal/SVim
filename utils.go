package main
import (
    "fmt"
    "os"
    "golang.org/x/term"
)
func GetDigits(num int) (out int) {
    if (num==0) {
        return 0
    }
    for num!=0 {
        out++
        num/=10
    }
    return out
}
func GetLineNoSpace(total int) (int) {
    total=GetDigits(total)
    if total<3 {
        total=3
    }
    total++
    return total
}
func GetKey() (char byte,code byte) {
    fd:=int(os.Stdin.Fd())
    old,err := term.MakeRaw(fd)
    if err!=nil {
        panic(err)
    }
    defer term.Restore(fd,old)
    buf:=make([]byte,5)
    n,_:=os.Stdin.Read(buf)
    if n==1 {
        char=buf[0]
        code=0 
        return
    } else if n==3 && buf[1]=='[' {
        char=0
        code=buf[2]
    } else {
        char=0 
        code=0
    }
    return
}
func abs(num int) (int) {
    if num<0 {
        return -num
    }
    return num
}
func MoveCur(code byte,curpos []int,offset []int,data [][]byte) ([]int,[]int) {
    fd:=int(os.Stdin.Fd())
    w,h,err:=term.GetSize(fd)
    w-=GetLineNoSpace(len(data))
    h-=2
    if err!=nil {
        panic(err)
    }
    switch code {
        case ARROW_UP:
            curpos[0]--
            break
        case ARROW_DOWN:
            curpos[0]++
            break
        case ARROW_RIGHT:
            curpos[1]++
            savepos=curpos[1]
            break
        case ARROW_LEFT:
            curpos[1]--
            savepos=curpos[1]
            break
    }
    curpos[1]=savepos
    if curpos[0]<0 {
        curpos[0]=0 
    } else if curpos[0]>len(data)-1 {
        curpos[0]=len(data)-1
    }
    if curpos[1]<0 {
        curpos[1]=0 
    } else if curpos[1]>len(data[curpos[0]]) {
        curpos[1]=len(data[curpos[0]])
    }
    if curpos[0]<offset[0] {
        offset[0]=curpos[0]
    } else if curpos[0]>offset[0]+h-4 {
        offset[0]=curpos[0]-h+4
    }
    if curpos[1]<offset[1] {
        offset[1]=curpos[1]
    } else if curpos[1]>offset[1]+w-4 {
        offset[1]=curpos[1]-w+4
    }
    return curpos,offset
}
func GetModeName(num int) (string) {
    if num==0 {
        return "EXIT"
    } else if num==1 {
        return "NORMAL"
    } else if num==2 {
        return "INSERT"
    } else {
        return "INVALID"
    } 
}
var savepos int = 0
func Print(b [][]byte,cur []int,offset []int,filename string,mode int) {
    fd:=int(os.Stdin.Fd())
    w,h,err:=term.GetSize(fd)
    l:=GetLineNoSpace(len(b))
    w-=l
    h-=2
    if err!=nil {
        panic(err)
    }
    fmt.Print(INVIS_CURSOR,HOME)
    for i:=offset[0];i<len(b)&&i<offset[0]+h;i++ {
        if i==cur[0] {
            fmt.Printf("%d%*s",i+1,l-GetDigits(i+1),"")
        } else {
            fmt.Printf("%*d ",l-1,abs(i-cur[0]))
        }
        if offset[1]>=len(b[i]) {
            fmt.Printf("%*s",w,"")
        } else if offset[1]+w>len(b[i]) {
            fmt.Printf("%s",b[i][offset[1]:])
            fmt.Printf("%*s",w-len(b[i])+offset[1],"")
        } else {
            fmt.Printf("%s",b[i][offset[1]:offset[1]+w])
        }
        if i!=len(b)-1&&i!=offset[0]+h-1 {
            fmt.Printf("\n")
        }
    }
    fmt.Print("\033[J")
    moveTo(h,0)
    SetBGColor(8)
    fmt.Printf("%-*s\n",w+l,filename)
    fmt.Print(RESET_COLOR)
    fmt.Printf("-- %s --",GetModeName(mode))
    
    moveTo(cur[0]-offset[0],cur[1]-offset[1]+l)
    fmt.Print(VIS_CURSOR)
}
func WriteFile(filename string,data [][]byte) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    for x,i:= range data {
        file.Write(i)
        if x!=len(data)-1 {
            file.Write([]byte{'\n'})
        }
    }
}
func LoadFile(fileName string) (data [][]byte) {
    d, err := os.ReadFile(fileName)
    len:=0
    data=[][]byte{{}}
    if err!=nil {
        if os.IsNotExist(err) {
            return
        }
        panic(err)
    }
    for _,i := range d {
        if i=='\n' {
            data=append(data,[]byte{})
            len++
        } else if i==9 {
            data[len]=append(data[len],' ',' ',' ',' ')
        } else {
            data[len]=append(data[len],i)
        }
    }
    return
}

