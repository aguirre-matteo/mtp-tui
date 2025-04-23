{
  outputs = { self, ... }: {
    package = import ./nix/package.nix;
    overlay = import ./nix/overlay.nix;
    nixosModule = import ./nix/modules/nixos.nix;
    homeManagerModule = import ./nix/modules/home-manager.nix;
  };
}
