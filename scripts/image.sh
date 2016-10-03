#!/bin/bash

while getopt "u:d" opt; do
  case $opt in
    u)
      URL=$OPTARG
      ;;
    d)
      DEST=$OPTARG
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      exit 1
      ;;
  esac
done

dir=`dirname $DEST`
base=`basename $DEST`
suffix

wget -O ${DEST?} $URL \
	&& convert $DEST -resize 582 ${base}_small
