# Feature à améliorer/modifier
**- Faire une meilleure gestion des websockets et des connexions client du côté server**
- Modifier le moment de connection du client
    - [x] Faire en sorte que le client se connect à la partie lorsqu'il fait P non lors de l'initialisation du client
    - [x] Faire en sorte que le client recoit du server suite à sa connexion à la partie son ID
    - [x] Ajouter un event de réponse que le server envoie
    - [x] Ajouter un event de réception dans le client qui lui permet de récupérer son ID
- Déplacement des élements qui fonctionnent sans problème sans obstacle 
    - Modifier la boucle de mise à jour coté client afin de pouvoir détecter si des unités ont été enlever

# Description du MVP
- Avec une logique pour faire la map avec des murs comme obstacle et donc algo de déplacement évitant les obstacles (djstra ou breathforsearch)
- Pouvoir jouer une partie avec des unitées très simple
    - Batiment disponible
        - TownCenter, supply, barraque
    - Unité disponible
        - Worker
        - Unité de base de combat
    - 1 Ressource disponible
- Avoir un hud de base avec hotkey hardcodé


# Feature pour MVP (minimum viable product)
- Mise en place de partie à deux joueurs pour de vrai
    - [x] Faire une gestion des joueurs lors de la connection à la partie
        - Gérer la création d'unité (voir à ajouter un event client-side qui ordonne au serveur de lui créer une nouvelle unitée)
    - [x] Faire la différentiation des unités dans la couleur
- Faire une gestion de map, comment créer une map dans le jeu
    - [ ] Créer plusieurs tiles avec leur textures animés qui change dans le temps
    - [ ] Faire une différence entre les coordonnées frontend et celle backend
        - Backend = postion dans la grille
        - Frontend = position absolue sur la window, faire en sorte que
    - [ ] Faire une gestion de la map qui inclus les type de tiles 
- Création d'un CC/TownCenter suite à l'action d'un worker
    - Mieux définir les différentes unités de base
    - Meilleure gestion des différents éléments sur l'écran
    - Mode construction où le joueur peut visualiser l'emplacement de son batiment
- Mise en place de l'intéraction entre les unités opposés
    - [ ] Faire la gestion de l'attaque d'une unité (backend)
        - Ajout d'une requete client attaque
        - Ajout d'une reception server à attaque
    - [ ] Faire la gestion des dégas entre les unités (frontend)
        - [ ] Faire un visuel entre l'unité attaquée et l'unité attaquante (genre un tir)
- Faire un in game HUD
    - Apprendre à utiliser imgui pour l'utiliser comme hud
    - Faire une scene de menu
    - Faire une scene de game avec les features qui sont implémentées
        - Faire HUD de game avec les options en fonction de l'unité selectionnée 
            - Si batiment, les options d'unités à créer
            - Si unité, les options de constructions, movement et effets spéciaux

# Feature à créer
- Creation du system de création d'unité
    - Mise en place du supply 
    - Nécessité de construire les supplys pour augmenter sa limite
    - Ajouter la freature de création de worker gràce à un hotkey
- Déplacement
    - [ ] Faire un algorithme de déplacement d'unité pour que l'unité évite les murs
- Faire des unités animés
    - https://zylinski.se/posts/gamedev-for-beginners-using-odin-and-raylib-3/


# Gameplay
- [ ] Système de ressource à déterminer
    - Aller retour comme dans SC2 prêt base, mais avec une 3ème plus loin
- [ ] Map

# Feature à créer optionnelle
- Ajouter une option de replay la partie
    - Voir comment gérer les données, comment les sauvegarder, voir le type de base de donnée à utiliser

