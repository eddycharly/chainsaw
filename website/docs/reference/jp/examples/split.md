# split

## Signature

`split(string, string)`

## Description

Splits the first string when the second string is found and converts it into an array.

## Examples

```
split('average|-|min|-|max|-|mean|-|median', '|-|', `3`) == ['average', 'min', 'max', 'mean|-|median']
```
