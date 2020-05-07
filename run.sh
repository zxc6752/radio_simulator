#!/bin/bash
exe() { echo "\$ $@" ; "$@" ; }

TUN="rantun"
# LISTEN_ADDR="127.0.0.2"
# UE_SUBNET="60.60.0.1/16"
SERVER_SUBNET="60.60.1.0/24"
GTPNL_PATH=lib/libgtp5gnl/tools

[ -n "$1" ] && SERVER_SUBNET=$1

PID_LIST=()

# Setup tunnel interface 
killall -9 gtp5g-link
./${GTPNL_PATH}/gtp5g-link add ${TUN} --ran &
PID_LIST+=($!)
sleep 0.2

# Add MTU on tun dev
ip l s dev ${TUN} mtu 1500

# Add Route to tunnel interface 
ip r add ${SERVER_SUBNET} dev ${TUN}

# Setup tunnel interface 
# exe sudo ip tunnel add ${TUN} mode ipip remote ${LISTEN_ADDR} local 127.0.0.1
# exe sudo ip link set ${TUN} up
# exe sudo ip addr add ${UE_SUBNET} peer ${SERVER_SUBNET} dev ${TUN}

./bin/simulator & 
PID_LIST+=($!)


function terminate()
{
    for ((idx=${#PID_LIST[@]}-1;idx>=0;idx--)); do
        kill -SIGINT ${PID_LIST[$idx]}
    done
   
}


trap terminate SIGINT SIGTERM

echo starting
# commands to start your services go here

wait ${PID_LIST}

# Del tunnel interface
exe ./${GTPNL_PATH}/gtp5g-link del ${TUN}
# exe sudo ip link del ${TUN}

sleep 1

# commands to shutdown your services go here
echo exited
