---
layout: default
title: Understanding the cryptographic entities
nav_order: 6
description: "Explaining the key concepts for encrypting user data."
permalink: /developing-khulnasoft/understanding-cryptography/
parent: For developers
---

<!--
Copyright 2020 - KhulnaSoft Authors <admin@khulnasoft.com>
SPDX-License-Identifier: Apache-2.0
-->

# Understanding the cryptographic entities in use

Event data in KhulnaSoft Fair Web Analytics is encrypted before leaving the user's browser and will also be stored like this at rest. Decryption only happens on the client side. In order to share data between users and operators the following cryptographic entities will be used:

## Account keypair

Each account owns a unique RSA keypair. The public key can be accessed by anyone that knows about the account's `AccountId`. This way users can encrypt secrets and make them available to certain accounts only.

## User secrets

Before a user sends data to an instance of KhulnaSoft Fair Web Analytics for the first time, the following procedure which roughly resembles a PGP exchange will happen:

- in the client, a random symmetric `UserSecret` will be created and persisted locally
- the account's public key will be used to encrypt the `UserSecret`
- the `EncryptedUserSecret` is sent to the server and associated with a hashed version of the client's user id, thus enabling the operator to generate a decrypted `UserSecret` when required
- event data is encrypted using the `UserSecret` and sent to the server

## Decrypting event data

### Users

Users can decrypt any event they sent using their local version of the respective `UserSecret`

### Operators

Operators can decrypt all events belonging to their account by decrypting the encrypted `UserSecret`s and then decrypting the event payloads using these secrets.
