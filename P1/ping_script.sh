#!/bin/bash

dst=""

if [ "$2" == "node1" ]
then
    dst="172.0.0.2"
elif [ "$2" == "node2" ]
then
    dst="172.0.0.3"
elif [ "$2" == "node3" ]
then
    dst="10.10.0.2"
elif [ "$2" == "node4" ]
then
    dst="10.10.0.3"
else [ "$2" == "router" ]
    if [ "$1" == "node1" ] || [ "$1" == "node2" ]
    then
        dst="172.0.0.1"
    else [ "$1" == "node3" ] || [ "$1" == "node4" ]
        dst="10.10.0.1"
    fi
fi

#echo $dst
echo "$1" is pinging "$dst"
ip netns exec $1 ping $dst