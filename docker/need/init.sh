#!/bin/bash
storageconf=/wpsep/storage/storage.toml

if [ ! -n "$AUTH_ACCESSKEY" ]; then
sed -i "s,AUTH_ACCESSKEY,1GL02rRYQxK8s7FQh8dV,g" $storageconf
sed -i "s,AUTH_SECRETKEY,2IDjaPOpFfkq5Zf9K4tKu8k5AKApY8S8eKV1zsRl,g" $storageconf
sed -i "s,HTTP_ADDR,:8080,g" $storageconf
sed -i "s,HTTP_SCHEME,http,g" $storageconf
sed -i "s,LOGS_ACCESS,false,g" $storageconf
sed -i "s,LOGS_PATH,,g" $storageconf
sed -i "s,LOGS_LEVEL,trace,g" $storageconf
else
sed -i "s,AUTH_ACCESSKEY,$AUTH_ACCESSKEY,g" $storageconf
sed -i "s,AUTH_SECRETKEY,$AUTH_SECRETKEY,g" $storageconf
sed -i "s,HTTP_ADDR,:$HTTP_ADDR,g" $storageconf
sed -i "s,HTTP_SCHEME,$HTTP_SCHEME,g" $storageconf
sed -i "s,LOGS_ACCESS,$LOGS_ACCESS,g" $storageconf
sed -i "s,LOGS_PATH,$LOGS_PATH,g" $storageconf
sed -i "s,LOGS_LEVEL,$LOGS_LEVEL,g" $storageconf
fi

sed -i "s,STORAGE_TYPE,$STORAGE_TYPE,g" $storageconf

sed -i "s,POSIX_LOCAL_PATH,$POSIX_LOCAL_PATH,g" $storageconf

sed -i "s,S3_ACCESSKEY,$S3_ACCESSKEY,g" $storageconf
sed -i "s,S3_SECRETKEY,$S3_SECRETKEY,g" $storageconf
sed -i "s,S3_ADDR,$S3_ADDR,g" $storageconf

sed -i "s,OSS_ACCESSKEY,$OSS_ACCESSKEY,g" $storageconf
sed -i "s,OSS_SECRETKEY,$OSS_SECRETKEY,g" $storageconf
sed -i "s,OSS_ADDR,$OSS_ADDR,g" $storageconf

sed -i "s,KDFS_ACCOUNT,$KDFS_ACCOUNT,g" $storageconf
sed -i "s,KDFS_ADDR,$KDFS_ADDR:$KDFS_ADDR_PATH,g" $storageconf

sed -i "s,SWIFT_ADDR,$SWIFT_ADDR,g" $storageconf
sed -i "s,SWIFT_AUTHURL,http://$SWIFT_AUTHURL:$SWIFT_AUTHURL_PORT_PATH,g" $storageconf
sed -i "s,SWIFT_TENANTNAME,$SWIFT_TENANTNAME,g" $storageconf
sed -i "s,SWIFT_USERNAME,$SWIFT_USERNAME,g" $storageconf
sed -i "s,SWIFT_PROJNAME,$SWIFT_PROJNAME,g" $storageconf
sed -i "s,SWIFT_PASSWORD,$SWIFT_PASSWORD,g" $storageconf

touch /wpsep/init.lock

