#!/bin/sh

initlock="/go/init.lock"
conf=/go/storage.toml

if [ ! -f "$initlock" ]; then
	sh /go/init.sh
fi

/go/agent -config=$conf
