# Helm SSM Parameter Store and Secrets Manager Plugin

This is a **helm3** plugin to help platform engineers to inject values coming from AWS SSM
parameters and AWS Secrets Manager , on the `values.yaml` file.

## Usage

Loads a template file,  and writes the output.

Simply add placeholders like `{{ssm "path" "option1=value1" }}` in your
file, where you want it to be replaced by the plugin.

Currently the plugin supports the following options:
- `region=eu-west-1` - to resolve that parameter in a specific region (if you don't specify it will assume us-east-1 which is default region)

### Values file

```yaml
service:
ingress:
  enabled: false
  hosts:
    - service.{{ssm "/exists/subdomain" }}
    - service1.{{sm "/acmsecret" }}
    - service2.{{ssm "/exists/subdomain" "region=eu-west-1" }}
    - service3.{{sm "dbendpoint" "region=eu-west-1" }}

```

### Command

```sh
$ make install 
$ helm ssm-sm [flags]
```

### Flags

```sh
  -d, --dry-run                 doesn't replace the file content
  -h, --help                    help for ssm
  -f, --values valueFilesList   specify values in a YAML file (can specify multiple) (default [])
  -v, --verbose                 show the computed YAML values file/s
```

## Install

Choose the latest version from the releases and install the
appropriate version for your OS as indicated below.

```sh
$ helm plugin add https://github.com/ugurcancaykara/helm-ssm-sm
```

### Developer (From Source) Install

If you would like to handle the build yourself, instead of fetching a binary,
this is how we recommend doing it.

- Make sure you have [Go](http://golang.org) installed.

- Clone this project

- In the project directory run
```sh
$ make install
```

## License
helm-ssm is available under the MIT license. See the LICENSE file for more info.


### Installed packages
- go get -u github.com/spf13/cobra
- go get github.com/aws/aws-sdk-go-v2/aws
- go get github.com/aws/aws-sdk-go-v2/config
- go get github.com/aws/aws-sdk-go-v2/service/ssm
- go get -u github.com/aws/aws-sdk-go-v2/service/secretsmanager
- go get -u gotest.tools/v3/assert
- go get github.com/stretchr/testify/assert

