{ lib
, stdenv
, boost179
, cmake
, fetchFromGitHub
, zlib
, protobuf
, sqlite
, geos
, pkg-config
, lz4
, ...
}:



stdenv.mkDerivation rec {
   name = "valhalla";
  rev= "3.5.1";

  src = fetchFromGitHub {
    inherit rev;
    owner = "valhalla";
    repo = "valhalla";
    # rev = "${rev}";
    sha256 = "sha256-v/EwoJA1j8PuF9jOsmxQL6i+MT0rXbyLUE4HvBHUWDo=";
    fetchSubmodules = true;
  };


  cmakeFlags = [
    "-DENABLE_CCACHE=OFF"
    "-DCMAKE_BUILD_TYPE=Release"
    "-DBUILD_SHARED_LIBS=OFF"
    "-DENABLE_STATIC_LIBRARY_MODULES=ON"
    "-DENABLE_BENCHMARKS=OFF"
    "-DENABLE_PYTHON_BINDINGS=OFF"
    "-DENABLE_TESTS=OFF"
    "-DENABLE_TOOLS=OFF"
    "-DENABLE_SERVICES=OFF"
    "-DENABLE_HTTP=OFF"
    "-DENABLE_DATA_TOOLS=OFF"
  ];

  nativeBuildInputs = [ cmake protobuf ];

  buildInputs = [
    zlib
    boost179
    sqlite
    geos
    pkg-config
    lz4
    protobuf
    
  ];

  # postFixup = ''
  #   patchelf --set-rpath "${lib.makeLibraryPath buildInputs}:$out/lib" $out/lib/libvalhalla.so
  # '';

  # install necessary headers
  # postInstall = ''
  #   mkdir -p $out/{include,proto}
  #   cp -r $src/third_party/rapidjson/include/* $out/include
  #   cp -r $src/third_party/date/include/* $out/include
  #   cp -r $src/proto/* $out/proto
  # '';

}