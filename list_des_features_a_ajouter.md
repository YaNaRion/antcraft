# Feature à améliorer/modifier
**- Faire une meilleure gestion des websockets et des connexions client du côté server**
- Modifier le moment de connection du client
    - Faire en sorte que le client se connect à la partie lorsqu'il fait P non lors de l'initialisation du client
- Faire en sorte que le client recoit du server suite à sa connection à la partie son ID
    - Ajouter un event de réponse que le server envoie
    - Ajouter un event de réception dans le client qui lui permet de récupérer son ID
- Déplacement des élements qui fonctionnent sans problème sans obstacle 
    - Modifier la boucle de mise à jour coté client afin de pouvoir détecter si des unités ont été enlever

# Feature à créer
- Mise en place de partie à deux joueurs pour de vrai, ou des unités peuvent s'attaquer
    - Faire une gestion des joueurs lors de la connection à la partie
        - Gérer la création d'unité (voir à ajouter un event client-side qui ordonne au serveur de lui créer une nouvelle unitée)
    - Faire la différentiation des unités dans la couleur
    - Faire la gestion de l'attaque d'une unité
        - Ajout d'une requete client attaque
        - Ajout d'une reception server à attaque
    - Faire la gestion des dégas entre les unités
        - Faire un visuel entre l'unité attaquée et l'unité attaquante (genre un tir)
- Création d'un CC/TownCenter suite à l'action d'un worker
    - Mieux définir les différentes unités de base
    - Meilleure gestion des différents éléments sur l'écran
    - Mode construction où le joueur peut visualiser l'emplacement de son batiment
- Creation du system de création d'unité
    - Mise en place du supply 
    - Nécessité de construire les supplys pour augmenter sa limite
    - Ajouter la freature de création de worker gràce à un hotkey
- Faire un in game HUD
    - Apprendre à utiliser imgui pour l'utiliser comme hud
    - Faire une scene de menu
    - Faire une scene de game avec les features qui sont implémentées
        - Faire HUD de game avec les options en fonction de l'unité selectionnée 
            - Si batiment, les options d'unités à créer
            - Si unité, les options de constructions, movement et effets spéciaux

# Feature à créer optionnelle
- Ajouter une option de replay la partie
    - Voir comment gérer les données, comment les sauvegarder, voir le type de base de donnée à utiliser

