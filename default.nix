with import <nixpkgs> {};

buildGoPackage rec {
  name = "slayrl";
  src = ./.;
  buildInputs = [
    pkg-config
    libGL
    xorg.libX11
    xorg.libXrandr
    xorg.libXcursor
    xorg.libXinerama
    xorg.libXi
    xorg.libXxf86vm
    xorg.libXext
  ];
  goDeps = ./deps.nix;
  goPackagePath = "github.com/zakkor/slayrl";
}