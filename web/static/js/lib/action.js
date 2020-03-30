import {gameId} from './game.js';
import {getSessionId} from './session.js';

async function sendAction(action, body, {showAlert = true} = {}) {
  const res = await fetch(`/game/${gameId}/${action}`, {
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

export function registerPlayer(player) {
  const name = prompt("Your name?");
  sendAction('register', {player, name});
}

export function unregisterPlayer() {
  sendAction('unregister');
}

export async function move(from, to) {
  try {
    await sendAction('move', {from, to}, {showAlert: false});
  } catch (err) {
    console.error(err);
  }
}

export async function promote(to) {
  try {
    await sendAction('promote', {to}, {showAlert: false});
  } catch (err) {
    console.error(err);
  }
}

export async function reset() {
  await sendAction('reset');
}
