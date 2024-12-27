package main

import (
	"fmt"
	"os"
	"golang.org/x/term"
)

const ARROW_UP = 65
const ARROW_DOWN = 66
const ARROW_RIGHT = 67
const ARROW_LEFT = 68
var LASTCHAR byte
func moveTo(i int,j int) {
    fmt.Printf("\033[%d;%dH",i+1,j+1)
}
func getKey() (char byte,code byte) {
    fd:=int(os.Stdin.Fd())
    old,err := term.MakeRaw(fd)
    if err!=nil {
        panic(err)
    }
    defer term.Restore(fd,old)
    buf:=make([]byte,1)
    os.Stdin.Read(buf)
    char=buf[0]
    if LASTCHAR==27 && char==91 {
        char=0
        os.Stdin.Read(buf)
        code=buf[0]
    }
    LASTCHAR=char
    if char==27 {
        char=0 
        code=0
    }
    return
}
func editFile(data [][]byte,curpos []int,b bool) ([][]byte,[]int,bool) {
    char,code:=getKey()
    if char=='-' {
        b=false
    } else if char==0 {
        switch code {
            case ARROW_UP:
                curpos[0]--
                break
            case ARROW_DOWN:
                curpos[0]++
                break
            case ARROW_RIGHT:
                curpos[1]++
                break
            case ARROW_LEFT:
                curpos[1]--
                break
        }
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
    } else if char==13 {
        x:=make([][]byte,curpos[0])
        copy(x,data[:curpos[0]])
        x=append(x,data[curpos[0]][:curpos[1]],data[curpos[0]][curpos[1]:])
        data=append(x,data[curpos[0]+1:]...)
        curpos[0]++
        curpos[1]=0
    } else if char==9 {
        line:=make([]byte,curpos[1])
        copy(line,data[curpos[0]][:curpos[1]])
        line=append(line,' ',' ',' ',' ');
        data[curpos[0]] = append(line, data[curpos[0]][curpos[1]:]...)
        curpos[1]+=4
    } else {
        line:=make([]byte,curpos[1])
        copy(line,data[curpos[0]][:curpos[1]])
        line=append(line, char);
        data[curpos[0]] = append(line, data[curpos[0]][curpos[1]:]...)
        curpos[1]++
    }
    return data,curpos,b
}
func Print(b [][]byte,cur []int) {
    fmt.Print("\033[H\033[J")
    for _,i:=range b {
        fmt.Printf("%s\n",i)
    }
    moveTo(cur[0],cur[1])
}
func WriteFile(filename string,data [][]byte) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    for _,i:= range data {
        file.Write(i)
        file.Write([]byte{'\n'})
    }
}
func ReadFile(fileName string) (data [][]byte) {
    d, err := os.ReadFile(fileName)
    if err!=nil {
        panic(err)
    }
    len:=0
    data=[][]byte{{}}
    for _,i := range d {
        if i=='\n' {
            data=append(data,[]byte{})
            len++
        } else {
            data[len]=append(data[len],i)
        }
    }
    return
}

func main() {
    args:=os.Args
    if len(args)!=2 {
        return
    }
    d:=ReadFile(args[1])
    cur:=[]int{0,0}
    editing:=true
    for editing {
        Print(d,cur)
        d,cur,editing=editFile(d,cur,editing)
    }
    WriteFile(args[1],d)
    fmt.Print("\033[H\033[J")
}
