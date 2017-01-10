#!/bin/bash
agentconf=/wpsep/agent/agent.toml

if [ ! -n "$AUTH_ACCESSKEY" ]; then
sed -i "s,AUTH_ACCESSKEY,1GL02rRYQxK8s7FQh8dV,g" $agentconf
sed -i "s,AUTH_SECRETKEY,2IDjaPOpFfkq5Zf9K4tKu8k5AKApY8S8eKV1zsRl,g" $agentconf
sed -i "s,HTTP_ADDR,:20808,g" $agentconf
sed -i "s,HTTP_SCHEME,http,g" $agentconf
sed -i "s,LOGS_ACCESS,false,g" $agentconf
sed -i "s,LOGS_PATH,,g" $agentconf
sed -i "s,LOGS_LEVEL,trace,g" $agentconf
else
sed -i "s,AUTH_ACCESSKEY,$AUTH_ACCESSKEY,g" $agentconf
sed -i "s,AUTH_SECRETKEY,$AUTH_SECRETKEY,g" $agentconf
sed -i "s,HTTP_ADDR,:$HTTP_ADDR,g" $agentconf
sed -i "s,HTTP_SCHEME,$HTTP_SCHEME,g" $agentconf
sed -i "s,LOGS_ACCESS,$LOGS_ACCESS,g" $agentconf
sed -i "s,LOGS_PATH,$LOGS_PATH,g" $agentconf
sed -i "s,LOGS_LEVEL,$LOGS_LEVEL,g" $agentconf
fi

sed -i "s,STORAGE_TYPE,$STORAGE_TYPE,g" $agentconf

sed -i "s,POSIX_LOCAL_PATH,$POSIX_LOCAL_PATH,g" $agentconf

sed -i "s,NFS_LOCAL_PATH,$NFS_LOCAL_PATH,g" $agentconf
sed -i "s,NFS_REMOTE_ADDR,$NFS_REMOTE_ADDR:$NFS_REMOTE_PATH,g" $agentconf

sed -i "s,S3_ACCESSKEY,$S3_ACCESSKEY,g" $agentconf
sed -i "s,S3_SECRETKEY,$S3_SECRETKEY,g" $agentconf
sed -i "s,S3_ADDR,$S3_ADDR,g" $agentconf

sed -i "s,KDFS_ACCOUNT,$KDFS_ACCOUNT,g" $agentconf
sed -i "s,KDFS_ADDR,$KDFS_ADDR:$KDFS_ADDR_PATH,g" $agentconf

sed -i "s,SWIFT_ADDR,$SWIFT_ADDR,g" $agentconf
sed -i "s,SWIFT_AUTHURL,http://$SWIFT_AUTHURL:$SWIFT_AUTHURL_PORT_PATH,g" $agentconf
sed -i "s,SWIFT_TENANTNAME,$SWIFT_TENANTNAME,g" $agentconf
sed -i "s,SWIFT_USERNAME,$SWIFT_USERNAME,g" $agentconf
sed -i "s,SWIFT_PROJNAME,$SWIFT_PROJNAME,g" $agentconf
sed -i "s,SWIFT_PASSWORD,$SWIFT_PASSWORD,g" $agentconf

touch /wpsep/init.lock

