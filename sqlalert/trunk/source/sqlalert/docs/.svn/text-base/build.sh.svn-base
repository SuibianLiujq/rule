#!/bin/bash


TARGET=$1
LANG=$2
LATEX=xelatex
ROOT="src"
OUTPUT="output"
VERSION="1.1.0"
PDFTMP="template-pdf.latex"
FLAGS="--toc --toc-depth=5 -N"
PREFIX="sqlalert"
GENCMD="pandoc --latex-engine=${LATEX} --template=${PDFTMP} ${FLAGS} -V linestretch=1.3 --highlight-style=tango"


function build_dir() {
	target=$1
	path=${ROOT}/${target}
	tmpfile=".temfile.md"

	for lang in `ls ${path}`; do
		if [ "${LANG}" != "" ] && [ "${LANG}" != ${lang} ]; then
			continue
		fi

		rm -f ${tmpfile}
		for file in `ls ${path}/${lang}/*.md 2>/dev/null`; do
			echo ${path}/${lang}/${file}
			cat ${file} >> ${tmpfile}
			echo >> ${tmpfile}
		done

		if [ -f "${tmpfile}" ]; then
			output="${OUTPUT}/${PREFIX}-${target}-${lang}-${VERSION}.pdf"

			echo "Building '${target} (${lang})' ${output} ..." 
			${GENCMD} -o ${output} ${tmpfile}
		fi
	done
}


function build() {
	for dir in `ls ${ROOT}`; do
		if [ "${TARGET}" != "" ] && [ "${TARGET}" != "${dir}" ]; then
			continue
		fi

		build_dir ${dir}
	done
}


build
