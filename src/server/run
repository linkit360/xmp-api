#!/usr/bin/env bash
clear
go build -o ../../bin/acceptor
if [ $? != 0 ]; then
    printf "\n\n BUILD FAIL \n\n"
    exit 1
fi
../../bin/acceptor
