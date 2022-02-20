#!/bin/bash

function DEBUG_PRINT_ARRAY() {
	array=$1
	i=0
	for elem in ${array[@]}; do
		echo "array[$i] = ${elem}" 1>&2
		let i++
	done
}

function DELETE_IMAGE_FILES() {
	image_file_list=$1
	for image_file in ${image_file_list[@]}; do
		rm -rf $image_file
	done
}

function ASSERT() {
	res=$1
	exp=$2
	if [ "$1" != "$2" ]
	then
		echo "Assert faile : $res should be $exp" 1>&2
		exit 1
	fi
}

function TEST_MANDATORY() {
	echo "mandatory test : START"

	cd ..

	# ファイルのリスト
	declare -a jpg_list=()
	declare -a png_list=()

	# imageディレクトリからjpgファイルのリストを作成
	for file in `ls images/*.jpg`
	do
		if [ $(file $file | awk '{print $2}') = "JPEG" ]
		then
			jpg_list+=($file)
		fi
	done

	# pngファイルのリストを作成
	for elem in ${jpg_list[@]}; do
		png_file=$(echo $elem | sed -e "s/\.[^.]*$/.png/")
		png_list+=($png_file)
	done

	echo "run : ./convert images"
	./convert images > /dev/null 2>&1

	# jpegファイルがpngファイルに変換されているか
	for elem in ${jpg_list[@]}; do
		png_file=$(echo $elem | sed -e "s/\.[^.]*$/.png/")
		ls $png_file > /dev/null 2>&1
		ASSERT $? 0
		ASSERT "$(file $elem | awk '{print $2}')" "JPEG"
	done

	# 作成されたファイルを削除
	DELETE_IMAGE_FILES $png_list

	cd .test
	echo "mandatory test : OK"
}

function TEST_EXT1_TO_EXT2() {
	in_file_ext="$1"
	out_file_ext="$2"

	# ファイルのリスト
	declare -a in_file_list=()
	declare -a out_file_list=()

	# imageディレクトリから各ファイルのリストを作成
	for file in `ls images/*.$in_file_ext`
	do
		if [ "$in_file_ext" = "jpg" ]
		then
			if [ $(file $file | awk '{print $2}') = "JPEG" ]
			then
				in_file_list+=($file)
			fi
		elif [ "$in_file_ext" = "gif" ]
		then
			if [ $(file $file | awk '{print $2}') = "GIF" ]
			then
				in_file_list+=($file)
			fi
		elif [ "$in_file_ext" = "png" ]
		then
			if [ $(file $file | awk '{print $2}') = "PNG" ]
			then
				in_file_list+=($file)
			fi
		else
			echo "Error: Invalid image format" 1>&2
			exit 1
		fi
	done

	# pngファイルのリストを作成
	for elem in ${in_file_list[@]}; do
		out_file_list+=($(echo $elem | sed -e "s/\.[^.]*$/.$out_file_ext/"))
	done

	echo "run : ./convert -i=$in_file_ext -o=$out_file_ext images"
	./convert -i=$in_file_ext -o=$out_file_ext images > /dev/null 2>&1

	# jpegファイルがpngファイルに変換されているか
	for elem in ${in_file_list[@]}; do
		out_file=$(echo $elem | sed -e "s/\.[^.]*$/.$out_file_ext/")
		ls $out_file > /dev/null 2>&1
		ASSERT $? 0
		if [ "$in_file_ext" = "jpg" ]
		then
			ASSERT "$(file $elem | awk '{print $2}')" "JPEG"
		elif [ "$in_file_ext" = "gif" ]
		then
			ASSERT "$(file $elem | awk '{print $2}')" "GIF"
		elif [ "$in_file_ext" = "png" ]
		then
			ASSERT "$(file $elem | awk '{print $2}')" "PNG"
		fi
	done

	# 作成されたファイルを削除
	DELETE_IMAGE_FILES $out_file_list
}

function TEST_BONUS() {
	echo "bonus test : START"
	cd ..

	TEST_EXT1_TO_EXT2 "jpg" "png"
	TEST_EXT1_TO_EXT2 "jpg" "gif"
	TEST_EXT1_TO_EXT2 "png" "jpg"
	TEST_EXT1_TO_EXT2 "png" "gif"
	TEST_EXT1_TO_EXT2 "gif" "jpg"
	TEST_EXT1_TO_EXT2 "gif" "png"

	cd .test
	echo "bonus test : OK"
}

if [ $# -eq 0 ]
then
	TEST_MANDATORY
elif [ "$@" = "bonus" ]
then
	TEST_BONUS
else
	echo "Usage: bash $0 [bonus]"
fi
