#!/bin/bash

SVNHOST="svn://192.168.0.66"
SVNROOT="${SVNHOST}/src/alert"

PATH_SRC="source"
PATH_PWD=`pwd`

CFG_CLEAN=".release.clean"
CFG_MODULES=".release.modules"

declare -A MODULES
MODULES["sqlalert"]="${SVNROOT}/sqlalert/tags/1.0.3"
MODULES["rules"]="${SVNROOT}/sqlalert-rules/tags/1.0.0"

function check_svn_path() {
	printf "Checking svn paths ..."
	for name in ${SVNROOT} ${MODULES[@]}; do
		if [ ! "$?" = "0" ]; then
			echo "ERROR: \"${name}\" not found!"
			exit 1
		fi
	done
	printf "\rChecking svn paths ... [OK]\n"
}

function export_svn_codes() {
	for name in ${!MODULES[@]}; do
		printf "Exporting ${MODULES[${name}]} ..."
		svn export ${MODULES[${name}]} ${name} > /dev/null
		printf "\rExporting ${MODULES[${name}]} ... [OK]\n"
	done
}

function output_info() {
	echo "rm -rf ${PATH_SRC}"     > ${CFG_CLEAN}
	echo "rm -rf ${CFG_CLEAN}"   >> ${CFG_CLEAN}
	echo "rm -rf ${CFG_MODULES}" >> ${CFG_CLEAN}

	echo "MODULES  = ${!MODULES[@]}"  > ${CFG_MODULES}
	echo "SRC      = ${PATH_SRC}"    >> ${CFG_MODULES}
}

function run() {
	check_svn_path
	rm -rf ${PATH_SRC} > /dev/null 2>&1

	mkdir ${PATH_SRC}
	cd ${PATH_SRC}
	export_svn_codes
	cd ${PATH_PWD}

	output_info
}

run
