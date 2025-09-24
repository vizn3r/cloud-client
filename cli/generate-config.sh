#!/bin/bash

read -p "Server host: " SERVER_HOST
echo

mkdir -p $HOME/.cloud

cat <<EOF >$HOME/.cloud/cli.json
{
  "server": {
    "host": "$SERVER_HOST"
  },
}
EOF

echo "Config saved to $HOME/.cloud/cli.json"
