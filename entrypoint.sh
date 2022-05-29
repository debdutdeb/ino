#! /bin/sh

output=${INPUT_output_path:-/tmp/index.md}
yaml_path=$INPUT_ino_path

echo "[DEBUG] output_path=$output_path"
echo "[DEBUG] yaml_path=$yaml_path"

ino -ino-path $yaml_path -output $output

if [ $? -ne 0 ]; then exit $?; fi

echo "::set-output name=output_file::$output"