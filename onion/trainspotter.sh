#!/bin/sh

DIR=$(dirname $0)

shut_off_led() {
   sh "$DIR"/expled-off.sh > /dev/null
}
shut_off_led
trap shut_off_led SIGHUP SIGINT SIGTERM

while read MESSAGE
do
    echo "$MESSAGE"

    STATUS=${MESSAGE%% *}
    BODY=${MESSAGE#* }
    case $STATUS in
        GO)
            expled 0x00ff00 > /dev/null
            ;;
        WAIT)
            expled 0x0000ff > /dev/null
            ;;
        ERROR)
            expled 0xff0000 > /dev/null
            ;;
    esac
done
