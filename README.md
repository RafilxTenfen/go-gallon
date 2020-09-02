# Go-Gallon

### Problem 
Dado um conjunto de garrafas d'água, com volumes de água diferentes entre si, e um galão de água.
Crie um algoritmo, na linguagem que achar melhor, para escolher as garrafas a serem utilizadas para encher o galão, de acordo:
- 1) O galão deve ser completamente preenchido com o volume das garrafas;
- 2) Procure esvaziar totalmente as garrafas escolhidas;
- 3) Quando não for possível esvaziar todas garrafas escolhidas, deixe a menor sobra possível;
- 4) utilize o menor número de garrafas possível;
Exemplo 1: galão de 7L, entre 5 garrafas: [1L, 3L, 4.5L, 1.5L, 3.5L ]
   resposta: [1L, 4.5L, 1.5L], sobra 0L;
Exemplo 2: galão de 5L, entre 4 garrafas: [1L, 3L, 4.5L, 1.5L ]
   resposta: [1L, 4.5L]; sobra 0.5L;

- All the values of water used are in (ML)

### How to Run

#### Using Docker
- Build the docker image
```shell
$~ docker build -t go-gallon .
```

- Run the binary inside the image
```shell
$~  docker run go-gallon:latest ./gallon run -h
```

- Testing the example 1
```shell
$~ docker run go-gallon:latest ./gallon run -w 7000 -g 1000,3000,4500,1500,3500
```

- Testing the example 2
```shell
$~ docker run go-gallon:latest ./gallon run -w 5000 -g 1000,3000,4500,1500
```


#### Using golang

- Type the following command line to build the binary `make`

- This will generate the binary `gallon`
- To see the available commands, run:
```shell
$~ ./gallon run -h
```

- To set the value of water to test, use the parameter `-w`, like for 4L o water it would be `-w 4000`
- To set the value of gallons to test, use the parameter `-g`, like for 3 gallons [1L, 3L, 4.5L] it would be `-g 1000,3000,4500`

- Example of testing the example 1
```shell
$~ ./gallon run -w 7000 -g 1000,3000,4500,1500,3500
```
__Expected Result__
```shell
INFO[0000] The bottles to be used to fill the water are  Bottles(ML)="[{1000} {1500} {4500}]" Rest of Water(ML)=0 Total of Water(ML)=7000
```


- Example of testing the example 2
```shell
$~ ./gallon run -w 5000 -g 1000,3000,4500,1500
```

__Expected Result__
```shell
INFO[0000] The bottles to be used to fill the water are  Bottles(ML)="[{1000} {4500}]" Rest of Water(ML)=500 Total of Water(ML)=5000
```


### How to Test
- Type the following command line to run the tests `make check`