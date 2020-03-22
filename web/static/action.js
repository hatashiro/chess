import {$} from './query.js';
import {getGameId} from './game.js';
import {getSessionId} from './session.js';
import * as enums from './enums.js';

async function sendAction(action, body) {
  const res = await fetch(`/game/${getGameId()}/${action}`, {
    method: 'post',
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify({
      session: getSessionId(),
      ...body,
    }),
  });

  if (res.status !== 200) {
    const text = await res.text();
    alert(`Error: ${text}`);
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
