import { debounce } from '../../lib/utils'
import { Component } from '../component'
export class StyleEditor extends Component {
  constructor() {
    super()
    this.socket = new WebSocket(`wss://${window.location.host}/ws/style-editor`)
    this.socket.onopen = (evt) => {
      console.log('Connection open...')
    }
    this.socket.onmessage = (msg) => {
      console.log('Got message', msg.data)
      this.slots.editor.value = msg.data
    }
    this.socket.onerror = (err) => {
      console.error(err)
    }
  }
  handleInput(evt) {
    console.log('input', evt.target.value)
    this.socket.send(evt.target.value)
  }
  handleKeyPress(evt) {
    if (evt.keyCode === 9) {
      // TODO: allow tab
    }
  }
  connectedCallback() {
    this.listen(
      this.slots.editor,
      'input',
      debounce(this.handleInput.bind(this), 500)
    )
    this.listen(this.slots.editor, 'keydown', this.handleKeyPress)
  }
}
customElements.define('linky-style-editor', StyleEditor)
