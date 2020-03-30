export const gameId = location.pathname.split('/')[1];

export async function loadGame() {
  const res = await fetch(`/game/${gameId}`);
  return res.json();
}
