#!/bin/bash


# test_string="[SESSION] ID=10,DNN=internet,SST=1,SD=010203,UEIP=60.60.0.1,ULAddr=10.200.200.102,ULTEID=2,DLAddr=10.200.200.1,DLTEID=1"
SESS_FORMAT=$'\[SESSION\] ID=([0-9]+),DNN=([^,]+),SST=([0-9]+),SD=([0-9]+),UEIP=([^,]+),ULAddr=([^,]+),ULTEID=([0-9]+),DLAddr=([^,]+),DLTEID=([0-9]+)'
# if [[ $test_string =~ $SESS_FORMAT ]]; then echo "DNN=${BASH_REMATCH[1]},UEIP=${BASH_REMATCH[4]}"; fi

# TUN="uetun"
# TUN_ADDR="60.60.0.1"

function show_usage {
    echo
    echo "Ue Simulator"
    echo "Usage: $0 ip port supi [-k|--keep-alive] [-id=] [-t|--time=] [-s|--slice=]"
    echo
    echo "Arguments:"
    echo "  -k|--keep-alive    : do not deregsiter"
    echo "  -id=               : pdu session id (default=10)"
    echo "  -t|--time=n        : pdu session would exist n[seconds](if not set when script terminate would be released)"
    echo "  -s|--slice=        : specific pdu session slice info, format: \"dnn=%s,sst=%d,sd=%s\""
}

if [ -z "$3" ]
then
    show_usage
    exit 1
fi


HOST=$1
PORT=$2
SUPI=$3
ALIVE=false
TIME=0

shift 3

for i in "$@"
do
case $i in
    -k|--keep-alive)
    ALIVE=true
    shift
    ;;
    -id=*)
    ID="${i#*=}"
    shift
    ;;
    -t=*|--time=*)
    TIME=${i#*=}
    shift # past argument with no value
    ;;
    -s=*|--slice)
    SLICE="${i#*=}"
    shift
    ;;
esac
done


function check_error() {
    if [[ "$1" == *"[ERROR]"* ]] && [[ "$1" == *"FAIL"* ]]; then
        exit 1
    fi
}
send_msg() { 
    echo "\$ $1" 
    echo "$1" >&$2
}
read_msg() {
    read -r msg_in <&$1
    check_error "$msg_in"
    echo "$msg_in"
}
get_ueip(){
    if [[ $1 =~ $SESS_FORMAT ]]
    then 
        echo "${BASH_REMATCH[5]}"
    fi 
}

exec 3<>/dev/tcp/${HOST}/${PORT}

read -r msg_in <&3
echo $msg_in
# send SUPI

send_msg "$SUPI" 3
read_msg 3

# Register
send_msg "reg" 3
read_msg 3

# ADD Session
msg_out="sess 10 add"
[ -n "$ID" ] && msg_out="sess $ID add"
[ -n "$SLICE" ] && msg_out="$msg_out"" ${SLICE}"
send_msg "$msg_out" 3
msg_in=$(read_msg 3)
echo "$msg_in"

# Add Ip in tun dev
# UEIP=$(get_ueip "$msg_in")
# if [ -n "${UEIP}" ] && [ "${UEIP}" != ${TUN_ADDR} ]
# then
#     sudo ip addr add ${UEIP} dev ${TUN}
# fi


if [ $TIME -gt 0 ]
then 
    echo "Wait $TIME seconds"
    sleep $TIME &
else 
    sleep infinity & 
fi

function terminate(){
    if $ALIVE;
    then
        # send rel pdu sess
        send_msg "$(echo "${msg_out}" | sed -e "s/add/del/g")" 3
        read_msg 3
    else 
        # send del reg
        send_msg "dereg" 3
        read_msg 3
    fi
    exit 1
}

trap terminate SIGINT
wait 

terminate
