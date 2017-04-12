export default {
  title: "What is the DOM?",
  link: "what-is-the-dom",
  body: `
    What is the Document Object Model? 
    <br/>
    <br/>
    HTML != DOM 
    <br/>
    <br/>
    Is my HTML document DOM? No. Is a page’s source code DOM? No. What about when we inspect a page? Are we see DOM then? Closer! 
    <br/>
    <br/>
    First up, the DOM is the logical representation of a webpage derived from the page’s HTML. The DOM can and usually is manipulated by Javascript in some way, and in a way we can think of Javascript as a DOM API (which may offend some). Most think of the DOM as an API itself. 
    <br/>
    <br/>
    Via the DOM we can (1) add content to an HTML document, (2) delete content from an HTML document, and (3) Change contents within an HTML document. All without a page refresh. We are granted the opportunity to interact with a webpage because of the DOM. 
    <br/>
    <br/>
    What happens when a browser interprets out HTML? When do the tags <​span​>, <​ul​>, and <​p​>’s stop being HTML? The browser receives our HTML and renders it into DOM immediately. HTML elements (<​head​>, <​body​>, and <​nav​>’s) become DOM objects/nodes. This is where the DOM API (lol) Javascript comes into play, Javascript can interact very nicely with objects. Calling APIs on them to do the heavy, hard-hitting add/deleting/changing. 
    <br/>
    <br/>
    Example: document .getElementById('intro') .style .display = 'none'; would stop displaying the DOM object created by the HTML block with the id attribute of \`intro.\` 
    <br/>
    <br/>
    Is a DOM node different from a DOM object? A DOM object is sort of a parent, or superset of the node. Things that are liable to be changed are what nodes are. HTML elements, text within elements, HTML attributes (classes, styles, href’s), and so on. A collection of nodes are stored within every DOM object. 
    <br/>
    <br/>
    What can we do with the DOM? Absolute everything. Changing styles, Animation, talking to the user, attacking the user (looking at you MDE), propaganda (looking at you North Korea), and everything else. 
    <br/>
    <br/>
    The DOM is the web developer’s medium.
`,
  image: require("./whatIsTheDom.image.jpg")
};
