import { Component } from '../component'

// use this element to wrap content inside the body
// it mainly serves a sa controller for the rest
// of the components
class App extends Component {}
customElements.define('linky-app', App)
