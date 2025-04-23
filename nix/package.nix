{ pkgs }:

pkgs.buildGoModule {
  pname = "mtp-tui";
  version = "unstable";
  src = ../.;
  vendorHash = "sha256-wk/UMSAL6K9UssnmEtFdpQ/tT0qKws11pldgnL9/dYA=";
}
