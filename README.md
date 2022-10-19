# Badges

[![CI](https://github.com/xh-polaris/cat-community-svc/actions/workflows/static-analysis.yml/badge.svg)](https://github.com/xh-polaris/cat-community-svc/actions/workflows/static-analysis.yml)
[![Build](https://github.com/xh-polaris/cat-community-svc/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/xh-polaris/cat-community-svc/actions/workflows/docker-publish.yml)

## Get started

**Start services**

```bash
go run rpc/cat_community.go -f rpc/etc/cat_community.yaml
```

```bash
go run api/cat_community.go -f api/etc/cat_community.yaml
```

Before starting the server, please replace the default config file in `etc` directory.
