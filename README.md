# Mamenzul/sesam-vital-core-utils

**Mamenzul/sesam-vital-core-utils** est une bibliothèque open source en Go destinée à fournir les outils essentiels de gestion de facturation SESAM-Vitale pour les éditeurs de logiciels santé.

Elle est conçue pour être :

- conforme au CDC SESAM-Vitale v1.40 Addendum 8
- écrite en Go pur (librairie standard uniquement)
- orientée DDD (Domain-Driven Design)
- intégrable nativement dans un backend Go + SQL

## ✨ Fonctionnalités

- Parsing de fichiers de facturation (CSV)
- Validation métier selon les règles SESAM-Vitale (RG_xxx)
- Génération des flux FSE/DRE
- Sécurisation : signature, chiffrement (Annexe 4)
- Génération des fichiers à transmettre

## 🧱 Architecture

```
CSV → Parse → Facture (Go struct) → Validation → Flux sécurisé → Fichier
```

## 🚀 Démarrage rapide

```bash
git clone https://github.com/Mamenzul/sesam-vital-core-utils.git
cd sesam-vital-core-utils
go build ./cmd/coreutils-cli
```

## 📄 Conformité

Ce projet suit le Cahier des Charges Éditeurs SESAM-Vitale :

- CDC v1.40 Addendum 8
- Normes B2-2007, DRE-2009

## 📌 Avertissement

Le projet ne vise **pas** à implémenter un client complet ou frontal. Il constitue le cœur technique **réutilisable dans des logiciels métiers**.

## 🧑‍💻 Contribuer

Voir [`CONTRIBUTING.md`](./CONTRIBUTING.md)

## 🛣️ Roadmap

Voir [`ROADMAP.md`](./ROADMAP.md)

## 🧠 🧭 Scénario réel de facturation SESAM-Vitale

Voici le workflow typique dans un cabinet médical :

    🪪 Le patient présente sa carte Vitale

    📥 Le logiciel lit la carte via un lecteur (interface PC/SC ou autre)

    👨‍⚕️ Le PS s’identifie via sa carte CPS (authentification)

    📝 Le praticien renseigne l’acte ou les actes réalisés

    🧮 Le logiciel valorise la facture selon les règles (convention, tarifs, exonérations…)

    📤 Le logiciel génère un flux SESAM-Vitale (FSE ou DRE)

    🔐 Le flux est signé, chiffré, puis transmis

    📥 Les ARL (accusés de réception) et RSP (rejets/paiements) sont récupérés
