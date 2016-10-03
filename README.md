# MSMQReader-Golang
This Project is about reading messages from MSMQ using Golang 

Sample cmd to exec the exe file with params :

MSMQReader -server yourservername -queuename yourqueuename -journal -search 27080619  -write E:\\output.txt 


Mandatory fields :

* server
* queuename

Optional fields:

* journal  (This flag ensures the MSMQ messages are read from journal if flag is not present it will read main queue)
* search   (This flag will check for any strings we need to search in MSMQ and return only the messages contains the string from queue, if            no search flag given it will return all messages from queue)
* write    (This flag is to write the output to specific file, if not present the output will just print in cmd line)


