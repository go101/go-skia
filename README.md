**go-skia** is a Go skia package based on skia c library through cgo.

The project is still in early stage.

The current skia libraries are built from skia C++ repository
revision 8534723c7be1c079711d8bac45e93728d6524f8a, with the
following commands:
```
$ bin/gn gen out/Static --args='is_official_build=true skia_use_dng_sdk=false skia_use_libwebp=false skia_use_freetype=false skia_use_fontconfig=false skia_use_expat=false skia_enable_fontmgr_empty=true skia_use_icu=false extra_cflags = [ "-w" ]'

$ ninja -C out/Static
```
