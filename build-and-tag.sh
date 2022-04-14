nix build
image_name=$(docker load < result | cut -d ' ' -f 3)
docker image tag $image_name state-transition:latest
