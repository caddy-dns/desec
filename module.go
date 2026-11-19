package desec

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/desec"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *desec.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.desec",
		New: func() caddy.Module { return &Provider{new(desec.Provider)} },
	}
}

// Provision sets up the module.
// Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.Token = caddy.NewReplacer().ReplaceAll(p.Provider.Token, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	desec {
//	     token <token>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	tokenSet := false
	for d.Next() {
		if d.NextArg() {
			p.Provider.Token = d.Val()
			tokenSet = true
			if d.NextArg() {
				return d.ArgErr()
			}
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "token":
				if tokenSet {
					return d.Err("token already set")
				}
				if !d.NextArg() {
					return d.ArgErr()
				}
				p.Provider.Token = d.Val()
				tokenSet = true
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if !tokenSet || p.Provider.Token == "" {
		return d.Err("missing token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
