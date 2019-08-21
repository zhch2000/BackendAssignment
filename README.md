# Backend Assignment in Golang
    
## Pre-requisition to build and test this assignment:
    
1. Having GO installed on your system. Refer to https://golang.org/dl/ to download and intall GO
    
2. Make sure set GOPATH and test your installation follow the instruction on the download site
    
3. Clone this repository under your GOPATH
    
## Step to build and test this assignment:
    
1. enter the assignment root folder which includes main.go as current working folder. 
    
2. Build binary, run the follow from command line, it should build an exe file such as BackendAssignment.exe in windows   
```    
go build
```

3. Please refer to Test Cases.txt for the correct command line to test this assignment
```
go run main.go "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\",\"time\":200}"
```
