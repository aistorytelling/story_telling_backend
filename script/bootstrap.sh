
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=story_telling_backend
echo "$CURDIR/bin/${BinaryName}"
nohup $CURDIR/bin/${BinaryName}  > out.log 2>&1 &
