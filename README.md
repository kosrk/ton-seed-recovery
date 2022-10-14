# TON seed recovery utility

Utility to recover corrupted seed phrase for TON wallet (V4R2 or V3R2).

## Use cases
* One word lost
* One word invalid

## How to use
Write seed to `const SEED = ""` at `main.go` (24 words for seed with invalid word or 23 words for seed with lost word)
If you have error `invalid word: some_word` it means that word is nt from default seed dict. You can remove it and retry as `one word lost` case.
