#!/bin/bash

dir="/home/test/go/src/github.com/girishg4t/go-grpc-graphql"
user=$1
host=$2

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
    scp ${arrFiles[$1]} $user@$host:/path/to/whereyouwant/${arrFiles[$1]}    
}

declare -a arrFiles
store_file_name()
{
    target=$dir/$1

    for file in "$target"/*
    do
        arrFiles=(${arrFiles[*]} "$file")
    done

}

trap 'echo signal received!; kill "${child_pid}"; wait "${child_pid}"; cleanup' SIGINT
store_file_name
arraylength=${#arrFiles[@]}
file_index=0
function execute()
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
  # Our cleanup code goes here
}

execute &

child_pid="$!"
wait "${child_pid}"
