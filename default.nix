with import <nixpkgs> {};
stdenv.mkDerivation {
    name = "dev-environment"; # Probably put a more meaningful name here
    buildInputs = [ go x11 xorg.libXrandr libGL xorg.libXcursor xorg.libXinerama xorg.libXxf86vm xorg.libXi mpv glfw pkg-config ];
}
