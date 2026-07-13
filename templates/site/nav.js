/* nav.js - keyboard navigation across the site.
   On the list pages: left/right switches between blog, ideas and projects,
   up/down moves a highlight through the entry links, enter opens the
   highlighted entry. On a post, idea or project: esc goes back to its list
   page. Purely client-side: the site stays static. */
(function () {
  "use strict";

  var scriptEl = document.querySelector("script[data-blog-root]");
  var root = scriptEl ? scriptEl.getAttribute("data-blog-root") : "";

  var SECTIONS = ["blog", "ideas", "projects"];
  var SECTION_PAGES = {
    blog: "index.html",
    ideas: "ideas.html",
    projects: "projects.html"
  };

  /* Which page is this? The generated nav bar renders the current section as
     a plain <span class="current"> on the list pages ("blog", "ideas" or
     "projects") and no .current on entry pages, so read it from the DOM
     rather than guessing from the URL (which varies with how the site is
     served). Fall back to the URL shape for pages generated before the
     .current nav existed. */
  var currentEl = document.querySelector("header.site-nav .current");
  var section = currentEl ? currentEl.textContent.trim().toLowerCase() : "";
  var isListPage = section !== "" || root === "";
  if (section === "" || SECTIONS.indexOf(section) < 0) {
    section = "blog";
    if (window.location.pathname.indexOf("/ideas") >= 0) {
      section = "ideas";
    } else if (window.location.pathname.indexOf("/projects") >= 0) {
      section = "projects";
    }
  }

  var links = [];
  var selected = -1;

  var styleEl = document.createElement("style");
  styleEl.textContent =
    ".blog-kbd-selected { background: #f2f8fb; outline: 2px solid #5a9ab5; outline-offset: 2px; }";
  document.head.appendChild(styleEl);

  function collectLinks() {
    links = [].slice.call(document.querySelectorAll(".post-list a, .season a"));
  }

  function select(index) {
    if (selected >= 0 && links[selected]) {
      links[selected].classList.remove("blog-kbd-selected");
    }
    selected = index;
    if (selected >= 0 && links[selected]) {
      links[selected].classList.add("blog-kbd-selected");
      links[selected].scrollIntoView({ block: "nearest" });
    }
  }

  function move(delta) {
    collectLinks();
    if (!links.length) {
      return;
    }
    var next;
    if (selected < 0) {
      next = delta > 0 ? 0 : links.length - 1;
    } else {
      next = (selected + delta + links.length) % links.length;
    }
    select(next);
  }

  document.addEventListener("keydown", function (ev) {
    if (ev.altKey || ev.ctrlKey || ev.metaKey || ev.shiftKey) {
      return;
    }
    var target = ev.target;
    if (target && (target.tagName === "INPUT" || target.tagName === "TEXTAREA" ||
        target.tagName === "SELECT" || target.isContentEditable)) {
      return;
    }
    /* The double-shift search popup owns the keyboard while it is open. */
    if (document.querySelector(".blog-search-overlay")) {
      return;
    }

    if (isListPage) {
      var at = SECTIONS.indexOf(section);
      if (ev.key === "ArrowLeft" && at > 0) {
        window.location.href = SECTION_PAGES[SECTIONS[at - 1]];
      } else if (ev.key === "ArrowRight" && at < SECTIONS.length - 1) {
        window.location.href = SECTION_PAGES[SECTIONS[at + 1]];
      } else if (ev.key === "ArrowDown") {
        ev.preventDefault();
        move(1);
      } else if (ev.key === "ArrowUp") {
        ev.preventDefault();
        move(-1);
      } else if (ev.key === "Enter" && selected >= 0 && links[selected]) {
        links[selected].click();
      }
    } else if (ev.key === "Escape") {
      window.location.href = root + SECTION_PAGES[section];
    }
  });
})();
