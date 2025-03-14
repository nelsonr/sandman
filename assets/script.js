const button = document.querySelector("button");
const emojiEl = document.querySelector(".emoji");

const emoji = ["🙊", "🙉", "🙈"];
let index = 0;

button.addEventListener("click", () => {
    emojiEl.textContent = emoji[index];
    index++;

    if (index === emoji.length) {
        index = 0;
    }
});
