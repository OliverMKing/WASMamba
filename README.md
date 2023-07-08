# WASMamba

A [Battlesnake](https://play.battlesnake.com/) written with Spin WASM and Go.

## Local

Install [Spin](https://developer.fermyon.com/spin/quickstart#install-spin), [TinyGo](https://tinygo.org/getting-started/install/), and the [Battlesnake CLI](https://github.com/BattlesnakeOfficial/rules/blob/main/cli/README.md).

Run the snake with `spin build` then `spin up`. Run the server with `battlesnake play --name "WASMamba" --url http://localhost:3000 --name "WASMamba2" --url http://localhost:3000 --browser`.

## Infrastructure

Infrastructure can be provisioned by running `terraform init && terraform apply` in the `tf` directory.

Ocassionally, we may need to update manifests. Install the [spin k8s plugin](https://github.com/chrismatteson/spin-plugin-k8s). Run `spin k8s scaffold wasmamba -o` to get the latest Dockerfile and manifests. Compare the manifests to the current [manifests](./manifests/). Additionally, see [spin kubernetes docs](https://developer.fermyon.com/spin/kubernetes) for latest best practices.