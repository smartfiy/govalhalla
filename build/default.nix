let


  pkgs = import (builtins.fetchGit {
    url = "https://github.com/NixOS/nixpkgs";
    ref = "nixos-24.05";
    rev = "087f43a1fa052b17efd441001c4581813c34bc19";       #"666e1b3f09c267afd66addebe80fb05a5ef2b554"; "nixos-24.11";
  }) { };
in
pkgs.callPackage ./bindings { } #{ inherit pkgs; }


# "nixos-unstable";
#     rev = "9abb87b552b7f55ac8916b6fc9e5cb486656a2f3"; 