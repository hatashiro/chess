import {html} from 'https://unpkg.com/lit-html@1.2.1/lit-html.js?module';
import {piece as renderPiece} from '../lib/piece.js';
import {move} from '../lib/action.js';

const ADJUST_BOARD_INTERVAL = 100;

let height = 0;
function adjustBoard() {
  const $board = document.querySelector('#board');
  if ($board && height !== $board.offsetWidth) {
    height = $board.offsetWidth;
    $board.style.height = `${height}px`;

    const fontSize = height / 8;
    $board.style.lineHeight = `${fontSize}px`;
    $board.style.fontSize = `${fontSize}px`;
  }
}
setInterval(adjustBoard, ADJUST_BOARD_INTERVAL);

let lastUpdated = 0;
let cache;
export const board = (state) => {
  if (lastUpdated < state.lastUpdated) {
    lastUpdated = state.lastUpdated;
    cache = boardImpl({board: state.board});
    clearFrom();
  }
  return cache;
};

const MAX_RANK = 8;

const location = (row, col) => row * MAX_RANK + col;

const indexRange = [...Array(MAX_RANK).keys()];

const boardImpl = ({board}) => html`
  <div id="board">
    ${indexRange.map((row) => html`
      <div>
      ${indexRange.map((col) =>
        cell({row, col, piece: board[location(row, col)]})
      )}
      </div>
    `)}
  </div>
`;

const from = () => document.querySelector('.from');
const setFrom = ($el) => $el.classList.add('from');
const clearFrom = () => document.querySelectorAll('.from')
  .forEach($el => $el.classList.remove('from'));
const getLocation = ($el) => {
  const row = Number($el.dataset.row);
  const col = Number($el.dataset.col);
  return location(row, col);
};

function handleCellClick(e) {
  if (from()) {
    move(getLocation(from()), getLocation(e.target));
    clearFrom();
  } else {
    setFrom(e.target);
  }
}

const cell = ({row, col, piece}) => html`
  ${piece ?
    html`
      <div class="piece" data-row=${row} data-col=${col} @click=${handleCellClick}>
        ${renderPiece({player: piece.owner, type: piece.type})}
      </div>
    ` :
    html`<div data-row=${row} data-col=${col} @click=${handleCellClick}/>`
  }
`;
