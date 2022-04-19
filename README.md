# Mutant validation :zap:

This project corresponds to the backend of a mutant validator, developed in [Go](https://go.dev/), and using [MongoDB](https://www.mongodb.com/) as a database engine.

## Installation / local-environment

Requirements
- go 1.18 
- Beego 2.0.2
- MongoDB 5.0.7

```shell
git clone https://github.com/AlaskaRising/mutants-project.git
cd mutants-project
bee run
```
>After executing the commands, the app will be receiving requests on port 8080, it is essential to have a MongoDB engine running on the default port 27017.

## Installation / docker

## Considerations
* A DNA sequence is considered to be mutant if the NxN matrix has more than one sequence of four identical letters, obliquely, horizontally or vertically.
* The letters of the Strings can only be: (A,T,C,G).

## Features

* POST → /mutant/  In case of verifying a mutant, it should return an HTTP 200-OK, otherwise a 403-Forbidden
* GET  → /stats/   Return a Json with the DNA verification statistics, e.g.: {“count_mutant_dna”:40, “count_human_dna”:100: “ratio”:0.4}

:tiger2: Thanks !!
