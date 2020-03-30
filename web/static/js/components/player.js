import {html} from 'https://unpkg.com/lit-html@1.2.1/lit-html.js?module';
import {chess} from '../lib/enums.js';
import {registerPlayer} from '../lib/action.js';

const label = (player) => player === chess.Player.P1 ? 'P1' : 'P2';

export const player = ({player, playerName, indicator}) => html`
  <p>
    ${label(player)}:

    ${playerName ?
      html`<span>${playerName}</span>` :
      html`<button @click=${() => registerPlayer(player)}>Register</button>`
    }

    ${indicator ?
      html`<span class="indicator">← ${indicator}</span>` : null
    }
  </p>
`;
