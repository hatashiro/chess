import {$} from './query.js';
import {MAX_RANK, intLocation} from './game.js';
import * as enums from './enums.js';

export function render(game) {
  renderPlayer(game, enums.chess.Player.P1);
  renderPlayer(game, enums.chess.Player.P2);

  hide('#board');
  hide('#promotion');
  if (game.phase !== enums.game.Phase.WAITING) {
    show('#board');

    renderState(game.state);

    if (game.state.promotion) {
      show('#promotion');
      renderPromotion(game);
    }
  }
}

function show(selectorOrEl) {
  if (typeof selectorOrEl === 'string') {
    $.all(selectorOrEl).forEach(($el) => {
      $el.style.display = '';
    });
  } else {
    selectorOrEl.style.display = '';
  }
}

function hide(selectorOrEl) {
  if (typeof selectorOrEl === 'string') {
    $.all(selectorOrEl).forEach(($el) => {
      $el.style.display = 'none';
    });
  } else {
    selectorOrEl.style.display = 'none';
  }
}

function renderPlayer(game, player) {
  const isMyTurn = game.state.turn === player;
  const name = game.players[player];
  const prefix = player === enums.chess.Player.P1 ? 'p1' : 'p2';

  const $name = $(`#${prefix}-name`);
  const $turn = $(`#${prefix}-turn`)
  const $register = $(`#${prefix}-register`)
  const $unregister = $('#unregister')

  hide($name);
  hide($turn);
  hide($register);
  hide($unregister);

  if (name) {
    show($name);
    $name.textContent = name;
    if (game.phase === enums.game.Phase.WAITING) {
      show($unregister);
    } else if (isMyTurn) {
      show($turn);
    }
  } else {
    show($register);
  }
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
      if (piece) {
        $cell.textContent = renderPiece(piece);
        $cell.classList.add('piece');
      } else {
        $cell.textContent = '';
        $cell.classList.remove('piece');
      }
      $cell.classList.remove('from');
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

function renderPromotion(game) {
  const player = game.state.turn;
  const name = game.players[player];
  const className = player === enums.chess.Player.P1 ? "p1" : "p2"
  hide('#promotion > button');
  show(`#promotion > button.${className}`);
  $('#promotion .name').textContent = name;
}
