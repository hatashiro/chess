import {$} from './query.js';

let height = 0;

function adjustBoard() {
  const board = $('#board');
  if (height !== board.offsetWidth) {
    height = board.offsetWidth;
    board.style.height = `${height}px`;

    const fontSize = height / 8;
    board.style.lineHeight = `${fontSize}px`;
    board.style.fontSize = `${fontSize}px`;
  }
}

setInterval(adjustBoard, 100);
