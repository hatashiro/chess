import {$} from './query.js';
import {MAX_RANK, intLocation} from './game.js';
import * as enums from './enums.js';

export function render(game) {
  renderPlayer(game, enums.chess.Player.P1);
  renderPlayer(game, enums.chess.Player.P2);

  if (game.phase === enums.game.Phase.WAITING) {
    hide('info');
    hide('board');
  } else {
    show('info');
    show('board');
    renderInfo(game);
    renderState(game.state);
  }
}

function show(id) {
  $(`#${id}`).style.display = '';
}

function hide(id) {
  $(`#${id}`).style.display = 'none';
}

function renderPlayer(game, player) {
  const name = game.players[player];
  const prefix = player === enums.chess.Player.P1 ? 'p1' : 'p2';
  if (name) {
    show(`${prefix}-name`);
    hide(`${prefix}-register`);
    $(`#${prefix}-name`).textContent = name;

    if (game.phase === enums.game.Phase.WAITING) {
      show('unregister');
    } else {
      hide('unregister');
    }
  } else {
    hide(`${prefix}-name`);
    show(`${prefix}-register`);
    hide('unregister');
  }
}

function renderInfo(game) {
  const turn = game.state.turn;
  const name = game.players[turn];
  $('#turn').textContent = name;
}

let lastUpdated = 0;
function renderState(state) {
  if (lastUpdated === state.lastUpdated) {
    return;
  }

  renderBoard(state.board);

  lastUpdated = state.lastUpdated;
}

function renderBoard(board) {
  for (let row = 0; row < MAX_RANK; row++) {
    for (let col = 0; col < MAX_RANK; col++) {
      const $cell = $(`#cell-${row}-${col}`);
      const piece = board[intLocation(row, col)];
      $cell.textContent = piece ? renderPiece(piece) : '';
    }
  }
}

const pieceSymbolMap = {
  [enums.chess.Player.P1]: {
    [enums.chess.PieceType.KING]:   '♚',
    [enums.chess.PieceType.QUEEN]:  '♛',
    [enums.chess.PieceType.ROOK]:   '♜',
    [enums.chess.PieceType.BISHOP]: '♝',
    [enums.chess.PieceType.KNIGHT]: '♞',
    [enums.chess.PieceType.PAWN]:   '♟',
  },
  [enums.chess.Player.P2]: {
    [enums.chess.PieceType.KING]:   '♔',
    [enums.chess.PieceType.QUEEN]:  '♕',
    [enums.chess.PieceType.ROOK]:   '♖',
    [enums.chess.PieceType.BISHOP]: '♗',
    [enums.chess.PieceType.KNIGHT]: '♘',
    [enums.chess.PieceType.PAWN]:   '♙',
  },
};

function renderPiece(piece) {
  return pieceSymbolMap[piece.owner][piece.type];
}
