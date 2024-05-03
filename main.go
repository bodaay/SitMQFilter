package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ibm-messaging/mq-golang/v5/ibmmq"
)

func main() {
	QueueManager := "QM1"
	// QueueName := "DEV.QUEUE.1"
	AppChannelName := "DEV.ADMIN.SVRCONN"
	Host := "127.0.0.1"
	Port := "1414"
	Username := "admin"
	Password := "password"
	MQConnStr := fmt.Sprintf("%s(%s)", Host, Port)

	var qMgrName string
	var err error
	var qMgr ibmmq.MQQueueManager
	var rc int
	qMgrName = QueueManager

	// Allocate the MQCNO and MQCD structures needed for the CONNX call.
	cno := ibmmq.NewMQCNO()
	cd := ibmmq.NewMQCD()

	// Fill in required fields in the MQCD channel definition structure
	cd.ChannelName = AppChannelName
	cd.ConnectionName = MQConnStr

	// Reference the CD structure from the CNO and indicate that we definitely want to
	// use the client connection method.
	cno.ClientConn = cd
	cno.Options = ibmmq.MQCNO_CLIENT_BINDING

	// MQ V9.1.2 allows applications to specify their own names. This is ignored
	// by older levels of the MQ libraries.
	cno.ApplName = "Golang ApplName"

	// Also fill in the userid and password if the MQSAMP_USER_ID
	// environment variable is set. This is the same variable used by the C
	// sample programs such as amqsput shipped with the MQ product.

	csp := ibmmq.NewMQCSP()
	csp.AuthenticationType = ibmmq.MQCSP_AUTH_USER_ID_AND_PWD
	csp.UserId = Username

	// For simplicity (it doesn't help with understanding the MQ parts of this program)
	// don't try to do anything special like turning off console echo for the password input
	csp.Password = Password
	// Make the CNO refer to the CSP structure so it gets used during the connection
	cno.SecurityParms = csp

	// And now we can try to connect. Wait a short time before disconnecting.
	qMgr, err = ibmmq.Connx(qMgrName, cno)
	if err == nil {
		fmt.Printf("Connection to %s succeeded.\n", qMgrName)
		d, _ := time.ParseDuration("30s")
		time.Sleep(d)
		qMgr.Disc() // Ignore errors from disconnect as we can't do much about it anyway
		rc = 0
	} else {
		fmt.Printf("Connection to %s failed.\n", qMgrName)
		fmt.Println(err)
		rc = int(err.(*ibmmq.MQReturn).MQCC)
	}

	fmt.Println("Done.")
	os.Exit(rc)
}
