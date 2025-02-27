# Pour lancer le projet :

## Prérequis :
    - Mysql
    - Golang

## Sur mysql en utilisateur root sur votre ordinateur dans un terminal ,vous allez créer un utilisateur avec un mot de passe :
    Utilisateur : AdminSupreme
    Mot de passe : AdminSupreme123! 

Commande à entrée :
```
CREATE USER 'AdminSupreme'@'localhost' IDENTIFIED BY 'AdminSupreme123!';
```
### Créer une base de donnée :
    Nom de la base de donnée : Database_Forum

Commande à entrée :
```
CREATE DATABASE Database_Forum;
```
### Mettre les bons privilèges au nouvel utilisateur :
Cette commande est pour mettre les droits sur toute la base de donnée du projet :
```
GRANT ALL PRIVILEGES ON Database_Forum.* TO 'AdminSupreme'@'localhost' WITH GRANT OPTION;
```
Et celle là pour rafraîchir les privilèges :
```
FLUSH PRIVILEGES;
```
>--------------------------------------
## Une fois tout ça effectué on peut faire : 
Sur terminal dans le dossier du ```Faux_Rome``` , faites cette commande pour lancer le code du site :

```go
go run ./main.go
```
>--------------------------------
## Après tout ca Allez sur votre navigateur et mettez ça dans l'url : 
```
http://localhost:1678
```