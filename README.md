# Whatsapp

## Dependencies

```txt

_ "github.com/joho/godotenv/autoload"
_ "github.com/lib/pq"
"github.com/google/uuid"
```
- https://whimsical.com/botmaker-9FSPjqbf1gpfTGAaS4WByy


## MODELOS DE MENSAGENS - USUÁRIO RESPONDEU ESCREVENDO
```json
[
  {
    "object":"whatsapp_business_account",
    "entry":[
      {
        "id":"110880075418235",
        "changes":[
          {
            "value":{
              "messaging_product":"whatsapp",
              "metadata":{
                "display_phone_number":"5511952992327",
                "phone_number_id":"107768129070051"
              },
              "contacts":[
                {
                  "profile":{
                    "name":"Rafael Carvalho"
                  },
                  "wa_id":"5511966896414"
                }
              ],
              "messages":[
                {
                  "from":"5511966896414",
                  "id":"wamid.HBgNNTUxMTk2Njg5NjQxNBUCABIYFjNFQjAyMUQ0Mzk5Q0IzNDk1MjQ2NDUA",
                  "timestamp":"1714959271",
                  "text":{
                    "body":"yyyyyyyyyyy"
                  },
                  "type":"text"
                }
              ]
            },
            "field":"messages"
          }
        ]
      }
    ]
  },
  {
    "object":"whatsapp_business_account",
    "entry":[
      {
        "id":"110880075418235",
        "changes":[
          {
            "value":{
              "messaging_product":"whatsapp",
              "metadata":{
                "display_phone_number":"5511952992327",
                "phone_number_id":"107768129070051"
              },
              "contacts":[
                {
                  "profile":{
                    "name":"Rafael Carvalho"
                  },
                  "wa_id":"5511966896414"
                }
              ],
              "messages":[
                {
                  "from":"5511966896414",
                  "id":"wamid.HBgNNTUxMTk2Njg5NjQxNBUCABIYFjNFQjA4RjhBMTBCOUZGNjI4NTQ1RUYA",
                  "timestamp":"1714959668",
                  "text":{
                    "body":"somente msg"
                  },
                  "type":"text"
                }
              ]
            },
            "field":"messages"
          }
        ]
      }
    ]
  }
]
```


# DEEP BOT

```view
Meta integration project for managing messages via WhatsApp
```

## Settings

```view

1 - go run cmd/api/main.go
2 - ngrok http 3000
3 - confirme através do primeiro get
4 - create table users
    CREATE TABLE users (
        id text PRIMARY KEY,
        "name" text NOT null,
        email text NOT null,
        password text NOT null,
        phone text NOT NULL
    );
5 - crie uma tabela de messages
    CREATE TABLE messages (
        id text PRIMARY KEY,
        message text NOT null,
    );

```

## Project directory structure

### `/cmd`

```view
Principais aplicações para este projeto.
O nome do diretório para cada aplicação deve corresponder ao nome do executável que você deseja ter.
```

### `/internal`

```view
Aplicação privada e código de bibliotecas.
Este é o código que você não quer que outras pessoas importem em suas aplicações ou bibliotecas.
```

### `/api`

```view
Especificações OpenAPI/Swagger, arquivos de esquema JSON, arquivos de definição de protocolo.
```

### `/configs`

```view
Modelos de arquivo de configuração ou configurações padrão.
Coloque seus arquivos de modelo confd ou consul-template aqui.
```

### `/scripts`

```view
Scripts para executar várias operações de construção, instalação, análise, etc.
Esses scripts mantêm o Makefile de nível raiz pequeno e simples.
```

### `/build`

```view
Empacotamento e integração contínua.
Coloque suas configurações de pacote e scripts em nuvem (AMI), contêiner (Docker), sistema operacional (deb, rpm, pkg) no diretório /build/package.
```

### `/test`

```view
Aplicações de testes externos adicionais e dados de teste.
```

### `migrations`

```bash
   migrate create -ext=postgresql -dir=sql/migrations -seq "name of the migration"

#   $ migrate -source file://path/to/migrations -database postgres://localhost:5432/database up 2

#$ docker run -v {{ migration dir }}:/migrations --network host migrate/migrate
#    -path=/migrations/ -database postgres://localhost:5432/database up 2
```

### `sqlc`

```bash
    Não usaremos o sqlc porq gera muito codigo desnecessario sendo que podemos fazer através do nativo
```

### `Dependencies`

```view
    golang.org/x/crypto
    go get -u github.com/golang-jwt/jwt/v5
```

### (AIzaSyDhYBsS65ncfIByRmHAgM_hbkiFSdKV2aU)