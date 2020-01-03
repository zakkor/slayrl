with import <nixpkgs> {};

buildGoPackage rec {
  name = "slayrl";
  src = ./.;
  buildInputs = [
    libGL
    xorg.libX11
    xorg.libXrandr
    xorg.libXcursor
    xorg.libXinerama
    xorg.libXi
    xorg.libXxf86vm
  ];
  goDeps = ./deps.nix;
  goPackagePath = "github.com/zakkor/slayrl";
}