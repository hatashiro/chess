import {html, render} from 'https://unpkg.com/lit-html@1.2.1/lit-html.js?module';
import {loadGame} from './lib/game.js';
import {promotion as renderPromotion} from './components/promotion.js';
import {player as renderPlayer} from './components/player.js';
import {board as renderBoard} from './components/board.js';
import {game, chess} from './lib/enums.js';
import {unregisterPlayer, reset} from './lib/action.js';

const RENDER_INTERVAL = 1000;

const app = ({
  phase,
  players: playerNames,
  state,
}) => {
  const player = state.turn;
  const playerName = playerNames[player];

  const indicators = {
    [chess.Player.P1]: null,
    [chess.Player.P2]: null,
  };
  if (phase === game.Phase.ACTIVE) {
    indicators[player] = 'Turn';
  } else if (phase === game.Phase.DONE) {
    indicators[-player] = 'Win';
  }

  return html`
    <div class="content">
      ${renderPlayer({
        player: chess.Player.P1,
        playerName: playerNames[chess.Player.P1],
        indicator: indicators[chess.Player.P1],
      })}

      ${phase !== game.Phase.WAITING ?
        renderBoard(state) : null
      }

      ${renderPlayer({
        player: chess.Player.P2,
        playerName: playerNames[chess.Player.P2],
        indicator: indicators[chess.Player.P2],
      })}

      <p>
        ${phase === game.Phase.WAITING && (playerNames[chess.Player.P1] || playerNames[chess.Player.P2]) ?
          html`<button @click=${unregisterPlayer}>Unregister</button>` : null
        }

        ${phase === game.Phase.DONE ?
          html`<button @click=${reset}>Reset</button>` : null
        }
      </p>

      ${state.promotion ?
        renderPromotion({player, playerName: playerNames[player]}) : null
      }
    </div>
  `;
}

async function renderApp() {
  render(app(await loadGame()), document.body);
}

renderApp();
setInterval(renderApp, RENDER_INTERVAL);
