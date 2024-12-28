package main
func InsertMode(data [][]byte,curpos []int,offset []int,b int) ([][]byte,[]int,[]int,int) {
    char,code:=GetKey()
    if char==27 {
        b=1
    } else if char==0 {
        curpos,offset=MoveCur(code,curpos,offset,data)
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

