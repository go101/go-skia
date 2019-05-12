**go-skia** is a Go skia binding based on skia C library through cgo.

Note: the project is still in early stage, and it only supports Linux now. 
The next release (in developing) will be much more mature.

If your project supports Go modules, then the import path is `go101.org/go-skia`,
and please remember add the following line in your `go.mod` file manually.
(`go get github.com/go-gfx/go-skia` doesn't work.)

```
replace go101.org/go-skia => github.com/go-gfx/go-skia v0.0.1
```

The current skia static library is built from [the skia fork](https://github.com/go-graphics/skia),
which forked from skia C++ repository revision 8534723c7be1c079711d8bac45e93728d6524f8a.
The commands used to build the library:
```
$ bin/gn gen out/Static --args='is_official_build=true skia_use_libwebp=false skia_use_dng_sdk=false skia_use_freetype=false skia_use_fontconfig=false skia_use_expat=false skia_enable_fontmgr_empty=true skia_use_icu=false skia_use_system_libjpeg_turbo=false skia_use_system_libpng=false skia_use_system_zlib=false extra_cflags = [ "-w" ]'
$
$ ninja -C out/Static
```



