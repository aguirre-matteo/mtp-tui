{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    naersk.url = "github:nix-community/naersk";
    systems.url = "github:nix-systems/default-linux";
  };

  outputs =
    {
      self,
      nixpkgs,
      naersk,
      systems,
      ...
    }:
    let
      eachSystem = nixpkgs.lib.genAttrs (import systems);
    in
    {
      packages = eachSystem (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
          naersk' = pkgs.callPackage naersk { };
        in
        {
          mtp-tui = import ./nix/package.nix { inherit naersk' pkgs; };
          default = self.packages.${system}.mtp-tui;
        }
      );

      nixosModule = self.nixosModules.default;
      nixosModules = {
        mtp-tui = import ./nix/modules/nixos.nix naersk;
        default = self.nixosModules.mtp-tui;
      };

      homeManagerModule = self.homeManagerModules.default;
      homeManagerModules = {
        mtp-tui = import ./nix/modules/home-manager.nix naersk;
        default = self.homeManagerModules.mtp-tui;
      };
    };
}
