#! /bin/sh

output_path=${1:-/tmp/index.md}
yaml_path=$2

echo "[DEBUG] output_path=$output_path"
echo "[DEBUG] yaml_path=$yaml_path"

ino -ino-path $yaml_path -output $output_path

if [ $? -ne 0 ]; then exit $?; fi

echo "::set-output name=output_file::$output_path"