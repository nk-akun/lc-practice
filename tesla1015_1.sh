#!/usr/bin/env bash

LOG_FILE="$1"

tail -f $LOG_FILE|while read line
do
    is_error=`echo $line|grep "500"|wc -l`
    if [ "$is_error" == "1" ] ; then
        code=`echo $line|awk -F ' ' '{print $9}'`
        file_path=`echo $line|awk -F ' ' '{print $7}'`
        mail "alert@project.com" "HTTP $code on $file_path"
    fi
done