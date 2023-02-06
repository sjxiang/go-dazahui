#!/bin/bash


for n in $@; do
    (sleep $n && echo $n)&
done
wait


# bash sleep_sort.sh 2 3 1