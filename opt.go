package structgraph

func WithScope(path ...string) Option {
	return func(d *drawer) {
		d.scopes = append(d.scopes, path...)
	}
}
