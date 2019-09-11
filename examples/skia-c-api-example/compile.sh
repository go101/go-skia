#!/bin/bash

gcc ./*.c ../../skia/lib/static/linux/amd64/libskia.* \
	-o skia-c-example \
	-I ../../skia/include/c \
	-Wl,-rpath,../../skia/out/static \
	-lstdc++ -lm -lpng -ljpeg -lz -lpthread
	# -lwebp -lwebpmux -lpng16 -ldl -luuid -lfreetype -lfontconfig -lexpat
