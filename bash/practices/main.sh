#!/bin/bash
#this is commenting
echo "This script is running in $(pwd)"
name="Sara"
echo "Hello ${name}"
time=$(date +%H:%m:%S)
echo "Time is ${time}"
echo "Using Array"
Array=(one two three four "This is five")
echo "${Array[0]}"
echo "${Array[4]}"
echo "the first argument is $1 the second is $2"
exit 0

