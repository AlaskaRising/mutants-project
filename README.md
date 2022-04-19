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

#### Another way to run the project is by importing it into an editor like IntelliJ or SpringToolSuite.

## Considerations
* A DNA sequence is considered to be mutant if the NxN matrix has more than one sequence of four identical letters, obliquely, horizontally or vertically.
* The letters of the Strings can only be: (A,T,C,G).



## Features


Endpoints are available to play from the creation of the roulette, opening bets, placing bets, closing bets and a query of the results obtained by all bettors.

#### The available endpoints correspond to: 


:tiger2: Thanks !!
