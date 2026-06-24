# Kuaforia - Casos como Código

Este directorio contiene la definición de casos de negocio en formato YAML,
versionados junto con el código fuente.

## Estructura

- `.kuaforia.yaml` — Configuración del workspace
- `casos/` — Casos de negocio
- `examples/` — Ejemplos de referencia

## Uso

```bash
# Validar todos los casos
kuaforia validate ./kuaforia/

# Sincronizar con Kuaforia
kuaforia sync           # push (local → remoto)
kuaforia sync --pull    # pull (remoto → local)

# Importar un caso individual
kuaforia import casos/mi-caso.yaml
```
