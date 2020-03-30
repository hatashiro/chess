import {html, render} from 'https://unpkg.com/lit-html@1.2.1/lit-html.js?module';

const index = () => html`
  <div class="content">
    <h1>♞♘ CHESS</h1>
    <p>
      <input
        type="text"
        placeholder="Room name..."
        @input=${handleInput}>
    </p>
    <p>
      <button @click=${join}>Join</button>
    </p>
  </div>
`;

let gameId = '';

function handleInput(e) {
  gameId = e.target.value.trim();
}

function join() {
  if (!gameId) {
    alert("Enter a game ID.");
    return;
  }
  location.href = `/${gameId}`
}

render(index(), document.body);
