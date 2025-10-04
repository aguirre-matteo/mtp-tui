final: prev: {
  mtp-tui-unwrapped = prev.callPackage ./package-unwrapper.nix { };
  mtp-tui = prev.callPackage ./package.nix { };
}
