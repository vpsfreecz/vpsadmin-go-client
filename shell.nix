let
  pkgs = import <nixpkgs> {};
  stdenv = pkgs.stdenv;

in stdenv.mkDerivation rec {
  name = "vpsadmin-go-client";

  buildInputs = with pkgs;[
    git
    go
    gotools
  ];
}
