{
  pkgs ? import <nixpkgs> { },
}: with pkgs;

let
  valhalla = (import ./valhalla) { };
  
  pkgunstable = import (builtins.fetchGit {
    url = "https://github.com/NixOS/nixpkgs";
    ref = "nixos-unstable";
    rev = "9abb87b552b7f55ac8916b6fc9e5cb486656a2f3";        #"666e1b3f09c267afd66addebe80fb05a5ef2b554"; "nixos-24.11";
  }) { };

  swig = pkgunstable.swig;
  protobuf = pkgunstable.protobuf;

  zlib = pkgsStatic.zlib;
  

in

stdenv.mkDerivation {
  pname = "govalhalla"; #"govalhalla-${valhallaCustom.rev}";
  version = "0.0.1";
  src = ./.;

  nativeBuildInputs =  [ pkgunstable.swig pkgs.pkg-config pkgs.boost.dev valhalla ];
  buildInputs =  [
    pkgunstable.swig
    valhalla
    pkgs.go
    pkgunstable.protobuf
    zlib 
    pkgs.boost.dev
  ];

  # preBuildPhase = ''
  #   export PKG_CONFIG_PATH=${pkgs.valhalla}/lib/pkgconfig:$PKG_CONFIG_PATH
  # '';

  buildPhase = ''


    # export CC="${stdenv.cc.targetPrefix}gcc"
    # export CXX="${stdenv.cc.targetPrefix}g++"
    

    # Generate SWIG bindings
    echo "Generating SWIG bindings: ${swig}/share/swig/${swig.version}"
    echo "Compiler : $CXX"
    

  ${swig}/bin/swig -c++ -v -go -cgo -intgosize 64 \
    -cpperraswarn \
    -o govalhalla_wrap.cxx \
    -I${valhalla}/include \
    -I${protobuf}/include \
    -I${boost.dev}/include \
    govalhalla.i
 
    
    

    # Compile shared library
    echo "wrapping shared library cd src ....."
    

    $CXX -fPIC -std=c++17 -c -v govalhalla_wrap.cxx \
        -I. \
        -I${valhalla}/include \
        -I${valhalla}/include/valhalla/third_party \
        -I${protobuf}/include \
        -I${boost.dev}/include \
        -I${go}/share/go/src/runtime/cgo \
        -I${valhalla}/include/valhalla/proto

    echo "Compiling govalhalla_actor.cpp..."

    $CXX -fPIC -std=c++17 -c govalhalla_actor.cpp \
      -I. \
      -I${valhalla}/include \
      -I${valhalla}/include/valhalla/third_party \
      -I${protobuf}/include \
      -I${boost.dev}/include \
      -I${go}/share/go/src/runtime/cgo \
      -I${valhalla}/include/valhalla/proto

    echo "Linking final shared library..."

    ls -l

    $CXX -shared  -fPIC  govalhalla_actor.o govalhalla_wrap.o  \
      -o libgovalhalla.so \
      -L${valhalla}/lib \
      -L${boost.dev}/include \
      -L${go}/share/go/src/runtime/cgo \
      -Wl,-Bstatic \
      -lvalhalla \
      -lz \
      -Wl,-Bdynamic \
      -lpthread \
      -Wl,--verbose
  '';

  installPhase = ''
    mkdir -p $out/{lib,govalhalla}
    cp libgovalhalla.so $out/lib/
    cp govalhalla_wrap.cxx $out/lib/
    cp -r *.go $out/govalhalla/

    # cp govalhalla_wrap.cxx $out/govalhalla/
    # cp -r ${valhalla}/include/proto/*.go $out/valhalla/
  '';


}
