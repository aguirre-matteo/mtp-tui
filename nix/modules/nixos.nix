{ lib, pkgs, config, ... }:
let
  inherit (lib)
  types mkIf mkEnableOption mkOption;

  cfg = config.programs.mtp-tui;
  yamlFormat = pkgs.formats.yaml { };
in
{
  options.programs.mtp-tui = {
    enable = mkEnableOption "mtp-tui";
    package = mkOption {
      type = with types; nullOr package;
      default = pkgs.callPackage ../package.nix { };
      defaultText = "pkgs.mtp-tui";
      description = "The mtp-tui package to use.";
    };

    settings = mkOption {
      type = yamlFormat.type;
      default = { };
      example = {
        mount = {
          point = "/home/youruser/Documents/mtp";
          options = "default_permissions";
        };
      };
      description = ''
        Settings for mtp-tui. All available options can be found in the documentation at
        <https://github.com/aguirre-matteo/mtp-tui?tab=readme-ov-file#configuration>.
      '';
    };
  };

  config = mkIf cfg.enable {
    environment.systemPackages = mkIf (cfg.package != null) [ cfg.package ];
    environment.etc = mkIf (cfg.settings != { }) {
      "mtp-tui.yml".source = (yamlFormat.generate "mtp-tui.yml" cfg.settings);
    };
  };
}
