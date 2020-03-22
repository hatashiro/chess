import './adjust_board.js';
import './action.js';
import {getGameId, loadGame} from './game.js';
import {render} from './render.js';

async function renderGame() {
  render(await loadGame());
}

renderGame();
setInterval(renderGame, 1000);
