package lib

import (
	"fmt"
	"strings"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// QueueMsg is struct created to store Queue Messages
type QueueMsg struct {
	Body string
}

// ReadQueue method returns a list of messages from MSMQ
func ReadQueue(server string, queuetype string, queuename string, search string) []QueueMsg {

	var list []QueueMsg
	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("MSMQ.MSMQQueueInfo")
	msmq, _ := unknown.QueryInterface(ole.IID_IDispatch)

	var formatName string
	if queuetype == "queue" {
		formatName = "direct=os:" + server + "\\private$\\" + queuename
	} else {
		formatName = "direct=os:" + server + "\\private$\\" + queuename + ";Journal"
	}
	oleutil.PutProperty(msmq, "FormatName", formatName)

	MSMQqueue := oleutil.MustCallMethod(msmq, "Open", 32, 0).ToIDispatch()

	// fmt.Println(oleutil.MustGetProperty(MSMQqueue, "IsOpen").Value())
	isOpen := oleutil.MustGetProperty(MSMQqueue, "IsOpen").Value().(int16)

	if isOpen == 1 {
		fmt.Println("Queue is open now....")
	}
	for {
		msg := oleutil.MustCallMethod(MSMQqueue, "PeekCurrent", 0, true, 1000, 0).ToIDispatch()
		if msg != nil {
			msgBody := oleutil.MustGetProperty(msg, "Body").Value().(string)
			// fmt.Println(msgBody)
			subStr := search
			if strings.Contains(msgBody, subStr) {
				list = append(list, QueueMsg{Body: msgBody})
			}
			oleutil.MustCallMethod(MSMQqueue, "PeekNext", 0, true, 1000, 0).ToIDispatch()
		} else {
			oleutil.MustCallMethod(MSMQqueue, "Close")
			fmt.Println("Queue is closed....")
			break
		}
	}

	return list
}
