import './app.css'

// use this element to wrap content inside the body
// it mainly serves a sa controller for the rest
// of the components
class App extends HTMLElement {}
customElements.define('linky-app', App)
