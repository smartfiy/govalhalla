{
  pkgs ? import <nixpkgs> { },
}: with pkgs;

let

  valhalla = (import ./valhalla) { };
  

in
stdenv.mkDerivation {
  pname = "govalhalla";
  version = "0.0.1";
  src = ./.;
  

  nativeBuildInputs = [ swig pkg-config  ]; # valhalla is a build-time dependency
  buildInputs = [ valhalla protobuf zlib.static lz4 boost179 ]; # These are runtime dependencies

  # NIX_CXXFLAGS_COMPILE = "-stdlib=libc++ -nostdinc++";
  # NIX_CXXFLAGS_COMPILE = "--stdlib=libc++";
  # NIX_LDFLAGS = "-stdlib=libc++ -lc++abi";
  # CXXFLAGS="-fexceptions";

  buildPhase = ''
    # Generate SWIG bindings
    mkdir -p $out/{lib,govalhalla}
    echo "Generating SWIG bindings: ${swig}/bin/swig (version: ${swig.version})"
    echo "Compiler: $CXX"



    ${swig}/bin/swig -c++ -v -go -cgo -intgosize 64 \
      -I${valhalla}/include \
      -module govalhalla \
      -soname libgovalhalla.so \
      -o govalhalla_wrap.cxx \
      govalhalla.i

    echo "Generated govalhalla_wrap.cxx"
    ls -lh

    # Compile shared library
    echo "Compiling govalhalla_wrap.cxx..."
    # CXXFLAGS="-nostdinc++ -isystem $(clang -print-resource-dir)/include/c++/v1"


    # $CXX -fPIC  -c -std=c++17 $CXXFLAGS -v govalhalla_wrap.cxx \
    #   -I${libcxx.dev}/include/c++/v1 \
    #   -o govalhalla_wrap.o 

    echo "Generated govalhalla_wrap.o"
    ls -lh
        

    # echo "Compiling govalhalla_actor.cpp..."
    # $CXX -fPIC -v -c  govalhalla_actor.cpp -o govalhalla_actor.o  \
    #   `pkg-config --cflags valhalla protobuf`

      
   

    echo "Linking final shared library..."
    # Get Abseil libraries dynamically
    # ABSL_LIBS=$(pkg-config --libs absl_log absl_strings absl_hash absl_synchronization absl_time | sed 's/-lrt//g')
    # ABSL_LIBS=$(pkg-config --static --libs \
    #             absl_cord absl_cord_internal \
    #             absl_cordz_info absl_cordz_handle absl_cordz_functions \
    #             absl_log_internal_message absl_log_internal_check_op \
    #             absl_log_internal_conditions absl_log_internal_format \
    #             absl_container_common absl_hash_function_defaults \
    #             absl_hash absl_city \
    #             absl_raw_hash_set \
    #             absl_strings absl_strings_internal \
    #             absl_base absl_spinlock_wait \
    #             absl_int128 absl_throw_delegate | sed 's/-lrt//g')

    # echo "ABSL_LIBS: $ABSL_LIBS"
     

    $CXX -shared -fPIC -std=c++17  govalhalla_wrap.cxx  \
      -o $out/lib/libgovalhalla.so \
      -Wl,-soname,libgovalhalla.so \
      -lvalhalla \
      `pkg-config --libs valhalla ` \
      -lprotobuf-lite \
      `pkg-config --libs protobuf-lite ` \
      -llz4 \
      -lz \
      -lpthread \
      -Wl,--verbose

  '';

  installPhase = ''
    
    # cp libgovalhalla.so $out/lib/
    cp govalhalla_wrap.cxx $out/lib/
    cp ${valhalla}/lib/libvalhalla.a $out/lib/
    cp -r *.go $out/govalhalla/
  '';
}