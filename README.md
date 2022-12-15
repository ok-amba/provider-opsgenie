# Provider OpsGenie

`provider-opsgenie` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
OpsGenie API.

This provider is using the Opsgenie Terraform provider version `0.6.18` and has support for the following resources.

 - Alert Policy
 - API Integration
 - Custom Role
 - Email Integration
 - Escalation
 - Heartbeat
 - Integration Action
 - Team
 - Team Routing Rule
 - User

### References
- Opsgenie Terraform provider <https://registry.terraform.io/providers/opsgenie/opsgenie/latest/docs>
- Opsgenie API Documentation <https://docs.opsgenie.com/docs>

## Installation

This provider is created using Upjet. The idea behind Upjet is to use it with Upbound Market, where you upload the package with their CLI tool.
At the time of writing, the Upbound Market is brand new and still a bit wonky. Their CLI isn't much different and there's limited documentation.
That is why I made this provider like the original Terrajet code generation tool. When running `make build` it will create two containers which you will have
to push to a container registry and reference in the `Provider` and `ControllerConfig`.

Prerequisites:

 - Go 1.19
 - Docker
 - Make

### Build the images
Clone this repository and checkout the desired version tag.

Install the submodules.
```
make submodules
```

Export some varibles to set the registry URL and image tag.
```
export BUILD_REGISTRY="some/registry-url"
export VERSION=v0.0.1
```

Start the build.
```
make build
```

Above build command will create two container images, push them to your registry.
```
REPOSITORY                                                                                 TAG       IMAGE ID       SIZE
some/registry-url/provider-opsgenie-amd64           latest    092d173576b7   153MB
some/registry-url/provider-opsgenie-amd64           v0.0.1    092d173576b7   153MB
some/registry-url/provider-opsgenie-package-amd64   latest    f87030437c60   264kB
some/registry-url/provider-opsgenie-package-amd64   v0.0.1    f87030437c60   264kB
```

### Deploy to Kubernetes

Go to your Opsgenie API Key Management page and create an API key.

Create below secret with the Opsgenie API Key.

```
cat << EOF | kubectl apply -f - 
apiVersion: v1
kind: Secret
metadata:
  name: opsgenie-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "<your-api-key-here>",
      "api_url": "api.eu.opsgenie.com"
    }
EOF
```

Apply the needed CRDs from this repository. The CRD's must match the version of the provider you are about to install. That means that if you update to a newer version, remeber to apply the CRD's again.
```
kubectl apply -f package/crds/
```

Apply below `Provider` and `ControllerConfig`. Replace the registry URL in both objects.
```
cat << EOF | kubectl apply -f - 
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-opsgenie
spec:
  package: some/registry-url/provider-opsgenie-package-amd64:v0.0.1
  controllerConfigRef:
    name: opsgenie-config
---
apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: opsgenie-config
spec:
  image: some/registry-url/provider-opsgenie-amd64:v0.0.1
EOF
```

You can now start deploying Opsgenie resources.
You can find examples in the `examples` folder.

## Start Developing

Prerequisites:
 - Go 1.19
 - Docker
 - Kubectl
 - Make
 - Kind

Install Kind (Kubernetes in Docker), <https://kind.sigs.k8s.io/>.

After installation, create a Kind cluster.

```
kind create cluster --name opsgenie-provider
```

Export the Kind `kubeconfig` file location to use with `kubectl`.
```
export KUBECONFIG=./kubeconfig
```

Create the `crossplane-system` namespace.
```
kubectl create namespace crossplane-system
```

Go to your Opsgenie API Key Management page and create an API key.

Create below secret with the Opsgenie API Key.

```
cat << EOF | kubectl apply -f - 
apiVersion: v1
kind: Secret
metadata:
  name: opsgenie-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "<your-api-key-here>",
      "api_url": "api.eu.opsgenie.com"
    }
EOF
```

Apply the needed CRDs from this repository.
```
kubectl apply -f package/crds/
```

Apply the Provider Config.
```
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

Install the needed submodules.
```
make submodules
```

Generate the code.
```
make generate
```

Run the provider.
```
make run
```

You can now add Opsgenie resources to your Kind cluster and the provider should now create the in Opsgenie.

To test your newly added code run the `reviewable` command.
```
make reviewable
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/ok-amba/provider-opsgenie/issues).
