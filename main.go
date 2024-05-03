package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ibm-messaging/mq-golang/v5/ibmmq"
)

func main() {
	QueueManager := "QM1"
	QueueName := "DEV.QUEUE.1"
	AppChannelName := "DEV.ADMIN.SVRCONN" //DEV.ADMIN.SVRCONN, if you want to use this, username should be app, and the password same as what I set in docker MQ_APP_PASSWORD
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
		d, _ := time.ParseDuration("3s")
		time.Sleep(d)
		// qMgr.Disc() // Ignore errors from disconnect as we can't do much about it anyway
		// rc = 0
		mqod := ibmmq.NewMQOD()
		var openOptions int32
		openOptions = ibmmq.MQOO_OUTPUT + ibmmq.MQOO_FAIL_IF_QUIESCING
		openOptions |= ibmmq.MQOO_INPUT_AS_Q_DEF

		mqod.ObjectType = ibmmq.MQOT_Q
		mqod.ObjectName = QueueName
		var qObject ibmmq.MQObject
		qObject, err = qMgr.Open(mqod, openOptions)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Opened queue", qObject.Name)
		}
		if err == nil {
			putmqmd := ibmmq.NewMQMD()
			pmo := ibmmq.NewMQPMO()

			pmo.Options = ibmmq.MQPMO_SYNCPOINT | ibmmq.MQPMO_NEW_MSG_ID | ibmmq.MQPMO_NEW_CORREL_ID

			putmqmd.Format = "MQSTR"
			msgData := "Hello from Go at " + time.Now().Format("02 Jan 2006 03:04:05")
			buffer := []byte(msgData)

			err = qObject.Put(putmqmd, pmo, buffer)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Put message to", qObject.Name)
			}
		}

		// The message was put in syncpoint so it needs
		// to be committed.
		if err == nil {
			err = qMgr.Cmit()
			if err != nil {
				fmt.Println(err)
			}
		}
		// MQGET all messages on the queue. Wait 3 seconds for any more
		// to arrive.
		if err == nil {
			msgAvail := true

			for msgAvail == true {
				var datalen int

				getmqmd := ibmmq.NewMQMD()
				gmo := ibmmq.NewMQGMO()
				gmo.Options = ibmmq.MQGMO_NO_SYNCPOINT | ibmmq.MQGMO_FAIL_IF_QUIESCING
				gmo.Options |= ibmmq.MQGMO_WAIT
				gmo.WaitInterval = 3000
				buffer := make([]byte, 32768)

				datalen, err = qObject.Get(getmqmd, gmo, buffer)

				if err != nil {
					msgAvail = false
					fmt.Println(err)
					mqret := err.(*ibmmq.MQReturn)
					if mqret.MQRC == ibmmq.MQRC_NO_MSG_AVAILABLE {
						// not a real error so reset err
						err = nil
					}
				} else {
					fmt.Printf("Got message of length %d: ", datalen)
					fmt.Println(strings.TrimSpace(string(buffer[:datalen])))
				}
			}
		}

		// MQCLOSE the queue
		if err == nil {
			err = qObject.Close(0)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Closed queue")
			}
		}

	} else {
		fmt.Printf("Connection to %s failed.\n", qMgrName)
		fmt.Println(err)
		rc = int(err.(*ibmmq.MQReturn).MQCC)
	}

	fmt.Println("Done.")
	os.Exit(rc)
}
