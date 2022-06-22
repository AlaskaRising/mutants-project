# Mutant validation :zap:

This project corresponds to the backend of a mutant validator, developed in [Go](https://go.dev/), and using [MongoDB](https://www.mongodb.com/) as a database engine.

## Installation / local-environment (Docker)

Requirements
- go 1.18 
- Beego 2.0.2
- MongoDB 5.0.7
- Docker

```shell
git clone https://github.com/AlaskaRising/mutants-project.git
cd mutants-project
docker-compose up
```
>After executing the commands, the app will be receiving requests on port 8085.

## Considerations
* A DNA sequence is considered to be mutant if the NxN matrix has more than one sequence of four identical letters, obliquely, horizontally or vertically.
* The letters of the Strings can only be: (A,T,C,G).

## Features

* POST → /mutant/  In case of verifying a mutant, it should return an HTTP 200-OK, otherwise a 403-Forbidden
>localhost:8085/mutant
* GET  → /stats/   Return a Json with the DNA verification statistics, e.g.: {“count_mutant_dna”:40, “count_human_dna”:100, “ratio”:0.4}
>localhost:8085/stats

## Architecture diagram
![](https://github.com/AlaskaRising/mutants-project/blob/master/diagrams/Blank%20diagram.png)

## Sequence diagram
![](https://github.com/AlaskaRising/mutants-project/blob/master/diagrams/Sequence%20diagram.png)

:tiger2: Thanks !!
