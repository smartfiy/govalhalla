{
  pkgs ? import <nixpkgs> { },
}: with pkgs;

let

  valhalla = (import ./valhalla) { };
  zlib = pkgs.pkgsStatic.zlib;
  lz4 = pkgs.pkgsStatic.lz4;

in
stdenv.mkDerivation {
  pname = "govalhalla";
  version = "0.0.1";
  src = ./.;

  nativeBuildInputs = [ swig pkg-config valhalla ]; # valhalla is a build-time dependency
  buildInputs = [ swig valhalla go protobuf zlib lz4 boost.dev ]; # These are runtime dependencies

  buildPhase = ''
    # Generate SWIG bindings
    echo "Generating SWIG bindings: ${swig}/bin/swig (version: ${swig.version})"
    echo "Compiler: $CXX"

    ${swig}/bin/swig -c++ -v -go -cgo -intgosize 64 \
      -o govalhalla_wrap.cxx \
      govalhalla.i

    # Compile shared library
    echo "Compiling govalhalla_wrap.cxx..."
    $CXX -fPIC -std=c++17 -c -v govalhalla_wrap.cxx \
        -I${valhalla}/third_party \ # Access third_party correctly
        -I${boost}/include \
        -I${go}/share/go/src/runtime/cgo

    echo "Compiling govalhalla_actor.cpp..."
    $CXX -fPIC -std=c++17 -c govalhalla_actor.cpp \
      -I${boost}/include \
      -I${go}/share/go/src/runtime/cgo

    echo "Linking final shared library..."

    $CXX -shared -fPIC govalhalla_actor.o govalhalla_wrap.o \
      -o libgovalhalla.so \
      -L${valhalla}/lib \
      -L${boost}/lib \
      -L${go}/share/go/src/runtime/cgo \
      -Wl,-Bstatic \
      -lvalhalla \
      -lz \
      -llz4 \
      -lprotobuf \
      -Wl,-Bdynamic \
      -lstdc++ \
      -lgcc_s \
      -lgcc \
      -lpthread

  '';

  installPhase = ''
    mkdir -p $out/{lib,govalhalla}
    cp libgovalhalla.so $out/lib/
    cp govalhalla_wrap.cxx $out/lib/
    cp -r *.go $out/govalhalla/
  '';
}