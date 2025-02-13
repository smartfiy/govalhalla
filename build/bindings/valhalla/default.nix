{
  pkgs ? import <nixpkgs> { },
}: with pkgs;


stdenv.mkDerivation (finalAttrs: {
  pname = "valhalla";
  version = "3.5.1";

  

  src = fetchFromGitHub {
    owner = "valhalla";
    repo = "valhalla";
    rev = finalAttrs.version;
    hash = "sha256-v/EwoJA1j8PuF9jOsmxQL6i+MT0rXbyLUE4HvBHUWDo=";
    fetchSubmodules = true;
  };


  nativeBuildInputs = [
    cmake
    pkg-config
    
  ];

  buildInputs = [
    boost179
    protobuf
    lz4
    zlib
    
    # gcc.cc.lib
  ];

  # CXXFLAGS="-fexceptions";
  

  # PKG_CONFIG_PATH = "${pkgs.valhalla}/lib/pkgconfig";
  # "-DENABLE_STATIC_LIBRARY_MODULES=ON"
  cmakeFlags = [
    "-DENABLE_TESTS=OFF"
    "-DENABLE_BENCHMARKS=OFF"
    "-DENABLE_CCACHE=OFF"
    "-DENABLE_SINGLE_FILES_WERROR=OFF"
    "-DENABLE_EXAMPLES=OFF"
    "-DENABLE_PYTHON_BINDINGS=OFF"
    "-DENABLE_TOOLS=OFF"
    "-DENABLE_SERVICES=OFF"
    "-DENABLE_HTTP=OFF"
    "-DENABLE_DATA_TOOLS=OFF"
    "-DCMAKE_BUILD_TYPE=Release"
    "-DCMAKE_POSITION_INDEPENDENT_CODE=ON"
    "-DCMAKE_CXX_STANDARD=17"
    
  ];

 postUnpack = ''
  # Copy files to tyr module
  cp ${../govalhalla_actor.cpp} $sourceRoot/src/tyr/govalhalla_actor.cc
  cp ${../govalhalla_actor.h} $sourceRoot/valhalla/tyr/govalhalla_actor.h

  # Add to sources list
  sed -i '/set(sources/a \ \ govalhalla_actor.cc' $sourceRoot/src/tyr/CMakeLists.txt
  '';

  #  env.CMAKE_CXX_FLAGS = "-DCMAKE_CXX_STANDARD=17";
  # postFixup = ''
  #   substituteInPlace "$out"/lib/pkgconfig/libvalhalla.pc \
  #     --replace '=''${prefix}//' '=/' \
  #     --replace '=''${exec_prefix}//' '=/'
  # '';

  passthru.tests = {
    pkg-config = testers.testMetaPkgConfig finalAttrs.finalPackage;
  };

  postInstall = ''
    cp -r $out/include/valhalla/third_party/* $out/include/
  '';

  meta = with lib; {
    changelog = "https://github.com/valhalla/valhalla/blob/${finalAttrs.src.rev}/CHANGELOG.md";
    description = "Open Source Routing Engine for OpenStreetMap";
    homepage = "https://valhalla.readthedocs.io/";
    license = licenses.mit;
    maintainers = [ maintainers.Thra11 ];
    platforms = platforms.linux;
  };
})
