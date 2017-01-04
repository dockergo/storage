#!/bin/sh

initlock="/wpsep/init.lock"
agentconf=/wpsep/agent/agent.toml

if [ ! -f "$initlock" ]; then
	sh /wpsep/init.sh
fi

/wpsep/agent/agent -config=$agentconf
