let
    pkgs = import (builtins.fetchGit {

        #name = "nixpkgs";
        url = "https://github.com/NixOS/nixpkgs";
        ref = "nixos-24.11";
        rev = "e24b4c09e963677b1beea49d411cd315a024ad3a"; 
    }){

    };


in
pkgs.callPackage ./bindings #{ inherit pkgs; }
{
  inherit (pkgs)
    lib stdenv  cmake fetchFromGitHub zlib
    boost179 pkg-config sqlite geos lz4 swig protobuf go;
    # swig = nixpkgs.swig; # Set the SWIG version to 4.0.2    
    # protobuf = nixpkgs.protobuf;
  
}