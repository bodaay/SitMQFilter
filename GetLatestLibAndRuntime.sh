#!/bin/bash

# Built this script, based on: https://github.com/ibm-messaging/mq-golang/blob/master/samples/runSample.deb.Dockerfile
rm -rf ibmmq_dist

runtimeFolder=ibmmq_dist
mkdir -p tmp
mkdir -p $runtimeFolder

cd tmp

MQARCH=X64
RDURL="https://public.dhe.ibm.com/ibmdl/export/pub/software/websphere/messaging/mqdev/redist"
RDTAR="IBM-MQC-Redist-Linux${MQARCH}.tar.gz"
VRMF=9.3.5.0


URL="$RDURL/$VRMF-$RDTAR"

LOCAL_FILE="./$VRMF-$RDTAR"


# Check if the file already exists
if [ -f "$LOCAL_FILE" ]; then
    echo "File already exists: $LOCAL_FILE"
else
    echo "Downloading file: $URL"
    # Use curl to download the file
    curl -o "$LOCAL_FILE" "$URL"
    echo "Download complete."
fi
# the genmqpkg.sh will generate a smaller runtime dist package and delete other non required shit

tar -zxf ./*.tar.gz \
  && bin/genmqpkg.sh -b ../$runtimeFolder

