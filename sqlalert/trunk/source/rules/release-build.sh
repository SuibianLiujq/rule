#!/bin/bash

PREFIX=$1

function release() {
	if [ "${PREFIX}" != "" ]; then
		mkdir -p "${PREFIX}/rules"

		for name in `ls`; do
			if [ -d "${name}" ]; then
				cp -r "./${name}" "${PREFIX}/rules"
			fi
		done
	fi
}

function run() {
	release

    # Delete configurations of YUN Li Lai.
	target="${PREFIX}/rules/globals/yll"
	if [ -d "${target}" ]; then
		rm -r "${target}"
	fi 
}

run
