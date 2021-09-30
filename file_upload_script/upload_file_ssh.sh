#!/bin/bash

from_dir="/home/girish/go/src/github.com/girishg4t/go-grpc-graphql"
to_path="/path/to/whereyouwant/"
user=$1
host=$2

declare -a arrFiles
trap 'echo signal received!; kill "${child_pid}"; wait "${child_pid}"; cleanup' SIGINT
file_index=0

function set_value 
{
  echo "${2}" > "${1}"
}
function read_value 
{
  cat "${1}"
}

transferfile()
{
    echo 'transferring: ' ${arrFiles[$1]} &    
    scp ${arrFiles[$1]} $user@$host:$to_path/${arrFiles[$1]}    
}

store_file_name()
{
    target=$from_dir/$1

    for file in "$target"/*
    do
        arrFiles=(${arrFiles[*]} "$file")
    done

}

store_file_name
arraylength=${#arrFiles[@]}

function start()
{   
    for (( i=0; i<${arraylength}; i++ ));
    do
       set_value "${file_index}" "${i}"
       transferfile ${i}
    done

    echo ""
    echo "Copied: $count"
}

cleanup() 
{
  e=$(read_value "${file_index}")
  e=$((e+1))
  echo "transfering last file ..."
  transferfile ${e}
}

start &

child_pid="$!"
wait "${child_pid}"
