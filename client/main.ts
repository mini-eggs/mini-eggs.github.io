import turbolinks from "turbolinks";
import marked from "marked";
import highlight from "highlight.js/lib/highlight.js";
import "./styles/main.scss";
import "./styles/aside.scss";
import "./styles/page.scss";
import "./styles/blog.scss";
import "highlight.js/styles/ocean.css";

highlight.registerLanguage("json", require("highlight.js/lib/languages/json"));
highlight.registerLanguage("javascript", require("highlight.js/lib/languages/javascript"));
highlight.registerLanguage("bash", require("highlight.js/lib/languages/bash"));

// flag hash, set to true to notify we have already loaded the image(s).
let images = {};

let fadeIn = node => {
  let src = node.getAttribute("src");
  if (!images[src]) {
    images[src] = true;
    node.className += " fade-in";
  }
};

let markdownToHTML = str => {
  let renderer = new marked.Renderer();

  renderer.code = (code, language) => {
    let tmp = document.createElement("div");
    tmp.innerHTML = code;
    //code = tmp.textContent;

    let valid = !!(language && highlight.getLanguage(language));
    let highlighted = valid ? highlight.highlight(language, code).value : code;

    return `<pre><code class="hljs ${language}">${highlighted}</code></pre>`;
  };

  return marked(str, { renderer });
};

document.addEventListener("DOMContentLoaded", () => {
  // Fade in aside when image is done loading
  // Only do this once!
  let asideEl = document.querySelector("aside");
  let asideImg = asideEl.querySelector("img");
  asideImg.addEventListener("load", () => fadeIn(asideEl));
});

document.addEventListener("turbolinks:before-render", event => {
  // We already faded in the aside
  // don't hide it!
  // @ts-ignore
  let body: HTMLBodyElement = event.data.newBody;
  let aside = body.querySelector("aside");
  aside.style.opacity = "1";

  // Show the images we have already displayed!
  for (let img of body.querySelectorAll("img")) {
    let src = img.getAttribute("src");
    if (images[src]) {
      img.style.opacity = "1";
    }
  }
});

document.addEventListener("turbolinks:load", () => {
  // Fade in article image as well
  let mainEl = document.querySelector("main");
  for (let img of mainEl.querySelectorAll("img")) {
    img.addEventListener("load", () => fadeIn(img));
  }

  // Display code/markdown correctly
  let articleEl = document.querySelector("template");
  if (articleEl) {
    let el = document.createElement("article");
    el.innerHTML = articleEl.innerHTML;
    el.innerHTML = markdownToHTML(el.textContent);
    articleEl.parentElement.replaceChild(el, articleEl);
  }
});

turbolinks.start();
