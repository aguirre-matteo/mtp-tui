{ naersk', pkgs }:

naersk'.buildPackage {
  name = "mtp-tui";
  version = "unstable";
  src = ../.;

  nativeBuildInputs = [ pkgs.makeWrapper ];
  buildInputs = with pkgs; [
    jmtpfs
    fuse
  ];

  postInstall = with pkgs; ''
    wrapProgram $out/bin/mtp-tui \
      --prefix PATH : "${
        lib.makeBinPath [
          jmtpfs
          fuse
        ]
      }"
  '';

  meta = {
    description = "A TUI for easily mounting and umounting your MTP devices!";
    homepage = "https://github.com/aguirre-matteo/mtp-tui";
    license = pkgs.lib.licenses.mit;
  };
}
