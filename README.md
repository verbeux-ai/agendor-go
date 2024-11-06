# Agendor Go

Essa biblioteca foi desenvolvida para o uso do [agendor](https://www.agendor.com.br/), crm via api.

Status: Em desenvolvimento

## Instalando

```
go get github.com/verbeux-ai/agendor-go
```

## Criando o Client

Este é um exemplo de criação de people

```go
import agendor "github.com/verbeux-ai/agendor-go"
import "context"

client = agendor.NewClient(
    agendor.WithToken(os.Getenv("TOKEN")),
)
people, err := client.CreatePeople(ctx, agendor.CreatePeopleRequest{
    Name: "Teste agendor-go",
})

if err != nil {
    panic(err)
}

fmt.Println(response)
```

## Criando People

Este é um exemplo de como criar uma people


## Features disponíveis

| Funcionalidade          | Implementado |
|-------------------------|--------------|
| Criar People            | ✔            |
| Listar People           | ✔            |
| Criar People Deal       | ✔            |
| Atualizar Deal          | ✔            |

### Futuro
Temos planos de manter o repositório, mas isso pode mudar. 


> Você está convidado a contribuir ao repositório!