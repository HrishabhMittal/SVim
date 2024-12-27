package main
import (
    "fmt"
    "os"
    "golang.org/x/term"
)
var LASTCHAR byte
func getKey() (char byte,code byte) {
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
func insertMoveCur(code byte,curpos []int,offset []int,data [][]byte) ([]int,[]int) {
    fd:=int(os.Stdin.Fd())
    w,h,err:=term.GetSize(fd)
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
var savepos int = 0
func insertMode(data [][]byte,curpos []int,offset []int,b int) ([][]byte,[]int,[]int,int) {
    char,code:=getKey()
    if char==27 {
        b=1
    } else if char==0 {
        curpos,offset=insertMoveCur(code,curpos,offset,data)
    } else if char==127 {
        if curpos[1]==0 {
            if curpos[0]!=0 {
                a:=len(data[curpos[0]-1])
                data[curpos[0]-1]=append(data[curpos[0]-1],data[curpos[0]]...)
                data=append(data[:curpos[0]],data[curpos[0]+1:]...)
                curpos[1]=a
                curpos[0]--
            }
        } else {
            data[curpos[0]]=append(data[curpos[0]][:curpos[1]-1],data[curpos[0]][curpos[1]:]...)
            curpos[1]--
        }
        savepos=curpos[1]
    } else if char==13 {
        x:=make([][]byte,curpos[0])
        copy(x,data[:curpos[0]])
        x=append(x,data[curpos[0]][:curpos[1]],data[curpos[0]][curpos[1]:])
        data=append(x,data[curpos[0]+1:]...)
        curpos[0]++
        curpos[1]=0
        savepos=curpos[1]
    } else if char==9 {
        line:=make([]byte,curpos[1])
        copy(line,data[curpos[0]][:curpos[1]])
        line=append(line,' ',' ',' ',' ');
        data[curpos[0]] = append(line, data[curpos[0]][curpos[1]:]...)
        curpos[1]+=4
        savepos=curpos[1]
    } else {
        line:=make([]byte,curpos[1])
        copy(line,data[curpos[0]][:curpos[1]])
        line=append(line, char);
        data[curpos[0]] = append(line, data[curpos[0]][curpos[1]:]...)
        curpos[1]++
        savepos=curpos[1]
    }
    return data,curpos,offset,b
}
func Print(b [][]byte,cur []int,offset []int) {
    fd:=int(os.Stdin.Fd())
    w,h,err:=term.GetSize(fd)
    if err!=nil {
        panic(err)
    }
    fmt.Print(INVIS_CURSOR,HOME)
    for i:=offset[0];i<len(b)&&i<offset[0]+h;i++ {
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
    moveTo(cur[0]-offset[0],cur[1]-offset[1])
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
func normalMoveCur(code byte,curpos []int,offset []int,data [][]byte) ([]int,[]int) {
    fd:=int(os.Stdin.Fd())
    w,h,err:=term.GetSize(fd)
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
    } else if curpos[1]>len(data[curpos[0]])-1 {
        curpos[1]=len(data[curpos[0]])-1
        if curpos[1]==-1 {
            curpos[1]=0
        }
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

func compare(bytearr []byte,str string) (bool) {
    if len(bytearr)!=len(str) {
        return false
    }
    for i,e:=range bytearr {
        if e!=str[i] {
            return false
        }
    }
    return true
}
var cmd []byte = []byte{}
func normalExecute(data [][]byte,filename string,b int) (int) {
    if compare(cmd,":w") {
        WriteFile(filename,data)
    } else if compare(cmd,":q") {
        b=0
    } else if compare(cmd,":wq") {
        WriteFile(filename,data)
        b=0
    }
    cmd=[]byte{}
    return b
}
func NormalMode(data [][]byte,curpos []int,offset []int,b int,filename string) ([][]byte,[]int,[]int,int) {
    char,code:=getKey()
    if len(cmd)!=0 {
        if char==13 {
            b=normalExecute(data,filename,b)
        } else if char==127 {
            cmd=cmd[:len(cmd)-1]
        } else if char!=0 {
            cmd=append(cmd,char)
        }
    } else if char=='i' {
        b=2
    } else if char=='a' {
        curpos,offset=insertMoveCur(ARROW_RIGHT,curpos,offset,data)
        b=2
    } else if char==':' {
        cmd=[]byte{':'}
    } else {
        if char=='h' {
            curpos,offset=normalMoveCur(ARROW_LEFT,curpos,offset,data)
        } else if char=='j' {
            curpos,offset=normalMoveCur(ARROW_DOWN,curpos,offset,data)
        } else if char=='k' {
            curpos,offset=normalMoveCur(ARROW_UP,curpos,offset,data)
        } else if char=='l' {
            curpos,offset=normalMoveCur(ARROW_RIGHT,curpos,offset,data)
        } else {
            curpos,offset=normalMoveCur(code,curpos,offset,data)
        }
    }
    return data,curpos,offset,b
}
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
            Print(d,cur,offset)
            d,cur,offset,mode=NormalMode(d,cur,offset,mode,args[1])
        } else if mode==2 {
            fmt.Print(BLINKING_THIN_CURSOR)
            Print(d,cur,offset)
            d,cur,offset,mode=insertMode(d,cur,offset,mode)
        }
    }
    fmt.Print(CLOSE_ALT_BUFFER,BLINKING_BLOCK_CURSOR)
}
