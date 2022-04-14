{
  description = "Implement a custom state transition function for the Deku sidechain with Go";
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs =
    { self
    , nixpkgs
    , flake-utils
    }:
      with flake-utils.lib;
      eachSystem defaultSystems (system:
      let pkgs = import nixpkgs { inherit system; }; in
      rec {
        devShell = pkgs.mkShell {
          nativeBuildInputs = with pkgs;
            [
              # Go developer tooling
              go
              gopls
              gore
            ];
        };
        packages = rec {
          binary = pkgs.buildGoModule {
            pname = "state-transition";
            version = "0.0.1";
            src = ./.;
            vendorSha256 = "sha256-xf0g5duhIw9wFnBmaUTmQJCNo0vtvtawcMl02kZwCEY=";
          };
          dockerImage = pkgs.dockerTools.buildImage {
            name = binary.pname;
            config = {
              Entrypoint = [ "${binary}/bin/state_transition" "/app/fifo"];
            };
          };
        };
        defaultPackage = packages.dockerImage;
      });
}
