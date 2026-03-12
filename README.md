# Collatz Conjecture

A Go implementation of the [Collatz Conjecture](https://en.wikipedia.org/wiki/Collatz_conjecture) problem.

## TODO

- Guard against silent integer overflow in `calculateNewTerm`: `term * 3 + 1` wraps silently for large `uint` values
- Consider an iterative implementation as a safeguard against stack overflow on extreme seeds
