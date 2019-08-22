# Backend Assignment in Golang
This is my first assignment writen in Golang 

## Pre-requisition to build and test this assignment:
    
1. Ensure GO is installed on your system. Refer to https://golang.org/dl/ to download and install GO
    
2. Make sure to set GOPATH and test your installation follow the instructions on the download site
    
3. Clone this repository under your GOPATH
    
## Step to build and test this assignment:
    
1. Enter the assignment root folder which includes main.go as current working folder. 
    
2. Build binary, run the follow from command line, it should build an exe file such as BackendAssignment.exe in Windows
```
go build
```

3. Run golang program
### Important 
How to pass multiple serialized JSON strings in command line:
1. Each string needs to be quoted by double quotation marks.     
2. The double quotation marks inside JSON string need to be replaced with ```\"```    

#### Example
```
go run main.go "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\",\"time\":200}"
```
#### Results
The results looks like the follow. The output sequence for addAction may be different since those are concurrent calls.
```
First getStats Call returns:    
[]    
Inside addAction => Parsed Input JSON String: {"action":"run", "time":75}. Action: run, Time: 75    
Inside addAction => Parsed Input JSON String: {"action":"jump", "time":100}. Action: jump, Time: 100    
Inside addAction => Parsed Input JSON String: {"action":"jump","time":200}. Action: jump, Time: 200    
addAction Call returns No error    
addAction Call returns No error    
addAction Call returns No error    
Second getStats Call returns:    
[{"action":"run","avg":75},{"action":"jump","avg":150}] 
```

## Test Cases

#### Test Case 1 - 3 correct JSON inputs:    
```    
go run main.go "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\",\"time\":200}"    
```    
    
#### Test Case 2 - Error in Time, time is not integer:        
```    
go run main.go "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"jump\", \"time\":75}" "{\"action\":\"jump\",\"time\":2s01}"    
```
Results:
```
First getStats Call returns:    
[]    
Inside addAction => Invalid Input JSON String: {"action":"jump","time":2s01}. Error: invalid character 's' after object key:value pair 
addAction Call returns Error: invalid character 's' after object key:value pair    
Inside addAction => Parsed Input JSON String: {"action":"jump", "time":75}. Action: jump, Time: 75    
Inside addAction => Parsed Input JSON String: {"action":"jump", "time":100}. Action: jump, Time: 100    
addAction Call returns No error    
addAction Call returns No error    
Second getStats Call returns:    
[{"action":"jump","avg":87.5}]    
```

#### Test Case 3 - Invalid JSON format:        
```    
go run main.go "{\"action"\:\"jump\", \"time\":100" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\",\"time\":200}"   
```

#### Test Case 4 - No action in JSON:        
```    
go run main.go "{\"act"\:\"jump\", \"time\":100" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\",\"time\":200}"    
```   
