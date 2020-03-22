export const MAX_RANK = 8;

export function getGameId() {
  return location.pathname.split('/')[1];
}

export async function loadGame() {
  const res = await fetch(`/game/${getGameId()}`);
  return res.json();
}

export function locationFromInt(intLoc) {
  const row = Math.floor(intLoc / MAX_RANK);
  const col = intLoc % MAX_RANK;
  return {row, col};
}

export function intLocation(row, col) {
  return row * MAX_RANK + col;
}
