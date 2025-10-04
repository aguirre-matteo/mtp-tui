{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    systems.url = "github:nix-systems/default";
  };

  outputs = { self, nixpkgs, systems, ... }: let
    eachSystem = nixpkgs.lib.genAttrs (import systems);
  in {
    packages = eachSystem (system: let
      callPackage = nixpkgs.legacyPackages.${system}.callPackage;
    in {
      mtp-tui-unwrapped = callPackage ./nix/package-unwrapper.nix { };
      mtp-tui = callPackage ./nix/package.nix { };
      default = self.packages.${system}.mtp-tui;
    });

    overlay = self.overlays.default;
    overlays = {
      mtp-tui = import ./nix/overlay.nix;
      default = self.overlays.mtp-tui;
    };
    
    nixosModule = self.nixosModules.default;
    nixosModules = {
      mtp-tui = import ./nix/modules/nixos.nix;
      default = self.nixosModules.mtp-tui;
    };

    homeManagerModule = self.homeManagerModules.default;
    homeManagerModules = {
      mtp-tui = import ./nix/modules/home-manager.nix;
      default = self.homeManagerModules.mtp-tui;
    };
  };
}
