# Contribuer à Mamenzul/sesam-vital-core-utils

Merci de votre intérêt pour ce projet. Voici comment contribuer efficacement.

## 📋 Prérequis

- Go >= 1.24
- Respect des conventions Go : `gofmt`, `golint`, `go vet`
- Adopter l'architecture DDD du projet

## 📦 Installer

```bash
git clone https://github.com/Mamenzul/sesam-vital-core-utils.git
cd sesam-vital-core-utils
go mod tidy
```

## 🧪 Lancer les tests

```bash
go test ./...
```

## 🧑‍💻 Proposer une contribution

1. Forkez le projet
2. Créez une branche : `feat/nom-fonctionnalite`
3. Commitez avec un message clair
4. Soumettez une pull request

## 📜 Code style

- Fichiers de domaine = pas de dépendance externe
- Garder les structures Go natives simples
- Toujours tester une règle métier avec au moins 2 cas

Merci 🙏
