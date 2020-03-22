const SESSION = 'session';

function generateSessionId() {
  return Math.floor(Math.random() * (Number.MAX_SAFE_INTEGER - 1)) + 1;
}

export function getSessionId() {
  let id = parseInt(localStorage.getItem(SESSION), 10);
  if (!id) {
    id = generateSessionId();
    localStorage.setItem(SESSION, id);
  }
  return id;
}
