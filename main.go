package main

/*==============================================================================
*   Date Created:  08/19/2019
*   Description:   Backend Assignment to write a small library class 
*                  that can perform concurrent calls in 2 functions
*
*   08/10/2019  CH  Added main, addAction, getStats function   
*   08/20/2019  CH  addAction return Error
*   08/20/2019  CH  getStats convert struct to JSON String
==============================================================================*/

import (
    "fmt"
    "os" 
    "encoding/json"           
    "errors"    
    "strings"
)

// Map to store num of calls and sum of time by action
// key is action
// value is a length 2 int array: int[0] is num of calls, int[1] is sum of time
var Actions = make(map[string][2]int)      

// Used to parse input JSON string to this Struct for addAction input
type ActionJson struct {
  Action string `json:"action"`
  Time int `json:"time"`
}

// Used to build JSON string from this struct for getStates output
type StatsJson struct {
  Action string `json:"action"`
  Avg float64 `json:"avg"`
}

/*
Input: a JSON serialized string of the form like: {"action":"jump", "time":100}
Output: If Input String invalid, returns Error, Otherwise returns nil  
*/
func addAction(actionString string, ch chan<-error) {

    var actionJson ActionJson 
	err := json.Unmarshal([]byte(actionString), &actionJson)    // Parse Input JSON string into Struct ActionJson
    
    if err != nil {    // if parsing error, return error
        fmt.Printf("Inside addAction => Invalid Input JSON String: %s. Error: %s\n", actionString, err)
        ch <- err
    } else {    // parsing success
        fmt.Printf("Inside addAction => Parsed Input JSON String: %s. Action: %s, Time: %d\n", actionString, actionJson.Action, actionJson.Time)
        
        if actionJson.Action == "" {    // Check if has action, return error if No Action
            err := errors.New("Action is Empty")
            ch <- err
        } else {
          if Actions[actionJson.Action][0] > 0 {    
              // Check if the action is already in map, Num of count plus 1, Sum of time plus the new Time
              array := [2]int{Actions[actionJson.Action][0] + 1, Actions[actionJson.Action][1] + actionJson.Time}
              Actions[actionJson.Action] = array    // Update to map
          } else {
              // If the action is not in map, add a new action key, initial with count = 1 and the new time
              array := [2]int{1, actionJson.Time}
              Actions[actionJson.Action] = array           
          }                                   
          ch <- nil    // No error, return nil
        }
    }
}

/*
Input: Null
Output: a serialized json array of the average time for each action that has been provided to the addAction function  
*/
func getStats(ch chan<-string) {
    
    var rtn = []string{}
    
    for key, sum := range Actions {   
        // Build StatsJson struct for each action        
        statsJson := &StatsJson{Action: key, Avg: float64(sum[1]*100/sum[0])/100}    // Avg time is the sum of time divided by num of count. Keep 2 decimal
        statsString, _ := json.Marshal(statsJson)    // Convert Struct to serialized JSON
        
        rtn = append(rtn, string(statsString))    // Append serialized JSON sting to return string array
    }     
    
    result := strings.Join(rtn, ",")    // Join string array to one string
    
    ch <- "[" + result + "]"
}

/*
Args: multiple JSON serialized string in the below format
For Example: 
go run main.go "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\", \"time\":200}"
*/
func main() {
    ch1 := make(chan error)    // For concurrent addAction calls
    ch2 := make(chan string)    // For concurrent getStats calls
    
    // Make concurrent calls to addAction for each Arg
    for _,action := range os.Args[1:] {
        go addAction(action, ch1)
    }
    
    // Make a concurrent call to getStats before unblock from the previous addAction calls and Print results
    go getStats(ch2)
    fmt.Printf("First getStats Call returns: \n%s\n", <-ch2)
    
    // Print the results from the previous addAction calls
    for range os.Args[1:] {
        err := <-ch1
        if err == nil {
            fmt.Printf("addAction Call returns No error\n")
        } else {
            fmt.Printf("addAction Call returns Error: %s\n", err)
        }
        
    }
    
    // Make a concurrent call to getStats after unblock and Print results
    go getStats(ch2)
    fmt.Printf("Second getStats Call returns: \n%s\n", <-ch2)
}