#!/bin/bash

for i in *.jpg; \
  do convert $i -resize 582 $(basename $i .jpg)_small.jpg; \
done
