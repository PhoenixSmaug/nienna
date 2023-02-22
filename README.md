# Nienna

This is an implementation of Knuth's Algorithm X in Go. It finds all solutions of the Exact Cover Problem, where you have given a (0,1)-Matrix and need to find all subsets of rows which add to a vector with 1 in all entries.

```math
\begin{pmatrix}0 & 0 & 1 & 0 & 1 & 1 & 0\\1 & 0 & 0 & 1 & 0 & 0 & 1\\0 & 1 & 1 & 0 & 0 & 1 & 0\\ 1 & 0 & 0 & 1 & 0 & 0 & 0\\0 & 1 & 0 & 0 & 0 & 0 & 1\\0 & 0 & 0 & 1 & 1 & 0 & 1 \end{pmatrix}
```

For this exemplary matrix only one solution exists, consisting of the first, fourth and fifth row:

```math
\begin{matrix} & 0 & 0 & 1 & 0 & 1 & 1 & 0\\ + & 1 & 0 & 0 & 1 & 0 & 0 & 0\\+ & 0 & 1 & 0 & 0 & 0 & 0 & 1\\\hline  & 1 & 1 & 1 & 1 & 1 & 1 & 1 \end{matrix}
```

This NP-complete problem can be extended with secondary columns, in whose entry the sum of the rows is no longer required to be exactly 1, but can also be 0. A lot of famous combinatorial problems can be nicely modelled as an Exact Cover Problem, such as:

* Sudoku

* n-Queens

* Polyomino Packing

Algorithm X and these three examples are implemented, a detailed description of the solution algorithm can be found in [Knuth's paper](https://arxiv.org/abs/cs/0011047).

(c) Mia Muessig