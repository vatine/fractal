# fractals

Various fractal generators. This was both a "get more comfortable with
spf13/cobra" as well as "hmm, fractals are kind of neat" project.

## Building

As can be seen from the go.mod, this requires (until merged) a patched
spf13/pflag, specifically the fork off github.com/vatine/pflag. Clone
that repo to somewhere locally, then change the rewrite in go.mod to
point somewhere sensible.

This also requires go 1.15 (or higher) to build.
