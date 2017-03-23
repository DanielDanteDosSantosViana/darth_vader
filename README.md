# GoSAD - API Rest do Sistema de Apontamento

# Instalando todas as dependências
```
   go get ./...
```
# Instalação
```
   na raiz do projeto:

   go install
```

# Modo de usar

Para utilizar o encurtador basta rodar em linha de comando:
```
   go run main.go

```

Existe um arquivo de configuração default utilizado no projeto chamado
conf.toml, porém pode se passar em linha de comando o caminho para outro
arquivo de configuração:
```
   go run main.go -config [path]

```

# Arquivo de configuração

O arquivo de configuração é essencial para o projeto, pois nele estão as
referências para os acessos na base e porta do serviço.

Arquivo config default:
```
[service]
port = ":8081"

[db]
mysqlread =  "root:1234@/sda?charset=utf8"
mysqlwrite = "root:1234@/sda?charset=utf8"

```


#Banco
```
MySQL

Obs:
  - o schema do banco está no arquivo schema.sql
  - Existem dois acessos de banco, mysqlread(onde colocaria o usuário com apenas permissão de leitura) e mysqlwrite(usuário com permissão de escrita). A ideia inicial era deixar o banco de leitura com Redis, para cachear as consultas.

```



# Teste

Para rodar todos os teste é necessário o seguinte  comando :
```
    na raiz do projeto:

    go test -v ./...

```
No projeto está sendo utilizado o goconvey(https://github.com/smartystreets/goconvey.git), então
você pode acessar  :
```
 $ $GOPATH/bin/goconvey
```
Os resultados dos testes serão mostrados na seguinte url:
```
http://localhost:8080

```
