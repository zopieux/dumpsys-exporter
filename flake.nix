{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, flake-utils, nixpkgs }: flake-utils.lib.eachDefaultSystem (system:
    let pkgs = import nixpkgs { inherit system; }; in
    {
      defaultPackage = pkgs.buildGoModule rec {
        pname = "dumpsys-exporter";
        version = "local";
        src = ./.;
        vendorHash = pkgs.lib.fakeHash;
      };
      devShell = pkgs.mkShell {
        buildInputs = with pkgs; [ go gopls go-tools protoc-gen-go protobuf ];
      };
    });
}
