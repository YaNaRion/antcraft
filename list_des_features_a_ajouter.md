# Feature à améliorer/modifier
- Modifier le moment de connection du client
    - Faire en sorte que le client se connect à la partie lorsqu'il fait P non lors de l'initialisation du client
- Faire en sorte que le client recoit du server suite à sa connection à la partie son ID
    - Ajouter un event de réponse que le server envoie
    - Ajouter un event de réception dans le client qui lui permet de récupérer son ID
- Déplacement des elements qui fonctionnent sans problème sans obstacle 
    - Modifier la boucle de mise à jour coté client afin de pouvoir détecter si des unités ont été enlever

# Feature à créer
- Mise en place de partie à deux joueurs pour de vrai, ou des unités peuvent s'attaquer
    - Faire une gestion des joueurs lors de la connection à la partie
        - Gérer la création d'unité (voir à ajouter un event client-side qui ordonne au serveur de lui créer une nouvelle unitée)
    - Faire la gestion de l'attaque d'une unité
        - Ajout d'une requete client attaque
        - Ajout d'une reception server à attaque
    - Faire la gestion des dégas entre les unités
        - Faire un visuel entre l'unité attaquée et l'unité attaquante (genre un tir)
- Création d'un CC/TownCenter suite à l'action d'un worker
    - Mieux définir les différentes unités de base
    - Meilleur gestion des différents éléments sur l'écran
    - Mode construction où le joueur peut visualiser l'emplacement de son batiment
