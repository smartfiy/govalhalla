let
  pkgs = import (builtins.fetchGit {
    url = "https://github.com/NixOS/nixpkgs";
    ref = "nixos-unstable";
    rev = "9abb87b552b7f55ac8916b6fc9e5cb486656a2f3";
  }) { };
in
pkgs.callPackage ./bindings { }
