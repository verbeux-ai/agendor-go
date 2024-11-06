# Agendor Go

Essa biblioteca foi desenvolvida para o uso do [agendor](https://www.agendor.com.br/), crm via api.

Status: Em desenvolvimento

## Criando o Client

Este é um exemplo de criação de people

```go
import myzap "github.com/verbeux-ai/agendor-go"
import "context"

client = agendor.NewClient(
    myzap.WithToken(os.Getenv("TOKEN")),
)
response, err := client.CreatePeopleDeal(ctx, peopleID, agendor.CreatePeopleDealRequest{
    Title: "People",
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