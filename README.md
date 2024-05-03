# SitMQFilter

# You can download latest IBM MQ Client from: https://public.dhe.ibm.com/ibmdl/export/pub/software/websphere/messaging/mqdev/redist/


# as of this date, this version is used: 9.2.0.0, lib64 foolder only extracted

# Tested on IBM mq docker version: 9.2.4.0-r1-amd64, 

```
docker pull ibmcom/mq:9.2.4.0-r1-amd64
```

To run it:

``` access webconsole on https://localhost:9443/ibmmq/console/login.html, username: admin, password: password
docker run -e LICENSE=accept -e MQ_ADMIN_PASSWORD=password -p 9443:9443 ibmcom/mq:9.2.4.0-r1-amd64
```