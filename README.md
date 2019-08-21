# Backend Assignment in Golang

Pre-requisition to build and test this assignment:

1. Having GO installed on your system. Refer to https://golang.org/dl/ to download and intall GO

2. Make sure set GOPATH and test your installation follow the instruction on the download site

3. Clone this repository under your GOPATH

Step to build and test this assignment:

1. enter the assignment root folder which includes main.go as current working folder. 

2. Build exe, run the follow from command line, it should build an exe file such as BackendAssignment.exe in windows   

go build

3. Run from main.go with 3 JSON input strings, run the following from command line

go run main.go "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\",\"time\":200}"

4. The results looks like the follow. The output sequence for addAction may be different since those are concurrent calls     
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

4. Run from Execution file such as BackendAssignment.exe with the follow from command line    

BackendAssignment "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\",\"time\":200}"    

5. Test cases (Working in Progress)    
Test Case 1:    

go run main.go "{\"action\":\"jump\", \"time\":100}" "{\"action\":\"run\", \"time\":75}" "{\"action\":\"jump\",\"time\":200}"    

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
