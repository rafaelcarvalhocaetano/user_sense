# Whatsapp

## URL: https://xodo-message.uc.r.appspot.com

## STEP

```txt


- Na raiz do projeto precisa ter um app.yaml
- Dentro do projeto execute os comandos
- O projeto só roda na porta 8080

-  gcloud projects list
-  gcloud config set project PROJECT_ID

1 - gcloud auth application-default xodo-message (xodo-message é um project existente na GCP)
2 - gcloud auth application-default login
3 - gcloud app deploy
4 - gcloud app browser
5 - gcloud app logs tail -s default

```


```txt

descriptor:                  [/home/rafael/github/wa_xodo/app.yaml]
source:                      [/home/rafael/github/wa_xodo]
target project:              [xodo-message]
target service:              [default]
target version:              [20240616t215423]
target url:                  [https://xodo-message.uc.r.appspot.com]
target service account:      [xodo-message@appspot.gserviceaccount.com]

```



## Mensagem 1
```txt
- template com imagem e btn ruim, bom ou exceleten

header: <imagem>

body: Você acabou de abastecer em um dos posto da Rede Xodó!

buttons: Avalie o abastecimento
```



## Mensagem 2
```txt
    Quer DESCONTO NO COMBUSTÍVEL todos os dias?
    Conheça o STP Club, e economize muito na hora de abastecer!
    Acesse o link e descubra todas as vantagens que só o STP Club tem.
    
    <https://www.instagram.com/reel/C4tBLGeuON2/?igsh=YW5ta3MxMDFjczF1>
```