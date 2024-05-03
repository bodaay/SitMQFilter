# SitMQFilter

* I'm not going to support any other crappy Windows shit, only Linux 64-bit supported

* Just use the script I created to properly get client dist package and prepare the folder ibmmq_dist: GetLatestLibAndRuntime.sh


* Tested on IBM mq docker version: 9.2.4.0-r1, 


To run it:

* access webconsole on https://localhost:9443/ibmmq/console/, running on port: 1414, webconsole at: 9443, username: admin, password: password
```bash
docker run -e MQ_LOGGING_CONSOLE_EXCLUDE_ID="" -e LICENSE=accept -e MQ_ADMIN_PASSWORD=password -e MQ_APP_PASSWORD=password -e MQ_QMGR_NAME=QM1 -p 1414:1414 -p 9443:9443 ibmcom/mq:9.2.4.0-r1-amd64
```

```bash
docker run -e LICENSE=accept -e MQ_ADMIN_PASSWORD=password -e MQ_QMGR_NAME=QM1 -p 1414:1414 -p 9443:9443 ibmcom/mq:9.2.4.0-r1-amd64
```

```bash
docker run -e LICENSE=accept -e MQ_QMGR_NAME=QM1 -p 1414:1414 -p 9443:9443 ibmcom/mq:9.2.4.0-r1-amd64
```

