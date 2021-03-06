# kubectl-istiolog [![build](https://github.com/TejaBeta/kubectl-istiolog/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/TejaBeta/kubectl-istiolog/actions/workflows/build.yml) [![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](./LICENSE)

A plugin to manipulate `istio-proxy` logging level and also 
provides tailing of logs.

## Installation

Download precompiled binaries from [GitHub Releases](https://github.com/TejaBeta/kubectl-istiolog/releases).

### Via [Krew](https://github.com/kubernetes-sigs/krew)

```bash
kubectl krew install istiolog
```

## Usage

```bash
kubectl istiolog <<podname>> -n <<namespace>> -l debug -f
```

`istio-proxy`(envoy instance) comes with a default logging level as `Warning`, 
the above command updates the logging level of Envoy instance to `debug` and 
also helps in tailing the logs of `istio-proxy` container which is a sidecar
in one go. 

`kubectl istiolog` supports all the logger names and logger levels similar
to `istio proxy-config`.

On exit, logging level of Envoy instance will be reverted back to default 
logging level `Warning`.

## Help Menu

```bash
A Kubectl plugin to manage and set envoy log levels

Usage:
  kubectl-istiolog [pod] [flags]
  kubectl-istiolog [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  version     print current kubectl-istiolog version

Flags:
  -f, --follow             Specify if the logs should be streamed
  -h, --help               help for kubectl-istiolog
  -l, --level string       Comma-separated minimum per-logger level of messages to output (default "warning")
  -n, --namespace string   Namespace in current context (default "default")
      --verbose            Verbose mode on

Use "kubectl-istiolog [command] --help" for more information about a command.
```

## Supported Logger Names

```
admin, aws, assert, backtrace, client, config, connection, conn_handler, dubbo, file, filter, forward_proxy, grpc, hc, health_checker, http, http2, hystrix, init, io, jwt, kafka, lua, main, misc, mongo, quic, pool, rbac, redis, router, runtime, stats, secret, tap, testing, thrift, tracing, upstream, udp, wasm
```

## Supported Logger Levels

```
trace, debug, info, warning, error, critical, off
```

## License
This project is distributed under the 
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see
[LICENSE](./LICENSE) for more information.
