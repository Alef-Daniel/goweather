# GO Weather

API de previsão do tempo em Go, com cache em Redis e integração à [Visual Crossing Weather API](https://www.visualcrossing.com/weather-api).

## Requisitos

- Go 1.25+
- Redis (para cache)
- Chave de API da [Visual Crossing](https://www.visualcrossing.com/weather-api)

## Configuração

Edite `config/config.json`:

- `weather_api.api_key`: sua chave da Visual Crossing
- `cache.endpoint`: endereço do Redis (`localhost:6379` local, `redis:6379` no Docker)

## Executar localmente

```bash
redis-server
go run ./cmd/api
```

A API sobe em `http://localhost:8080`.

## Executar com Docker

Ajuste `config/config.json` e defina `cache.endpoint` como `redis:6379`. Depois:

```bash
docker compose up --build
```

A API fica em `http://localhost:8080` e o Redis só acessível entre os containers.

## Endpoints

| Método | Caminho | Descrição |
|--------|---------|-----------|
| GET | `/health` | Health check |
| POST | `/api/weather` | Previsão para um local. Body: `{"location": "São Paulo, BR"}` |
| POST | `/api/weather/date` | Previsão por período. Body: `{"location": "...", "date_init": "2025-03-01", "date_end": "2025-03-05"}` |

## Licença

Ver [LICENSE](LICENSE).
