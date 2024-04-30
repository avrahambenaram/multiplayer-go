class Player {
  constructor(id, x, y, points) {
    this.id = id;
    this.x = x;
    this.y = y;
    this.points = points;
  }

  moveLeft() {
    this.x--;
  }

  moveRight() {
    this.x++;
  }

  moveUp() {
    this.y--;
  }

  moveDown() {
    this.y++;
  }
}

class Fruit {
  constructor(x, y, type) {
    this.x = x;
    this.y = y;
    this.type = type;
  }
}

class Game {
  constructor() {
    this.players = [];
    this.cleanFruits();
  }

  cleanFruits() {
    this.fruits = {
      0: [],
      1: [],
      2: [],
      3: [],
      4: [],
      5: [],
      6: [],
      7: [],
      8: [],
      9: [],
    };
  }

  addPlayer(player) {
    this.players.append(player);
  }

  removePlayer(playerId) {
    const player = this.players.find(player => player.id === playerId);
    const playerIndex = this.players.indexOf(player);
    if (playerIndex == -1) return;
    this.players.splice(playerIndex, 1);
  }

  movePlayer(props) {
    const player = this.players.find(player => player.id === props.playerId);
    if (!player) return;

    switch (props.direction) {
      case "left":
        player.moveLeft();
        break
      case "right":
        player.moveRight();
        break
      case "up":
        player.moveUp();
        break
      case "down":
        player.moveDown();
        break
    }
  }
}
