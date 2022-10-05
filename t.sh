#!/bin/bash
DIR=$(cd $(dirname $0) && pwd)
echo "DIR $DIR"
cd $DIR
pwd
source venv/bin/activate
read -p "Press [Enter] key to start backup..."