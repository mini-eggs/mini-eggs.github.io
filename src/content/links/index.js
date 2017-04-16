const resume =
  "https://docs.google.com/document/d/1kF0x3kyEueNo-7yuLIe2xywcLWEREoAGqy8l8XpEuew/edit?usp=sharing";
const github = "https://github.com/mini-eggs";
const twitter = "https://twitter.com/minieggs40";
const email = "mailto:evanjones4040@gmail.com";

const blogChildren = [
  { name: "Emojis and the URL", route: "/blog/🐶" },
  { name: "What Is The DOM", route: "/blog/what-is-the-dom" },
  {
    name: "Learning in 2016 and Onward",
    route: "/blog/learing-in-2016-and-onward"
  }
];

const projectsChildren = [
  { name: "Flippour", route: "/project/flippour" },
  { name: "Triangly", route: "/project/triangly" },
  {
    name: "SSR w/ CRA",
    route: "/project/server-side-rendering-with-create-react-app"
  }
];

export default [
  { name: "Home", route: "/page/home", children: [] },
  { name: "Resume", route: resume, children: [] },
  { name: "Projects", route: false, children: projectsChildren },
  { name: "Blog", route: false, children: blogChildren },
  { name: "Github", route: github, children: [] },
  { name: "Twitter", route: twitter, children: [] },
  { name: "Email", route: email, children: [] }
];
