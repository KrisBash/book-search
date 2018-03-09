#!/bin/bash
run=0
if [ -e "/tmp/initialize.sh" ]
then
    diff=`diff ./BuildEnv/initialize.sh /tmp/initialize.sh |wc -l`
    if [ $diff -gt 0 ]
    then    
        run=1
    fi
else
    run=1
fi

if [ $run -eq 1 ]
then
    #install ansible
    sudo apt-get update
    sudo apt-get install software-properties-common
    sudo apt-add-repository ppa:ansible/ansible
    sudo apt-get update
    sudo apt-get install -y ansible

fi
cp ./BuildEnv/initialize.sh /tmp/initialize.sh 

