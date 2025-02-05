let

  pkgunstable = import (builtins.fetchGit {
    url = "https://github.com/NixOS/nixpkgs";
    ref = "nixos-unstable";
    rev = "9abb87b552b7f55ac8916b6fc9e5cb486656a2f3";
  }) { };

  pkgs = import (builtins.fetchGit {
    url = "https://github.com/NixOS/nixpkgs";
    ref = "nixos-24.05";
    rev = "087f43a1fa052b17efd441001c4581813c34bc19";
  }) { 

    overlays  = [ # Make this a list
      (self: super: {
        swig = pkgunstable.swig;  # Use swig from unstable

        protobuf = pkgunstable.protobuf.overrideAttrs (oldAttrs: { # Use super.protobuf
          cmakeFlags = (oldAttrs.cmakeFlags or []) ++ [
            "-Dprotobuf_BUILD_SHARED_LIBS=OFF"
            "-Dprotobuf_BUILD_TESTS=OFF"
            "-Dprotobuf_BUILD_EXAMPLES=OFF"
          ];
        });
      }) # Close the function definition
    ]; # Close the list

  };

in

  pkgs.callPackage ./bindings/default.nix {}