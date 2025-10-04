{
  lib,
  stdenvNoCC,
  callPackage,
  jmtpfs,
  fuse,
  makeWrapper
}:

stdenvNoCC.mkDerivation {
  pname = "mtp-tui";
  version = "unstable";
  src = callPackage ./package-unwrapper.nix {};
  
  nativeBuildInputs = [ makeWrapper ];
  buildInputs = [
    jmtpfs
    fuse
  ];

  installPhase = ''
    mkdir -p $out/bin
    makeWrapper $src/bin/mtp-tui $out/bin/mtp-tui \
      --prefix PATH : "${lib.makeBinPath [ jmtpfs fuse ]}"
  '';

  meta = {
    description = "A TUI for easily mounting and umounting your MTP devices!";
    homepage = "https://github.com/aguirre-matteo/mtp-tui";
    license = lib.licenses.mit;
  };
}
