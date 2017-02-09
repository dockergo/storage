#!/bin/bash

conf=/go/storage.toml

sed -i "s,AUTH_ACCESSKEY,$AUTH_ACCESSKEY,g" $conf
sed -i "s,AUTH_SECRETKEY,$AUTH_SECRETKEY,g" $conf

sed -i "s,HTTP_ADDR,:$HTTP_ADDR,g" $conf
sed -i "s,HTTP_SCHEME,$HTTP_SCHEME,g" $conf

sed -i "s,LOGS_ACCESS,$LOGS_ACCESS,g" $conf
sed -i "s,LOGS_PATH,$LOGS_PATH,g" $conf
sed -i "s,LOGS_LEVEL,$LOGS_LEVEL,g" $conf

sed -i "s,STORAGE_TYPE,$STORAGE_TYPE,g" $conf

sed -i "s,POSIX_PATH,$POSIX_PATH,g" $conf

sed -i "s,S3_ACCESSKEY,$S3_ACCESSKEY,g" $conf
sed -i "s,S3_SECRETKEY,$S3_SECRETKEY,g" $conf
sed -i "s,S3_ADDR,$S3_ADDR,g" $conf

sed -i "s,OSS_ACCESSKEY,$OSS_ACCESSKEY,g" $conf
sed -i "s,OSS_SECRETKEY,$OSS_SECRETKEY,g" $conf
sed -i "s,OSS_ADDR,$OSS_ADDR,g" $conf

sed -i "s,QINIU_ACCESSKEY,$QINIU_ACCESSKEY,g" $conf
sed -i "s,QINIU_SECRETKEY,$QINIU_SECRETKEY,g" $conf
sed -i "s,QINIU_ADDR,$QINIU_ADDR,g" $conf

sed -i "s,SWIFT_ADDR,$SWIFT_ADDR,g" $conf
sed -i "s,SWIFT_AUTHURL,http://$SWIFT_AUTHURL:$SWIFT_AUTHURL_PORT_PATH,g" $conf
sed -i "s,SWIFT_TENANTNAME,$SWIFT_TENANTNAME,g" $conf
sed -i "s,SWIFT_USERNAME,$SWIFT_USERNAME,g" $conf
sed -i "s,SWIFT_PROJNAME,$SWIFT_PROJNAME,g" $conf
sed -i "s,SWIFT_PASSWORD,$SWIFT_PASSWORD,g" $conf

touch /go/init.lock

