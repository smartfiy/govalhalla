let

  # pkgunstable = import (builtins.fetchGit {
  #   url = "https://github.com/NixOS/nixpkgs";
  #   ref = "nixos-unstable";
  #   rev = "9abb87b552b7f55ac8916b6fc9e5cb486656a2f3";
  # }) { 
   
  # #  overlays = [(self: super: {
  # #     Stdenv = super.stdenv.override {
  # #       name = "clangStdenv";
  # #       cc = super.clang;
  # #       cxx = super.clangxx;
  # #     };
  # #   })];

  # };

  pkgs = import (builtins.fetchGit {
    url = "https://github.com/NixOS/nixpkgs";
    ref = "nixos-unstable";
    rev = "9abb87b552b7f55ac8916b6fc9e5cb486656a2f3";
  }) { 

    # config.replaceStdenv = { pkgs, ... }: pkgs.gcc13Stdenv;
    overlays  = [ # Make this a list
      (self: super: {

        # swig = pkgunstable.swig;  # Use swig from unstable

        # stdenv = super.stdenv.override {
        #   cc = super.gcc12;
        #   # cxx = super.gcc12;
        # };


        lz4 = super.lz4.overrideAttrs (oldAttrs: {
          stdenv = super.gcc13Stdenv;
          cmakeFlags = (oldAttrs.cmakeFlags or []) ++ [
            "-DBUILD_SHARED_LIBS=OFF"
            "-DCMAKE_POSITION_INDEPENDENT_CODE=ON"
          ];
         
        });  # Use lz4 from unstable

        # zlib = super.zlib.overrideAttrs (oldAttrs: {
        #   static = true;
        #   shared = false;
        # });

        # abseil-cpp = pkgunstable.abseil-cpp.override {
        #   static = true;
          
        # };
        
        boost = super.boost.override {
          # stdenv = super.gcc13Stdenv;
          enableStatic = true;
          enableShared = false;
          python = null;
        };
        
        protobuf = super.protobuf.override {
          # stdenv = super.gcc13Stdenv;
          enableShared = false;
          abseil-cpp = super.abseil-cpp.override {
          stdenv = super.gcc13Stdenv;
          static = true;
          };
        };
        
      }) # Close the function definition
    ]; # Close the list

  };

in pkgs.callPackage ./bindings/default.nix {}
