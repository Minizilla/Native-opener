# Microzilla

Un outil pour enregistrer des handlers d'URI personnalisés sur Linux, Windows et macOS.

## Comment ça fonctionne

- Microzilla sauvegarde dans le fichier de config de votre OS que ce dernier doit appeler un wrapper go quand il recoit `monapp://monfichier.pdf`

- Ce wrapper s'occupe de récuperer lees arguments comme `monfichier.pdf` et les passe à l'application que vous avait demandé dans l'enregistrelent de l'ui

## Utilisation

### 1. Compiler les outils

```bash
make all
```

### 2. Enregistrer un handler d'URI

```bash
./microzilla monapp /chemin/vers/votre/programme/a/executer
```

### 3. Utiliser l'URI

Quand quelqu'un clique sur `monapp://monfichier.pdf`, votre programme sera lancé avec `monfichier.pdf` comme argument.

