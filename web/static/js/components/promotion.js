import {html} from 'https://unpkg.com/lit-html@1.2.1/lit-html.js?module';
import {piece as renderPiece} from '../lib/piece.js';
import {chess} from '../lib/enums.js';
import {promote} from '../lib/action.js';

const pieceButton = ({player, type}) => html`
  <button @click=${() => promote(type)}>
    ${renderPiece({player, type})}
  </button>
`;

export const promotion = ({player, playerName}) => html`
  <div id="promotion">
    <p>${playerName}, please choose a piece type for the promotion.</p>
    ${pieceButton({player, type: chess.PieceType.QUEEN})}
    ${pieceButton({player, type: chess.PieceType.ROOK})}
    ${pieceButton({player, type: chess.PieceType.BISHOP})}
    ${pieceButton({player, type: chess.PieceType.KNIGHT})}
  </div>
`;


