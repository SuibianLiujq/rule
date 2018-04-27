#!/bin/bash

TARGET=$1

CFG_CLEAN=".release.clean"
CFG_MODULES=".release.modules"

SH_BUILD="release-build.sh"
PATH_PWD=`pwd`
PATH_TMP="release"

MODULES=""
SRC=""

PREFIX=""
PACKAGE=""

function load_cfg() {
	echo `sed "/^$1 *=/!d;s/.*=//" $2`
}


function load() {
	MODULES=`load_cfg "MODULES" ${CFG_MODULES}`
	SRC=`load_cfg "SRC" ${CFG_MODULES}`

	PREFIX="${PATH_PWD}/${PATH_TMP}"
	mkdir -p ${PREFIX}
	echo "rm -rf ${PREFIX}" >> ${CFG_CLEAN}
}


function build() {
	for name in ${MODULES}; do
		if [ ! ${TARGET} = "" ]; then
			if [ ! ${TARGET} = ${name} ]; then
				continue
			fi
		fi

		cd ${name}
			echo "Building ${name} ..."
			sh ${SH_BUILD} ${PREFIX} > /dev/null
		cd ..
	done
}


function packege() {
	printf "Packaging ..."
	if [ ! -x "${PREFIX}/bin/sqlalert" ]; then
		echo "Executable file 'sqlalert' not found"
		exit 1
	fi

	version=`${PREFIX}/bin/sqlalert -v | awk '{print $2}'`
	PACKAGE="sqlalert-${version}"

	cp -r ${PATH_TMP} ${PACKAGE}
	tar zcvf "${PACKAGE}.tar.gz" ${PACKAGE}/* > /dev/null 2>&1
	echo "rm -rf ${PACKAGE}.tar.gz" >> ${CFG_CLEAN}

	rm -rf ${PACKAGE}
	printf "\rPackaging ... [OK]\n"
}

function run() {
	if [ "${TARGET}" = "clean" ]; then
		sh ${CFG_CLEAN}
		exit
	fi

	if [ ! -f "${CFG_MODULES}" ]; then
		echo "File ${CFG_MODULES} not found, please run 'make co' first!"
		exit 1
	fi

	load
	echo "Building ${TARGET} ..."
	echo

	cd ${SRC}
		build
	cd ${PATH_PWD}

	packege
}

run
