#!/bin/sh

initlock="/flyaways/init.lock"
conf=/flyaways/storage.toml

if [ ! -f "$initlock" ]; then
	sh /flyaways/init.sh
fi

/flyaways/agent -config=$conf
