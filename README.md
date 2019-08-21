# Backend Assignment in Golang

Pre-requisition to test this assignment:
1. Having GO installed on your system. Refer to https://golang.org/dl/ to download and intall GO
2. Make sure set GOPATH and test your installation follow the instruction on the download site
3. Clone this repository under your GOPATH

Step to test this assignment:
1. Build exe
Go build

2. Run from main.go
go run main.go "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\",\"time\":200}"
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

3. Run from BackendAssignment.exe
Test Case 1:
BackendAssignment "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\",\"time\":200}"
First getStats Call returns:
[]
Inside addAction => Parsed Input JSON String: {"action":"jump", "time":100}. Action: jump, Time: 100
addAction Call returns No error
Inside addAction => Parsed Input JSON String: {"action":"jump","time":200}. Action: jump, Time: 200
addAction Call returns No error
Inside addAction => Parsed Input JSON String: {"action":"run", "time":75}. Action: run, Time: 75
addAction Call returns No error
Second getStats Call returns:
[{"action":"jump","avg":150},{"action":"run","avg":75}]

Test Case 2: Error in Time

BackendAssignment "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"jump\", \"time\":75}" "{\"action\":\"jump\",\"time\":2s01}"

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