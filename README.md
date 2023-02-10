## Remote Debugging

#### Prepreq
- [vscode](https://code.visualstudio.com/)
- [docker](https://www.docker.com/)
- [local kind k8s cluster](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [golang](https://go.dev/doc/install)

#### Build & Deploy
- Run `.\build.ps1; .\deploy.ps1`
  - This will build the docker image, load the image into kind cluster and deploy. If you have the image built and loaded into kind cluster and all you want is to change the kubernetes spec of your deployment, you can instead run `kubectl apply -f otel-custom.yaml`

#### Attach the Debugger 
- Run `kubectl port-forward svc/otel-custom-dlv 40000:40000`
- Attach the debugger by launching "Connect to server" from "launch.json" configs.

#### Important
In my case docker image build and deploy repeatedly took a lot of space, keep an eye on your disk space and periodically run `docker image prune`. Keep an eye on `$env:USERPROFILE\AppData\Local\Docker\wsl\data\ext4.vhdx` file, if it is too bloated, you can reclaim space by running `Optimize-VHD -Path $env:USERPROFILE\AppData\Local\Docker\wsl\data\ext4.vhdx -Mode full`. In order to run this, you need to stop docker and run `wsl --shutdown`.