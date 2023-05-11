[deSEC](https://desec.io) module for Caddy
==========================================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It
can be used to manage DNS records with [deSEC](https://desec.io).

## Caddy module name

```
dns.providers.desec
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "desec",
				"token": "YOUR_TOKEN"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns desec {
		token "YOUR_TOKEN"
	}
}
```

```
# one site
tls {
	dns provider_name desec {
		token "YOUR_TOKEN"
	}
}
```