# Agilizando arquitetura de modelos com Swagger e Protobuf

Exemplo de projeto de API usando Swagger, Protobuf & Bindata para facilitar a sua vida

## Instalação

### Docker

É necessário ter o docker para poder trabalhar com esse projeto

```sh
cd scripts/{SEU SO}
chmod +x ensure_docker.sh
./ensure_docker.sh
```

### Geradores de Classes

Para gerar modelos a partir de Swagger ou ProtoBuf, é necessário instalar as ferramentas e rodar os scripts, dentro da pasta scripts

### Swagger

```sh
cd scripts/{SEU SO}
chmod +x install_goswagger.sh
./install_goswagger.sh
```

### ProtoBuf

```sh
cd scripts/{SEU SO}
chmod +x install_proto.sh
./install_proto.sh
```

Obs: É importante lembrar que se os modelos forem deletados, é preciso remover manualmente pois o script não os remove.


### Gerador do Schema do MongoDB

Para embutir os schemas da documentação api/data/schema no código para que o sistema aplique a estrutura ali definida no MongoDB, é necessário instalar o go-bindata

#### BinData

```sh
cd scripts/{SEU SO}
chmod +x install_gobindata.sh
./install_gobindata.sh
```

## Desenvolvimento

### Requisitos

É necessário ter o [Go Dep](https://github.com/golang/dep) para gerenciar as dependências.
O jeito mais fácil para instalar no Ubuntu é:

#### MacOS
```sh
 brew install dep
```


#### Ubuntu
```sh
 sudo apt-get install go-dep
```
