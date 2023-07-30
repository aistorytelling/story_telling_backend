
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=story_telling_backend
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
