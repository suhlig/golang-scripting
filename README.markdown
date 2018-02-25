# Scripting with Go

Inspired by a [blog post](https://blog.cloudflare.com/using-go-as-a-scripting-language-in-linux/) I experimented a bit with this approach. This [gist](https://gist.github.com/posener/73ffd326d88483df6b1cb66e8ed1e0bd) goes into more details.

# Results

In general the approach seems to work fine, with the **major caveat** that a non-zero exit code is always reported as `1` instead of the real exit code.

The alternative of defining a new binary format with the `binfmt_misc` module is way more elaborate, but Ubuntu Trusty does not support the `binfmt_misc` module out of the box (Xenial has it).

# Test

You'll need Vagrant and Ruby installed; then:

```bash
vagrant up
./test
```
