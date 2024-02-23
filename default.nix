{ pkgs ? (let
  inherit (builtins) fetchTree fromJSON readFile;
  inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
in import (fetchTree nixpkgs.locks) {
  overlays = [ (import "${fetchTree gomod2nix.locked}/overlay.nix") ];
}), buildGoApplication ? pkgs.buildGoApplication }:
buildGoApplication {
  pname = "quickview";
  version = "main";
  pwd = ./.;
  src = ./.;
  subPackages = [ "cmd/quickview" ];
  modules = ./gomod2nix.toml;
  meta.mainProgram = "quickview";
}
