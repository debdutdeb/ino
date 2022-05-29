#! /bin/sh

output=${INPUT_output_path:-/tmp/index.md}
yaml_path=$INPUT_ino_path

ino -ino-path $yaml_path -output $output

echo "::set-output name=output_file::$output"