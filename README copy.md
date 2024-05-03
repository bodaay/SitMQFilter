# SitMQFilter

* Just use the script I created to properly get client dist package and prepare the folder ibmmq_dist: GetLatestLibAndRuntime.sh


* Tested on IBM mq docker version: 9.3.5.1-r1, 

```bash
docker pull icr.io/ibm-messaging/mq:9.3.5.1-r1
```

To run it:

* access webconsole on https://localhost:9443/ibmmq/console/, running on port: 1414, webconsole at: 9443, username: admin, password: password
```bash
docker run -e LICENSE=accept -e MQ_ADMIN_PASSWORD=password -e MQ_APP_PASSWORD=password -e MQ_QMGR_NAME=QM1 -p 1414:1414 -p 9443:9443 icr.io/ibm-messaging/mq:9.3.5.1-r1
```

```bash
docker run -e LICENSE=accept -e MQ_QMGR_NAME=QM1 -p 1414:1414 -p 9443:9443 icr.io/ibm-messaging/mq:9.3.5.1-r1
```