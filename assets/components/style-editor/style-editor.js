import { debounce } from '../../lib/utils'
import { Component } from '../component'
export class StyleEditor extends Component {
  constructor() {
    super()
    this.socket = new WebSocket(`wss://${window.location.host}/ws`)
    const socket = new WebSocket('wss://localhost:8080/ws')
    socket.onopen = (evt) => {
      console.log('Connection open...')
    }
    socket.onmessage = (msg) => {
      console.log('Got message', msg)
    }
    socket.onerror = (err) => {
      console.error(err)
    }
  }
  handleInput(evt) {
    console.log('input', evt)
  }
  handleKeyPress(evt) {
    if (evt.keyCode === 9) {
      // TODO: allow tab
    }
  }
  connectedCallback() {
    this.listen(this.slots.editor, 'input', debounce(this.handleInput))
    this.listen(this.slots.editor, 'keydown', this.handleKeyPress)
  }
}
customElements.define('linky-style-editor', StyleEditor)
