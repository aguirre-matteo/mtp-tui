{
  lib,
  buildGoModule,
}:

buildGoModule {
  pname = "mtp-tui-unwrapped";
  version = "unstable";
  src = ../.;
  vendorHash = "sha256-wk/UMSAL6K9UssnmEtFdpQ/tT0qKws11pldgnL9/dYA=";
  
  meta = {
    description = "A TUI for easily mounting and umounting your MTP devices!";
    homepage = "https://github.com/aguirre-matteo/mtp-tui";
    license = lib.licenses.mit;
  };
}
