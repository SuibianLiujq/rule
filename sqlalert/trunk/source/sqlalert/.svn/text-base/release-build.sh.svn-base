#!/bin/bash

PREFIX=$1

function build() {
	make
}


function release() {
	if [ "${PREFIX}" != "" ]; then
		mkdir -p ${PREFIX}/bin
		cp ./bin/sqlalert ${PREFIX}/bin
	fi
}


function run() {
	build
	release
}

run
