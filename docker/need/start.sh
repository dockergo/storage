#!/bin/sh

initlock="/wpsep/init.lock"
storageconf=/wpsep/storage/storage.toml

if [ ! -f "$initlock" ]; then
	sh /wpsep/init.sh
fi

/wpsep/storage/storage -config=$storageconf
