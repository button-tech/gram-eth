#!/bin/bash

if [ -d "./ssh_key/id_rsa" ]
then
    echo "Directory ./ssh_key exists."
else
    mkdir ./ssh_key
fi

cat ~/.ssh/sendgrid_rsa > ./ssh_key/id_rsa

chmod 600 ./ssh_key/id_rsa

docker build -t krboktv/send-eth .