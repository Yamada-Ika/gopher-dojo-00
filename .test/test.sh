#!/bin/bash

function DEBUG_PRINT_ARRAY() {
	array=(${1})
	i=0
	for elem in ${array[@]}; do
		echo "array[$i] = ${elem}" 1>&2
		let i++
	done
}

function DELETE_IMAGE_FILES() {
	image_file_list=(${1})
	for image_file in ${image_file_list[@]}; do
		rm -rf $image_file
	done
}

function ASSERT() {
	res="$1"
	exp="$2"
	if [ "$res" != "$exp" ]
	then
		echo "Assert faile : $res should be $exp" 1>&2
		exit 1
	fi
}

function TEST_ERROR_CASE() {
	echo -e "\nmandatory error test : START\n"

	echo "run : ../convert"
	ASSERT "$(../convert 2>&1)" "error: invalid argument"
	echo "run : ../convert nosuchdirectory"
	ASSERT "$(../convert nosuchdirectory 2>&1)" "error: nosuchdirectory: no such file or directory"
	echo "run : ../convert images/dummys/test1.jpg"
	../convert images/dummys/test1.jpg
	echo "run : ../convert images/dummys/test2.png"
	../convert images/dummys/test2.png
	echo "run : ../convert images/dummys/test3.gif"
	../convert images/dummys/test3.gif
	echo "run : ../convert images/dummys/test4.jpg"
	../convert images/dummys/test4.jpg
	echo "run : ../convert images/dummys/test5.png"
	../convert images/dummys/test5.png
	echo "run : ../convert images/dummys/test6.gif"
	../convert images/dummys/test6.gif
	echo "run : ../convert images/dummys/test7.jpg"
	../convert images/dummys/test7.jpg
	echo "run : ../convert images/dummys/test8.png"
	../convert images/dummys/test8.png
	echo "run : ../convert images/dummys/test9.gif"
	../convert images/dummys/test9.gif
	echo "run : ../convert images/dummys/test10.txt"
	../convert images/dummys/test10.txt
	echo "run : ../convert images/dummys/test11"
	../convert images/dummys/test11
	echo "run : ../convert images/dummys/gif.gif"
	../convert images/dummys/gif.gif
	echo "run : ../convert images/dummys/jpg.jpg"
	../convert images/dummys/jpg.jpg
	echo "run : ../convert images/dummys/png.png"
	../convert images/dummys/png.png
	echo "run : ../convert images/dummys/.gif.gif"

	echo -e "\nmandatory error test : OK\n"
}

function PRINT_IMAGE_LIST_TO_BE_CONVERTED() {
	echo -n "image to be converted : "
	file_list=(${1})
	for file in ${file_list[@]}; do
		echo -n "$file "
	done
	echo ""
}

function TEST_MANDATORY() {
	echo -e "\nmandatory test : START\n"

	# ファイルのリスト
	declare -a jpg_list=()
	declare -a png_list=()

	# imageディレクトリからjpgファイルのリストを作成
	for file in `find images`
	do
		if [ $(file $file | awk '{print $2}') = "JPEG" ]
		then
			jpg_list+=($file)
		fi
	done

	PRINT_IMAGE_LIST_TO_BE_CONVERTED "${jpg_list[*]}"

	# pngファイルのリストを作成
	for elem in ${jpg_list[@]}; do
		png_file=$(echo $elem | sed -e "s/\.[^.]*$/.png/")
		png_list+=($png_file)
	done

	echo "run : ../convert images"
	../convert images > /dev/null 2>&1

	# jpegファイルがpngファイルに変換されているか
	for elem in ${jpg_list[@]}; do
		png_file=$(echo $elem | sed -e "s/\.[^.]*$/.png/")
		ls $png_file > /dev/null 2>&1
		ASSERT $? 0
		ASSERT "$(file $elem | awk '{print $2}')" "JPEG"
	done

	# 作成されたファイルを削除
	DELETE_IMAGE_FILES "${png_list[*]}"

	echo -e "\nmandatory test : OK\n"
}

function TEST_EXT1_TO_EXT2() {
	in_file_ext="$1"
	out_file_ext="$2"

	# ファイルのリスト
	declare -a in_file_list=()
	declare -a out_file_list=()

	# imageディレクトリから各ファイルのリストを作成
	for file in `find images`
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

	PRINT_IMAGE_LIST_TO_BE_CONVERTED "${in_file_list[*]}"

	# pngファイルのリストを作成
	for elem in ${in_file_list[@]}; do
		out_file_list+=($(echo $elem | sed -e "s/\.[^.]*$/.$out_file_ext/"))
	done

	echo "run : ../convert -i=$in_file_ext -o=$out_file_ext images"
	../convert -i=$in_file_ext -o=$out_file_ext images > /dev/null 2>&1

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
	DELETE_IMAGE_FILES "${out_file_list[*]}"
}

function TEST_BONUS() {
	echo "bonus test : START"

	TEST_EXT1_TO_EXT2 "jpg" "png"
	TEST_EXT1_TO_EXT2 "jpg" "gif"
	TEST_EXT1_TO_EXT2 "png" "jpg"
	TEST_EXT1_TO_EXT2 "png" "gif"
	TEST_EXT1_TO_EXT2 "gif" "jpg"
	TEST_EXT1_TO_EXT2 "gif" "png"

	echo "bonus test : OK"
}

if [ $# -eq 0 ]
then
	TEST_MANDATORY
	TEST_ERROR_CASE
elif [ "$@" = "bonus" ]
then
	TEST_BONUS
	TEST_ERROR_CASE
else
	echo "Usage: bash $0 [bonus]"
fi
