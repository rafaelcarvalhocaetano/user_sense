# Typesense

## Usando banco de dado com registros

Criar um script ou codigo para ler os dados do banco e popular o document do typesense

psql -d sua_base_de_dados -c "COPY (SELECT row_to_json(t) FROM sua_tabela t) TO STDOUT;" > dados_para_typesense.jsonl

```shell
curl "http://localhost:8108/collections/nome_da_sua_colecao/documents/import?action=create" \
  -X POST \
  -H "Content-Type: text/plain" \
  -H "X-Typesense-API-Key: sua_api_key" \
  --data-binary @dados_para_typesense.jsonl
```

## Backup do typesense

- Existe um recurso nativo do typesense para fazer backup dos dados (snapshots)
- É um endpoint do typesense para criar um backup do estado e dos dados

- (Estado -->  Configs)
- (Dados  -->  Documentos que você indexou)

```shell
  curl "http://localhost:8108/snapshots?snapshot_path=/tmp/typesense-data-snapshot" \
  -X POST \
  -H "X-Typesense-API-Key: sua_api_key"
```

- Compactando o snapshot

```shell
tar -czvf typesense-backup-$(date +%Y-%m-%d).tar.gz -C /tmp/typesense-data-snapshot .
```

- Processo de envio para o s3 é manual

```shell
aws s3 cp typesense-backup-$(date +%Y-%m-%d).tar.gz s3://seu-bucket-s3/backups/
```


## Estrutura dos Arquivos do Snapshot

📁 db_snapshot/ (Diretório Principal)

Este é o snapshot completo do banco de dados Typesense.

📄 Arquivos de Metadados:

CURRENT - Ponteiro para o snapshot ativo atual
MANIFEST-000013 - Manifesto que lista todos os arquivos do snapshot
OPTIONS-000015 - Configurações e opções do banco
__raft_snapshot_meta - Metadados do Raft (consensus protocol)

📄 Arquivos de Dados:

000009.sst - Arquivo SST (Sorted String Table) com dados
000017.sst - Outro arquivo SST com dados

🎯 Para Migração Completa:
TODOS os arquivos são necessários para restaurar em outro servidor:


```txt
db_snapshot/
├── CURRENT          ← Estado atual
├── MANIFEST-000013  ← Lista de arquivos
├── OPTIONS-000015   ← Configurações
├── 000009.sst      ← Dados (schema + documentos)
├── 000017.sst      ← Dados (schema + documentos)
└── __raft_snapshot_meta ← Metadados Raft
```

```bash
tar -czf typesense_backup_$(date +%Y%m%d_%H%M%S).tar.gz -C /backups db_snapshot/
```