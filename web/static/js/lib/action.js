import {$} from './query.js';
import {getGameId, intLocation} from './game.js';
import {getSessionId} from './session.js';
import * as enums from './enums.js';

async function sendAction(action, body, {showAlert = true} = {}) {
  const res = await fetch(`/game/${getGameId()}/${action}`, {
    method: 'post',
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify({
      session: getSessionId(),
      ...body,
    }),
  });

  if (res.status !== 200) {
    if (showAlert) {
      const text = await res.text();
      alert(`Error: ${text}`);
    } else {
      throw new Error(await res.text());
    }
  }
}

function registerPlayer(player) {
  const name = prompt("Your name?");
  sendAction('register', {player, name});
}

$('#p1-register').addEventListener(
  'click',
  () => registerPlayer(enums.chess.Player.P1));

$('#p2-register').addEventListener(
  'click',
  () => registerPlayer(enums.chess.Player.P2));

function unregisterPlayer() {
  sendAction('unregister');
}

$('#unregister').addEventListener('click', unregisterPlayer);

function getIntLocation($el) {
  const [, row, col] = $el.id.split('-').map(x => Number(x));
  return intLocation(row, col);
}

function selectCell(e) {
  const $from = $('.from');

  const $cell = e.target;
  if ($from) {
    if ($from === $cell) {
      // Unselect.
      $cell.classList.remove('from');
    } else {
      // Move
      move(getIntLocation($from), getIntLocation($cell));
    }
  } else if ($cell.classList.contains('piece')) {
    $cell.classList.add('from');
  }
}

$('#board').addEventListener('click', selectCell);

async function move(from, to) {
  try {
    await sendAction('move', {from, to}, {showAlert: false});
  } catch (err) {
    // Ignore move errors and unselect $from.
    console.error(err);
    $('.from').classList.remove('from');
  }
}

async function promote(e) {
  if (e.target.tagName !== 'BUTTON') return;

  const to = Number(e.target.id.split('-')[1]);

  try {
    await sendAction('promote', {to}, {showAlert: false});
  } catch (err) {
    // Ignore move errors and unselect $from.
    console.error(err);
  }
}

$('#promotion').addEventListener('click', promote);

async function reset() {
  await sendAction('reset');
}

$('#reset').addEventListener('click', reset);
