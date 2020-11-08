# Usage

This module implements a [goldmark](https://github.com/yuin/goldmark) renderer
that can output manpages.

The _intended usage_ is to **transform** bits of `markdown` based documentation
into proper manpages by combining them with a skeleton prepared beforehand.

## codeblock

```sh
DEBUG=
echo ${DEBUG}
```

## lists

1. hi
2. what's
3. up

- oh
- oh2
- oh3

* nothing
* nothing2
* nothing3

## other fun

> just checking out
> blockquotes

### checking some stuff

[example.com][ex]

[ex]: https://example.com
