# TON seed recovery utility

Utility to recover corrupted seed phrase for TON wallet (V5R1, V4R2 or V3R2).

## Use cases
* One word lost
* One word invalid

## How to use

### RUN Go program
Write seed to `var SEED = ""` at `main.go` (24 words for seed with invalid word or 23 words for seed with lost word)
If you have error `invalid word: some_word` it means that word is not from default seed dict. You can remove it and retry as `one word lost` case.

### RUN via docker
1. Write seed to `SEED` at `docker-compose.yml` (24 words for seed with invalid word or 23 words for seed with lost word)
2. Run container `docker compose up`
3. Wait for an entry like `can not find valid seed` or `Valid seed` to appear in the logs.
4. Remove container `docker rm -f ton-seed-recovery`
5. It is not recommended to continue using the restored seed phrase, it is better to create a new wallet.