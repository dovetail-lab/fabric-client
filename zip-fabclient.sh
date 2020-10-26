#!/bin/bash
SCRIPT_DIR=$( dirname "${BASH_SOURCE[0]}" )

cd ${SCRIPT_DIR}
fzip=$(dirname "${PWD}")/fabclientExtension.zip
echo "create file ${fzip}"
cd ..
if [ -f "${fzip}" ]; then
  rm -f ${fzip}
fi

zip -r ${fzip} fabric-client
zip -d ${fzip} fabric-client/.git/\*
