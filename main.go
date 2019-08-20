package main

/*==============================================================================
*   Date Created:  08/19/2019
*   Description:   Main GO Class for Backend Assignment
*
*   08/10/2019  CH  Created main function   
==============================================================================*/

import (
    "fmt"
    "os" 
    "bytes"   
    "strconv"
)

var actions = make(map[string][2]int)

func addAction(action string, ch chan<-string) {
    fmt.Println(actions[action])
    if actions[action][0] > 0 {
        array := [2]int{actions[action][0] + 1, actions[action][1] + 100}
        actions[action] = array
    } else {
        array := [2]int{1, 100}
        actions[action] = array           
    }
    ch <- fmt.Sprintf("%d %d %s", actions[action][0], actions[action][1], action)
}

func main() {
    ch := make(chan string)
    for _,action := range os.Args[1:] {
        go addAction(action, ch)
    }
    for range os.Args[1:] {
        fmt.Println(<-ch)
    }
    //fmt.Printf("%s elapsed\n", action)
    go getStats(ch)
    fmt.Println(<-ch)
}

func getStats(ch chan<-string) {
    buf := bytes.Buffer{}      
  
    for key, action := range actions {
        buf.WriteString("{\"action\":\"")
        buf.WriteString(key)
        buf.WriteString("\", \"avg\":")
        buf.WriteString(strconv.FormatFloat(float64(action[1])/float64(action[0]+1),'f',2,64))             
        buf.WriteString("},\n")
    }
    result := buf.String()
    ch <- fmt.Sprintf(result)
}

