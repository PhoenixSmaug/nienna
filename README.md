# Nienna

This repository is still in early development and is subject to heavy change.

## Algorithm X

Donald Knuth describes his algorithm in detail [here](https://arxiv.org/abs/cs/0011047). It finds all solutions of the Exact Cover Problem, where you have given a (0,1)-Matrix and need to find all subsets of rows which add to a vector with 1 in all entries.

```math
\begin{pmatrix}0 & 0 & 1 & 0 & 1 & 1 & 0\\1 & 0 & 0 & 1 & 0 & 0 & 1\\0 & 1 & 1 & 0 & 0 & 1 & 0\\ 1 & 0 & 0 & 1 & 0 & 0 & 0\\0 & 1 & 0 & 0 & 0 & 0 & 1\\0 & 0 & 0 & 1 & 1 & 0 & 1 \end{pmatrix}
```

For this exemplary matrix only one solution exists, consisting of the first, fourth and fifth row. This NP-complete problem can be extended with secondary columns, in whose entry the sum of the rows is no longer required to be exactly 1, but can also be 0.

The following three combinatorial problems are implemented and have an example in main.go:

* **n-Queens** with one or all solutions

* **Sudoku** solver

* Pack **Pentominoes** into a rectangle

## Algorithm C

## Algorithm M