#!/bin/bash
cd ..
if [ -f ./fabclientExtension.zip ]; then
  rm -f ./fabclientExtension.zip
fi
zip -r ./fabclientExtension.zip fabric-client
