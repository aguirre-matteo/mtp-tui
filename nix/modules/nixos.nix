{ lib, pkgs, config, ... }:
let
  inherit (lib)
  mkIf mkEnableOption mkPackageOption mkOption;

  cfg = config.programs.mtp-tui;

  formatter = pkgs.formats.yaml { };
in
{
  options.programs.mtp-tui = {
    enable = mkEnableOption "mtp-tui";
    package = mkPackageOption pkgs "mtp-tui";
    settings = mkOption {
      type = formatter.type;
      default = { };
      example = ''
        mount = {
          point = "/home/youruser/Documents/mtp";
          options = "default_permissions";
        };
      '';
      description = ''
        Settings for mtp-tui. All available options can be found in the documentation at
        <https://github.com/aguirre-matteo/mtp-tui?tab=readme-ov-file#configuration>.
      '';
    };
  };

  config = mkIf cfg.enable {
    environment.systemPackages = [
      cfg.package 
      pkgs.jmtpfs
      pkgs.fuse
    ];

    environment.etc = mkIf (cfg.settings != { }) {
      "mtp-tui.yml".source = (formatter.generate "mtp-tui.yml" cfg.settings);
    };
  };
}
