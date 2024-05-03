# SitMQFilter

* You can download latest IBM MQ Client from: https://public.dhe.ibm.com/ibmdl/export/pub/software/websphere/messaging/mqdev/redist/


* as of this date, this version is used: 9.3.5.1. I have only extracted: inc,lib64 folders extracted, dont need other bullshit

* Tested on IBM mq docker version: 9.3.5.1-r1, 

```bash
docker pull icr.io/ibm-messaging/mq:9.3.5.1-r1
```

To run it:

* access webconsole on https://localhost:9443/ibmmq/console/login.html, running on port: 1414, webconsole at: 9443, username: admin, password: password
```bash
docker run -e LICENSE=accept -e MQ_ADMIN_PASSWORD=password -e MQ_QMGR_NAME=QM1 -p 1414:1414 -p 9443:9443 icr.io/ibm-messaging/mq:9.3.5.1-r1
```