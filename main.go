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
    "encoding/json"                                   
)

var actions = make(map[string][2]int)
type ActionJson struct {
  Action string
  Time int
}

func addAction(actionJson string, ch chan<-string) {

    //actionJson := `{"action": "jump","time":100}`
	var actionJ ActionJson 
	json.Unmarshal([]byte(actionJson), &actionJ)
	fmt.Printf("Parse Input => Action: %s, Time: %d\n", actionJ.Action, actionJ.Time)
    //if err!=nil{
        //fmt.Print("Error:",err)
    //}
    //json.Unmarshal(action, &conf)
    //fmt.Println("%s, %s", conf.action, conf.time)
    
    //fmt.Println(actions[actionJ.Action])
    if actions[actionJ.Action][0] > 0 {
        array := [2]int{actions[actionJ.Action][0] + 1, actions[actionJ.Action][1] + actionJ.Time}
        actions[actionJ.Action] = array
    } else {
        array := [2]int{1, actionJ.Time}
        actions[actionJ.Action] = array           
    }
    ch <- fmt.Sprintf("Action Added: %s, %d, total count: %d, total time: %d", actionJ.Action, actionJ.Time, actions[actionJ.Action][0], actions[actionJ.Action][1])
}

func getStats(ch chan<-string) {
    buf := bytes.Buffer{}      
    buf.WriteString("Get Stats:\n[\n")
    for key, action := range actions {
        buf.WriteString("{\"action\":\"")
        buf.WriteString(key)
        buf.WriteString("\", \"avg\":")
        buf.WriteString(strconv.FormatFloat(float64(action[1])/float64(action[0]),'f',2,64))  // 2 decimal Avg            
        buf.WriteString("},\n")
    }
    buf.WriteString("]")
    result := buf.String()
    
    ch <- fmt.Sprintf(result)
}

func main() {
    ch := make(chan string)
    
    for _,action := range os.Args[1:] {
        go addAction(action, ch)
    }
    for range os.Args[1:] {
        fmt.Println(<-ch)
    }
    
    go getStats(ch)
    fmt.Println(<-ch)
    
    //go run main.go "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\", \"time\":200}"
}