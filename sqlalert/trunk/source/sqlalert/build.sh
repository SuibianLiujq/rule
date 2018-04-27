#!/bin/bash

PWD=`pwd`
GO="/opt/go/bin/go"

if [ "${GOPATH}" == "" ]; then
	GOPATH="${PWD}"
else
	GOPATH="${PWD}:${GOPATH}"
fi
export GOPATH=${GOPATH}

${GO} $1 main/sqlalert
${GO} $1 main/test
