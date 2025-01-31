# HTTP Exporter
Simple HTTP exporter for Prometheus
```html
exported metrics:
    http_status - gauge


#### how to build

```shell
# example build for Apple macOS & Apple Silicon Chip
./scripts/build.sh darwin arm64
```

```shell
# example build for apple macOS & Intel Chip
./scripts/build.sh darwin amd64
```
```shell
# build with no args for help
./scripts/build.sh 
```

####  systemd service example

```unit file (systemd)
[Unit]
  Description=HTTP exporter
  Wants=network-online.target
  After=network-online.target

[Service]
  ExecStart=/usr/local/bin/http-exporter --config /etc/http-exporter/config.yaml
  SyslogIdentifier=http-exporter
  Restart=always

[Install]
  WantedBy=multi-user.target

```
####  ./config.yml example
```yaml
port: 8080
interface: 0.0.0.0
interval: 30s
urls:
  - "https://example.com"
  - "https://secureexample.com"
```
