#!/bin/bash

function DEBUG_PRINT_ARRAY() {
	array=$1
	i=0
	for e in ${array[@]}; do
		echo "array[$i] = ${e}" 1>&2
		let i++
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

echo "mandatory test : START"

cd ..

# ファイルのリスト
declare -a jpg_list=()
declare -a png_list=()

# jpgファイルのリストを作成
for file in `ls images/*.jpg`
do
	if [ $(file $file | awk '{print $2}') = "JPEG" ]
	then
		jpg_list+=($file)
	fi
done

# pngファイルのリストを作成
i=0
for elem in ${jpg_list[@]}; do
	png_file=$(echo $elem | sed -e "s/\.[^.]*$/.png/")
	png_list+=($png_file)
done

echo "run : ./convert images"
./convert images > /dev/null 2>&1

# jpegファイルがpngファイルに変換されているか
# 変換されたファイルが存在するかどうか
# 変換されたファイルがJPEGかどうか
i=0
for elem in ${jpg_list[@]}; do
	png_file=$(echo $elem | sed -e "s/\.[^.]*$/.png/")
	ls $png_file > /dev/null 2>&1
	ASSERT $? 0
	ASSERT "$(file $elem | awk '{print $2}')" "JPEG"
done

# 作成されたファイルを削除
for elem in ${png_list[@]}; do
	rm -rf $elem
done

echo "mandatory test : OK"
