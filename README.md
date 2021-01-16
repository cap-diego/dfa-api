### API para la [pag](https://dfa.diegobuceta.com) de minimizador AFD

#### Objetivo: escuchar por peticiones que tengan en el body un AFD, aplicarle el [algoritmo](https://github.com/cap-diego/dfa-minimization-algorithm) y retornar el AFD mínimo.


#### Para correrlo localmente:
  - tener los certificados para https: ca_bundle.crt y certificate.crt en el directorio principal
  - correr con ``` go run main.go ``` o bien crear el binario para correr en otro nodo: ``` GOOS=linux go build main.go  ``` (ejemplo para linux).
  - El body de la request tiene que tener la pinta:
      ```json
      {
          "states":[0,1],
          "alphabet":[0,1],
          "initialState": 1,
          "finalStates":[1],
          "delta":{
              "0":{"0":1,"1":0},
              "1":{"0":1,"1":0}}
      }
    ```
