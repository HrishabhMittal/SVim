package main
var cmd []byte = []byte{}
func normalExecute(data [][]byte,filename string,b int) (int) {
    if string(cmd)==":w" {
        WriteFile(filename,data)
    } else if string(cmd)==":q" {
        b=0
    } else if string(cmd)==":wq" {
        WriteFile(filename,data)
        b=0
    }
    cmd=[]byte{}
    return b
}
func NormalMode(data [][]byte,curpos []int,offset []int,b int,filename string) ([][]byte,[]int,[]int,int) {
    char,code:=GetKey()
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
        curpos,offset=MoveCur(ARROW_RIGHT,curpos,offset,data)
        b=2
    } else if char==':' {
        cmd=[]byte{':'}
    } else {
        if char=='h' {
            curpos,offset=MoveCur(ARROW_LEFT,curpos,offset,data)
        } else if char=='j' {
            curpos,offset=MoveCur(ARROW_DOWN,curpos,offset,data)
        } else if char=='k' {
            curpos,offset=MoveCur(ARROW_UP,curpos,offset,data)
        } else if char=='l' {
            curpos,offset=MoveCur(ARROW_RIGHT,curpos,offset,data)
        } else {
            curpos,offset=MoveCur(code,curpos,offset,data)
        }
    }
    return data,curpos,offset,b
}



